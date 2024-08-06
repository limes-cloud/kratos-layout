package app

import (
	"context"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"
	pb "poverty/api/poverty/chat/v1"
	"poverty/api/poverty/errors"
	"poverty/internal/conf"
	"poverty/internal/domain/service"
	"poverty/internal/infra/dbs"
	"poverty/internal/infra/httpx"
	"poverty/internal/infra/rpc"
	"poverty/internal/types"
)

type ChatApp struct {
	conf *conf.Config
	pb.UnimplementedChatServer
	srv *service.ChatService
}

func NewChatApp(conf *conf.Config) *ChatApp {
	bot, err := httpx.NewChatBotInfra(conf.AiBot.Secret, conf.AiBot.Appid)
	if err != nil {
		panic(err)
	}
	return &ChatApp{
		conf: conf,
		srv:  service.NewChatService(conf, dbs.NewChatInfra(), bot, rpc.NewUserInfra()),
	}
}

func init() {
	register(func(c *conf.Config, hs *http.Server, gs *grpc.Server) {
		srv := NewChatApp(c)
		pb.RegisterChatHTTPServer(hs, srv)
		pb.RegisterChatServer(gs, srv)

		cr := hs.Route("/")
		cr.POST("/poverty/client/v1/chat/send", srv.SendChatMessage)
	})
}

// SendChatMessage 发送会话信息
func (s *ChatApp) SendChatMessage(c http.Context) error {
	var (
		req pb.SendChatMessageRequest
	)
	if err := c.Bind(&req); err != nil {
		return errors.ParamsError()
	}

	h := c.Middleware(func(ctx context.Context, req any) (any, error) {
		return nil, s.send(ctx, c, req.(*pb.SendChatMessageRequest))
	})

	_, err := h(c, &req)
	return err
}

// SendChatMessage 发送会话信息
func (s *ChatApp) send(c context.Context, httpCtx http.Context, req *pb.SendChatMessageRequest) error {
	var ctx = kratosx.MustContext(c)
	// 第一次请求，自动携带sid
	if req.SessionId == nil || *req.SessionId == "" {
		sid, err := s.srv.GenSessionId(ctx)
		if err != nil {
			return err
		}
		req.SessionId = &sid
	}

	httpCtx.Response().Header().Set("Content-Type", "text/event-stream")
	httpCtx.Response().Header().Set("Cache-Control", "no-cache")
	httpCtx.Response().Header().Set("Connection", "keep-alive")

	flusher, ok := httpCtx.Response().(http.Flusher)
	if !ok {
		return errors.SystemError()
	}

	closer := make(chan struct{}, 1)

	// 发送聊天
	err := s.srv.SendChatMessage(ctx, &types.SendChatMessageRequest{
		SessionId: *req.SessionId,
		Message:   req.Message,
		Reply: func(msg string) {
			reply := pb.SendChatMessageReply{
				SessionId: *req.SessionId,
				Message:   msg,
			}
			data, _ := json.Marshal(&reply)
			data = append(data, []byte("\n")...)
			httpCtx.Response().Write(data)
			flusher.Flush()
		},
		Done: func() {
			closer <- struct{}{}
		},
	})
	if err != nil {
		return err
	}

	// 等待回复完成
	select {
	case <-closer:
	case <-ctx.Done():
	}
	close(closer)
	return nil
}

func (s *ChatApp) GetChatSetting(c context.Context, _ *pb.GetChatSettingRequest) (*pb.GetChatSettingReply, error) {
	reply := pb.GetChatSettingReply{}
	if err := valx.Transform(s.conf.AiBot.Setting, &reply); err != nil {
		kratosx.MustContext(c).Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.TransformError()
	}
	return &reply, nil
}

// ListChatRecord 获取会话记录信息列表
func (s *ChatApp) ListChatRecord(c context.Context, req *pb.ListChatRecordRequest) (*pb.ListChatRecordReply, error) {
	var ctx = kratosx.MustContext(c)
	result, total, err := s.srv.ListChatRecord(ctx, &types.ListChatRecordRequest{
		Page:      req.Page,
		PageSize:  req.PageSize,
		Order:     req.Order,
		OrderBy:   req.OrderBy,
		UserId:    req.UserId,
		SessionId: req.SessionId,
		Distinct:  req.Distinct,
		UserName:  req.UserName,
	})
	if err != nil {
		return nil, err
	}

	reply := pb.ListChatRecordReply{Total: total}
	if err := valx.Transform(result, &reply.List); err != nil {
		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	return &reply, nil
}

// DeleteChatRecord 删除会话记录信息
func (s *ChatApp) DeleteChatRecord(c context.Context, req *pb.DeleteChatRecordRequest) (*pb.DeleteChatRecordReply, error) {
	return &pb.DeleteChatRecordReply{}, s.srv.DeleteChatRecord(kratosx.MustContext(c), req.Id)
}
