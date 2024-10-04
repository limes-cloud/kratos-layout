package repository

import (
	"github.com/limes-cloud/kratosx"

	"partyaffairs/internal/domain/entity"
	"partyaffairs/internal/types"
)

type BannerRepository interface {
	// GetBanner 获取指定的轮播图信息
	GetBanner(ctx kratosx.Context, id uint32) (*entity.Banner, error)

	// ListBanner 获取轮播图信息列表
	ListBanner(ctx kratosx.Context, req *types.ListBannerRequest) ([]*entity.Banner, uint32, error)

	// CreateBanner 创建轮播图信息
	CreateBanner(ctx kratosx.Context, req *entity.Banner) (uint32, error)

	// UpdateBanner 更新轮播图信息
	UpdateBanner(ctx kratosx.Context, req *entity.Banner) error

	// DeleteBanner 删除轮播图信息
	DeleteBanner(ctx kratosx.Context, id uint32) error
}
