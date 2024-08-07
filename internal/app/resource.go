package app

import (
	"context"
	"poverty/internal/infra/rpc"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"

	"poverty/api/poverty/errors"
	pb "poverty/api/poverty/resource/v1"
	"poverty/internal/conf"
	"poverty/internal/domain/entity"
	"poverty/internal/domain/service"
	"poverty/internal/infra/dbs"
	"poverty/internal/types"
)

type ResourceApp struct {
	pb.UnimplementedResourceServer
	srv *service.ResourceService
}

func NewResourceApp(conf *conf.Config) *ResourceApp {
	return &ResourceApp{
		srv: service.NewResourceService(conf, dbs.NewResourceInfra(), rpc.NewFileInfra()),
	}
}

func init() {
	register(func(c *conf.Config, hs *http.Server, gs *grpc.Server) {
		srv := NewResourceApp(c)
		pb.RegisterResourceHTTPServer(hs, srv)
		pb.RegisterResourceServer(gs, srv)
	})
}

// GetResource 获取指定的文件信息
func (s *ResourceApp) GetResource(c context.Context, req *pb.GetResourceRequest) (*pb.GetResourceReply, error) {
	var ctx = kratosx.MustContext(c)

	ent, err := s.srv.GetResource(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	reply := pb.GetResourceReply{}
	if err := valx.Transform(ent, &reply); err != nil {
		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.TransformError()
	}
	return &reply, nil
}

// ListResource 获取文件信息列表
func (s *ResourceApp) ListResource(c context.Context, req *pb.ListResourceRequest) (*pb.ListResourceReply, error) {
	var ctx = kratosx.MustContext(c)
	result, total, err := s.srv.ListResource(ctx, &types.ListResourceRequest{
		Page:     req.Page,
		PageSize: req.PageSize,
		Order:    req.Order,
		OrderBy:  req.OrderBy,
		Title:    req.Title,
		Status:   req.Status,
	})
	if err != nil {
		return nil, err
	}

	reply := pb.ListResourceReply{Total: total}
	if err := valx.Transform(result, &reply.List); err != nil {
		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	return &reply, nil
}

// CreateResource 创建文件信息
func (s *ResourceApp) CreateResource(c context.Context, req *pb.CreateResourceRequest) (*pb.CreateResourceReply, error) {
	var (
		ent = entity.Resource{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &ent); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	id, err := s.srv.CreateResource(ctx, &ent)
	if err != nil {
		return nil, err
	}

	return &pb.CreateResourceReply{Id: id}, nil
}

// UpdateResource 更新文件信息
func (s *ResourceApp) UpdateResource(c context.Context, req *pb.UpdateResourceRequest) (*pb.UpdateResourceReply, error) {
	var (
		ent = entity.Resource{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &ent); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	if err := s.srv.UpdateResource(ctx, &ent); err != nil {
		return nil, err
	}

	return &pb.UpdateResourceReply{}, nil
}

// DeleteResource 删除文件信息
func (s *ResourceApp) DeleteResource(c context.Context, req *pb.DeleteResourceRequest) (*pb.DeleteResourceReply, error) {
	return &pb.DeleteResourceReply{}, s.srv.DeleteResource(kratosx.MustContext(c), req.Id)
}
