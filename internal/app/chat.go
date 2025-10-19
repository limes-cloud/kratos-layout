package app

import (
	"context"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx/pkg/value"
	"partyaffairs/api/chat"
	"partyaffairs/api/errors"
	"partyaffairs/internal/core"
	"partyaffairs/internal/domain/service"
	"partyaffairs/internal/infra/dbs"
	"partyaffairs/internal/infra/httpx"
	"partyaffairs/internal/infra/rpc"
	"partyaffairs/internal/types"
)

type ChatApp struct {
	chat.UnimplementedChatServer
	srv *service.ChatService
}

func NewChatApp() *ChatApp {
	return &ChatApp{
		srv: service.NewChatService(dbs.NewChat(), httpx.NewChatBot(), rpc.NewUser()),
	}
}

func init() {
	register(func(hs *http.Server) {
		srv := NewChatApp()
		chat.RegisterChatHTTPServer(hs, srv)

		cr := hs.Route("/")
		cr.POST("/partyaffairs/api/v1/chat/send", srv.SendChatMessage)
	})
}

// SendChatMessage 发送会话信息
func (s *ChatApp) SendChatMessage(c http.Context) error {
	var (
		req chat.SendChatMessageRequest
	)
	if err := c.Bind(&req); err != nil {
		return errors.ParamsError()
	}

	h := c.Middleware(func(ctx context.Context, req any) (any, error) {
		return nil, s.send(ctx, c, req.(*chat.SendChatMessageRequest))
	})

	_, err := h(c, &req)
	return err
}

// SendChatMessage 发送会话信息
func (s *ChatApp) send(c context.Context, httpCtx http.Context, req *chat.SendChatMessageRequest) error {
	var ctx = core.MustContext(c)
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
			reply := chat.SendChatMessageReply{
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

func (s *ChatApp) GetChatSetting(c context.Context, _ *chat.GetChatSettingRequest) (*chat.GetChatSettingReply, error) {
	var ctx = core.MustContext(c)
	reply := chat.GetChatSettingReply{}
	if err := value.Transform(ctx.Config().AiBot.Setting, &reply); err != nil {
		core.MustContext(c).Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.TransformError()
	}
	return &reply, nil
}

// ListChatRecord 获取会话记录信息列表
func (s *ChatApp) ListChatRecord(c context.Context, req *chat.ListChatRecordRequest) (*chat.ListChatRecordReply, error) {
	var ctx = core.MustContext(c)
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

	reply := chat.ListChatRecordReply{Total: total}
	if err := value.Transform(result, &reply.List); err != nil {
		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	return &reply, nil
}

// DeleteChatRecord 删除会话记录信息
func (s *ChatApp) DeleteChatRecord(c context.Context, req *chat.DeleteChatRecordRequest) (*chat.DeleteChatRecordReply, error) {
	return &chat.DeleteChatRecordReply{}, s.srv.DeleteChatRecord(core.MustContext(c), req.Id)
}
