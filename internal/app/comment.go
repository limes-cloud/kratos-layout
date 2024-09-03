package app

import (
	"context"
	"poverty/internal/infra/rpc"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"

	pb "poverty/api/poverty/comment/v1"
	"poverty/api/poverty/errors"
	"poverty/internal/conf"
	"poverty/internal/domain/entity"
	"poverty/internal/domain/service"
	"poverty/internal/infra/dbs"
	"poverty/internal/types"
)

type CommentApp struct {
	pb.UnimplementedCommentServer
	srv *service.CommentService
}

func NewCommentApp(conf *conf.Config) *CommentApp {
	return &CommentApp{
		srv: service.NewCommentService(conf, dbs.NewCommentInfra(), rpc.NewUserInfra()),
	}
}

func init() {
	register(func(c *conf.Config, hs *http.Server, gs *grpc.Server) {
		srv := NewCommentApp(c)
		pb.RegisterCommentHTTPServer(hs, srv)
		pb.RegisterCommentServer(gs, srv)
	})
}

// ListComment 获取留言信息列表
func (s *CommentApp) ListComment(c context.Context, req *pb.ListCommentRequest) (*pb.ListCommentReply, error) {
	var ctx = kratosx.MustContext(c)
	result, total, err := s.srv.ListComment(ctx, &types.ListCommentRequest{
		Page:     req.Page,
		PageSize: req.PageSize,
		Order:    req.Order,
		OrderBy:  req.OrderBy,
		UserName: req.UserName,
	})
	if err != nil {
		return nil, err
	}

	reply := pb.ListCommentReply{Total: total}
	if err := valx.Transform(result, &reply.List); err != nil {
		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	return &reply, nil
}

// CreateComment 创建留言信息
func (s *CommentApp) CreateComment(c context.Context, req *pb.CreateCommentRequest) (*pb.CreateCommentReply, error) {
	var (
		ent = entity.Comment{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &ent); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	id, err := s.srv.CreateComment(ctx, &ent)
	if err != nil {
		return nil, err
	}

	return &pb.CreateCommentReply{Id: id}, nil
}

// DeleteComment 删除留言信息
func (s *CommentApp) DeleteComment(c context.Context, req *pb.DeleteCommentRequest) (*pb.DeleteCommentReply, error) {
	return &pb.DeleteCommentReply{}, s.srv.DeleteComment(kratosx.MustContext(c), req.Id)
}
