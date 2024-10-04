package app

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"
	ktypes "github.com/limes-cloud/kratosx/types"

	pb "partyaffairs/api/partyaffairs/banner/v1"
	"partyaffairs/api/partyaffairs/errors"
	"partyaffairs/internal/conf"
	"partyaffairs/internal/domain/entity"
	"partyaffairs/internal/domain/service"
	"partyaffairs/internal/infra/dbs"
	"partyaffairs/internal/infra/rpc"
	"partyaffairs/internal/types"
)

type Banner struct {
	pb.UnimplementedBannerServer
	srv *service.BannerService
}

func NewBanner(conf *conf.Config) *Banner {
	return &Banner{
		srv: service.NewBannerService(conf, dbs.NewBanner(), rpc.NewFile()),
	}
}

func init() {
	register(func(c *conf.Config, hs *http.Server) {
		srv := NewBanner(c)
		pb.RegisterBannerHTTPServer(hs, srv)
	})
}

// ListBanner 获取轮播图信息列表
func (s *Banner) ListBanner(c context.Context, req *pb.ListBannerRequest) (*pb.ListBannerReply, error) {
	var ctx = kratosx.MustContext(c)
	result, total, err := s.srv.ListBanner(ctx, &types.ListBannerRequest{
		Page:     req.Page,
		PageSize: req.PageSize,
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
func (s *Banner) CreateBanner(c context.Context, req *pb.CreateBannerRequest) (*pb.CreateBannerReply, error) {
	id, err := s.srv.CreateBanner(kratosx.MustContext(c), &entity.Banner{
		Title:  req.Title,
		Src:    req.Src,
		Path:   req.Path,
		Weight: req.Weight,
		Status: req.Status,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateBannerReply{Id: id}, nil
}

// UpdateBanner 更新轮播图信息
func (s *Banner) UpdateBanner(c context.Context, req *pb.UpdateBannerRequest) (*pb.UpdateBannerReply, error) {
	if err := s.srv.UpdateBanner(kratosx.MustContext(c), &entity.Banner{
		BaseModel: ktypes.BaseModel{Id: req.Id},
		Title:     req.Title,
		Src:       req.Src,
		Path:      req.Path,
		Weight:    req.Weight,
		Status:    req.Status,
	}); err != nil {
		return nil, err
	}
	return &pb.UpdateBannerReply{}, nil
}

// DeleteBanner 删除轮播图信息
func (s *Banner) DeleteBanner(c context.Context, req *pb.DeleteBannerRequest) (*pb.DeleteBannerReply, error) {
	return &pb.DeleteBannerReply{}, s.srv.DeleteBanner(kratosx.MustContext(c), req.Id)
}
