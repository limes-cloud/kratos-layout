package app

import (
	"context"
	"partyaffairs/api/activity"
	"partyaffairs/internal/core"

	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx/pkg/value"

	"partyaffairs/api/errors"
	"partyaffairs/internal/domain/entity"
	"partyaffairs/internal/domain/service"
	"partyaffairs/internal/infra/dbs"
	"partyaffairs/internal/infra/rpc"
	"partyaffairs/internal/types"
)

type Activity struct {
	activity.UnimplementedActivityServer
	srv *service.ActivityService
}

func NewActivity() *Activity {
	return &Activity{
		srv: service.NewActivityService(dbs.NewActivity(), rpc.NewFile()),
	}
}

func init() {
	register(func(hs *http.Server) {
		srv := NewActivity()
		activity.RegisterActivityHTTPServer(hs, srv)
	})
}

// GetActivity 获取指定的活动信息
func (s *Activity) GetActivity(c context.Context, req *activity.GetActivityRequest) (*activity.GetActivityReply, error) {
	var ctx = core.MustContext(c)
	ent, err := s.srv.GetActivity(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	reply := activity.GetActivityReply{}
	if err := value.Transform(ent, &reply); err != nil {
		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.TransformError()
	}
	return &reply, nil
}

// ListActivity 获取活动信息列表
func (s *Activity) ListActivity(c context.Context, req *activity.ListActivityRequest) (*activity.ListActivityReply, error) {
	var ctx = core.MustContext(c)
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

	reply := activity.ListActivityReply{Total: total}
	if err := value.Transform(result, &reply.List); err != nil {
		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	return &reply, nil
}

// CreateActivity 创建活动信息
func (s *Activity) CreateActivity(c context.Context, req *activity.CreateActivityRequest) (*activity.CreateActivityReply, error) {
	var (
		ent = entity.Activity{}
		ctx = core.MustContext(c)
	)

	if err := value.Transform(req, &ent); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	id, err := s.srv.CreateActivity(ctx, &ent)
	if err != nil {
		return nil, err
	}

	return &activity.CreateActivityReply{Id: id}, nil
}

// UpdateActivity 更新活动信息
func (s *Activity) UpdateActivity(c context.Context, req *activity.UpdateActivityRequest) (*activity.UpdateActivityReply, error) {
	var (
		ent = entity.Activity{}
		ctx = core.MustContext(c)
	)

	if err := value.Transform(req, &ent); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	if err := s.srv.UpdateActivity(ctx, &ent); err != nil {
		return nil, err
	}

	return &activity.UpdateActivityReply{}, nil
}

// DeleteActivity 删除活动信息
func (s *Activity) DeleteActivity(c context.Context, req *activity.DeleteActivityRequest) (*activity.DeleteActivityReply, error) {
	return &activity.DeleteActivityReply{}, s.srv.DeleteActivity(core.MustContext(c), req.Id)
}
