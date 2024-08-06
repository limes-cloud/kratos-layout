package app

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"

	"poverty/api/poverty/errors"
	pb "poverty/api/poverty/notice/v1"
	"poverty/internal/conf"
	"poverty/internal/domain/entity"
	"poverty/internal/domain/service"
	"poverty/internal/infra/dbs"
	"poverty/internal/types"
)

type NoticeApp struct {
	pb.UnimplementedNoticeServer
	srv *service.NoticeService
}

func NewNoticeApp(conf *conf.Config) *NoticeApp {
	return &NoticeApp{
		srv: service.NewNoticeService(conf, dbs.NewNoticeInfra()),
	}
}

func init() {
	register(func(c *conf.Config, hs *http.Server, gs *grpc.Server) {
		srv := NewNoticeApp(c)
		pb.RegisterNoticeHTTPServer(hs, srv)
		pb.RegisterNoticeServer(gs, srv)
	})
}

// GetNotice 获取指定的通知信息
func (s *NoticeApp) GetNotice(c context.Context, req *pb.GetNoticeRequest) (*pb.GetNoticeReply, error) {
	var ctx = kratosx.MustContext(c)

	ent, err := s.srv.GetNotice(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	reply := pb.GetNoticeReply{}
	if err := valx.Transform(ent, &reply); err != nil {
		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.TransformError()
	}
	return &reply, nil
}

// ListNotice 获取通知信息列表
func (s *NoticeApp) ListNotice(c context.Context, req *pb.ListNoticeRequest) (*pb.ListNoticeReply, error) {
	var ctx = kratosx.MustContext(c)
	result, total, err := s.srv.ListNotice(ctx, &types.ListNoticeRequest{
		Page:     req.Page,
		PageSize: req.PageSize,
		Order:    req.Order,
		OrderBy:  req.OrderBy,
		Title:    req.Title,
		IsTop:    req.IsTop,
		Status:   req.Status,
	})
	if err != nil {
		return nil, err
	}

	reply := pb.ListNoticeReply{Total: total}
	if err := valx.Transform(result, &reply.List); err != nil {
		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	return &reply, nil
}

// CreateNotice 创建通知信息
func (s *NoticeApp) CreateNotice(c context.Context, req *pb.CreateNoticeRequest) (*pb.CreateNoticeReply, error) {
	var (
		ent = entity.Notice{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &ent); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	id, err := s.srv.CreateNotice(ctx, &ent)
	if err != nil {
		return nil, err
	}

	return &pb.CreateNoticeReply{Id: id}, nil
}

// UpdateNotice 更新通知信息
func (s *NoticeApp) UpdateNotice(c context.Context, req *pb.UpdateNoticeRequest) (*pb.UpdateNoticeReply, error) {
	var (
		ent = entity.Notice{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &ent); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	if err := s.srv.UpdateNotice(ctx, &ent); err != nil {
		return nil, err
	}

	return &pb.UpdateNoticeReply{}, nil
}

// DeleteNotice 删除通知信息
func (s *NoticeApp) DeleteNotice(c context.Context, req *pb.DeleteNoticeRequest) (*pb.DeleteNoticeReply, error) {
	return &pb.DeleteNoticeReply{}, s.srv.DeleteNotice(kratosx.MustContext(c), req.Id)
}
