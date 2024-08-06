package app

import (
	"context"
	"poverty/internal/infra/rpc"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"

	pb "poverty/api/poverty/activity/v1"
	"poverty/api/poverty/errors"
	"poverty/internal/conf"
	"poverty/internal/domain/entity"
	"poverty/internal/domain/service"
	"poverty/internal/infra/dbs"
	"poverty/internal/types"
)

type ActivityApp struct {
	pb.UnimplementedActivityServer
	srv *service.ActivityService
}

func NewActivityApp(conf *conf.Config) *ActivityApp {
	return &ActivityApp{
		srv: service.NewActivityService(conf, dbs.NewActivityInfra(), rpc.NewFileInfra()),
	}
}

func init() {
	register(func(c *conf.Config, hs *http.Server, gs *grpc.Server) {
		srv := NewActivityApp(c)
		pb.RegisterActivityHTTPServer(hs, srv)
		pb.RegisterActivityServer(gs, srv)
	})
}

// GetActivity 获取指定的活动信息
func (s *ActivityApp) GetActivity(c context.Context, req *pb.GetActivityRequest) (*pb.GetActivityReply, error) {
	var ctx = kratosx.MustContext(c)

	ent, err := s.srv.GetActivity(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	reply := pb.GetActivityReply{}
	if err := valx.Transform(ent, &reply); err != nil {
		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.TransformError()
	}
	return &reply, nil
}

// ListActivity 获取活动信息列表
func (s *ActivityApp) ListActivity(c context.Context, req *pb.ListActivityRequest) (*pb.ListActivityReply, error) {
	var ctx = kratosx.MustContext(c)
	result, total, err := s.srv.ListActivity(ctx, &types.ListActivityRequest{
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

	reply := pb.ListActivityReply{Total: total}
	if err := valx.Transform(result, &reply.List); err != nil {
		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	return &reply, nil
}

// CreateActivity 创建活动信息
func (s *ActivityApp) CreateActivity(c context.Context, req *pb.CreateActivityRequest) (*pb.CreateActivityReply, error) {
	var (
		ent = entity.Activity{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &ent); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	id, err := s.srv.CreateActivity(ctx, &ent)
	if err != nil {
		return nil, err
	}

	return &pb.CreateActivityReply{Id: id}, nil
}

// UpdateActivity 更新活动信息
func (s *ActivityApp) UpdateActivity(c context.Context, req *pb.UpdateActivityRequest) (*pb.UpdateActivityReply, error) {
	var (
		ent = entity.Activity{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &ent); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	if err := s.srv.UpdateActivity(ctx, &ent); err != nil {
		return nil, err
	}

	return &pb.UpdateActivityReply{}, nil
}

// DeleteActivity 删除活动信息
func (s *ActivityApp) DeleteActivity(c context.Context, req *pb.DeleteActivityRequest) (*pb.DeleteActivityReply, error) {
	return &pb.DeleteActivityReply{}, s.srv.DeleteActivity(kratosx.MustContext(c), req.Id)
}
