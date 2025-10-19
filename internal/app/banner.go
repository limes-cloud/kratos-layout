package app

import (
	"context"
	"github.com/limes-cloud/kratosx/model"
	"github.com/limes-cloud/kratosx/model/page"
	"google.golang.org/protobuf/types/known/emptypb"
	"partyaffairs/api/banner"
	"partyaffairs/internal/core"

	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx/pkg/value"

	"partyaffairs/api/errors"
	"partyaffairs/internal/domain/entity"
	"partyaffairs/internal/domain/service"
	"partyaffairs/internal/infra/dbs"
	"partyaffairs/internal/types"
)

type Banner struct {
	banner.UnimplementedBannerServer
	srv *service.BannerService
}

func NewBanner() *Banner {
	return &Banner{
		srv: service.NewBannerService(dbs.NewBanner()),
	}
}

func init() {
	register(func(hs *http.Server) {
		srv := NewBanner()
		banner.RegisterBannerHTTPServer(hs, srv)
	})
}

// ListVisibleBanner 获取客户端可见的轮播图信息列表
func (s *Banner) ListVisibleBanner(c context.Context, _ *emptypb.Empty) (*banner.ListBannerReply, error) {
	var ctx = core.MustContext(c)
	result, total, err := s.srv.ListBanner(ctx, &types.ListBannerRequest{
		Status: value.Pointer(true),
	})
	if err != nil {
		return nil, err
	}

	reply := banner.ListBannerReply{Total: total}
	if err := value.Transform(result, &reply.List); err != nil {
		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	return &reply, nil
}

// ListBanner 获取轮播图信息列表
func (s *Banner) ListBanner(c context.Context, req *banner.ListBannerRequest) (*banner.ListBannerReply, error) {
	var ctx = core.MustContext(c)
	result, total, err := s.srv.ListBanner(ctx, &types.ListBannerRequest{
		Search: &page.Search{
			Page:     req.Page,
			PageSize: req.PageSize,
		},
		Title:  req.Title,
		Status: req.Status,
	})
	if err != nil {
		return nil, err
	}

	reply := banner.ListBannerReply{Total: total}
	if err := value.Transform(result, &reply.List); err != nil {
		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	return &reply, nil
}

// CreateBanner 创建轮播图信息
func (s *Banner) CreateBanner(c context.Context, req *banner.CreateBannerRequest) (*banner.CreateBannerReply, error) {
	id, err := s.srv.CreateBanner(core.MustContext(c), &entity.Banner{
		Title:  req.Title,
		Key:    req.Key,
		Path:   req.Path,
		Weight: req.Weight,
		Status: req.Status,
	})
	if err != nil {
		return nil, err
	}
	return &banner.CreateBannerReply{Id: id}, nil
}

// UpdateBanner 更新轮播图信息
func (s *Banner) UpdateBanner(c context.Context, req *banner.UpdateBannerRequest) (*banner.UpdateBannerReply, error) {
	if err := s.srv.UpdateBanner(core.MustContext(c), &entity.Banner{
		BaseTenantModel: model.BaseTenantModel{Id: req.Id},
		Title:           req.Title,
		Key:             req.Key,
		Path:            req.Path,
		Weight:          req.Weight,
		Status:          req.Status,
	}); err != nil {
		return nil, err
	}
	return &banner.UpdateBannerReply{}, nil
}

// DeleteBanner 删除轮播图信息
func (s *Banner) DeleteBanner(c context.Context, req *banner.DeleteBannerRequest) (*banner.DeleteBannerReply, error) {
	return &banner.DeleteBannerReply{}, s.srv.DeleteBanner(core.MustContext(c), req.Id)
}
