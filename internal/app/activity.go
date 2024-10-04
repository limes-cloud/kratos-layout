package app

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"

	pb "partyaffairs/api/partyaffairs/activity/v1"
	"partyaffairs/api/partyaffairs/errors"
	"partyaffairs/internal/conf"
	"partyaffairs/internal/domain/entity"
	"partyaffairs/internal/domain/service"
	"partyaffairs/internal/infra/dbs"
	"partyaffairs/internal/infra/rpc"
	"partyaffairs/internal/types"
)

type Activity struct {
	pb.UnimplementedActivityServer
	srv *service.ActivityService
}

func NewActivity(conf *conf.Config) *Activity {
	return &Activity{
		srv: service.NewActivityService(conf, dbs.NewActivity(), rpc.NewFile()),
	}
}

func init() {
	register(func(c *conf.Config, hs *http.Server) {
		srv := NewActivity(c)
		pb.RegisterActivityHTTPServer(hs, srv)
	})
}

// GetActivity 获取指定的活动信息
func (s *Activity) GetActivity(c context.Context, req *pb.GetActivityRequest) (*pb.GetActivityReply, error) {
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
func (s *Activity) ListActivity(c context.Context, req *pb.ListActivityRequest) (*pb.ListActivityReply, error) {
	var ctx = kratosx.MustContext(c)
	result, total, err := s.srv.ListActivity(ctx, &types.ListActivityRequest{
		Page:     req.Page,
		PageSize: req.PageSize,
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
func (s *Activity) CreateActivity(c context.Context, req *pb.CreateActivityRequest) (*pb.CreateActivityReply, error) {
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
func (s *Activity) UpdateActivity(c context.Context, req *pb.UpdateActivityRequest) (*pb.UpdateActivityReply, error) {
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
func (s *Activity) DeleteActivity(c context.Context, req *pb.DeleteActivityRequest) (*pb.DeleteActivityReply, error) {
	return &pb.DeleteActivityReply{}, s.srv.DeleteActivity(kratosx.MustContext(c), req.Id)
}
