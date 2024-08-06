package app

import (
	"context"
	"poverty/internal/infra/rpc"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"

	pb "poverty/api/poverty/banner/v1"
	"poverty/api/poverty/errors"
	"poverty/internal/conf"
	"poverty/internal/domain/entity"
	"poverty/internal/domain/service"
	"poverty/internal/infra/dbs"
	"poverty/internal/types"
)

type BannerApp struct {
	pb.UnimplementedBannerServer
	srv *service.BannerService
}

func NewBannerApp(conf *conf.Config) *BannerApp {
	return &BannerApp{
		srv: service.NewBannerService(conf, dbs.NewBannerInfra(), rpc.NewFileInfra()),
	}
}

func init() {
	register(func(c *conf.Config, hs *http.Server, gs *grpc.Server) {
		srv := NewBannerApp(c)
		pb.RegisterBannerHTTPServer(hs, srv)
		pb.RegisterBannerServer(gs, srv)
	})
}

// ListBanner 获取轮播图信息列表
func (s *BannerApp) ListBanner(c context.Context, req *pb.ListBannerRequest) (*pb.ListBannerReply, error) {
	var ctx = kratosx.MustContext(c)
	result, total, err := s.srv.ListBanner(ctx, &types.ListBannerRequest{
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

	reply := pb.ListBannerReply{Total: total}
	if err := valx.Transform(result, &reply.List); err != nil {
		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	return &reply, nil
}

// CreateBanner 创建轮播图信息
func (s *BannerApp) CreateBanner(c context.Context, req *pb.CreateBannerRequest) (*pb.CreateBannerReply, error) {
	var (
		ent = entity.Banner{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &ent); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	id, err := s.srv.CreateBanner(ctx, &ent)
	if err != nil {
		return nil, err
	}

	return &pb.CreateBannerReply{Id: id}, nil
}

// UpdateBanner 更新轮播图信息
func (s *BannerApp) UpdateBanner(c context.Context, req *pb.UpdateBannerRequest) (*pb.UpdateBannerReply, error) {
	var (
		ent = entity.Banner{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &ent); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	if err := s.srv.UpdateBanner(ctx, &ent); err != nil {
		return nil, err
	}

	return &pb.UpdateBannerReply{}, nil
}

// DeleteBanner 删除轮播图信息
func (s *BannerApp) DeleteBanner(c context.Context, req *pb.DeleteBannerRequest) (*pb.DeleteBannerReply, error) {
	return &pb.DeleteBannerReply{}, s.srv.DeleteBanner(kratosx.MustContext(c), req.Id)
}
