package repository

import (
	"partyaffairs/internal/core"
	"partyaffairs/internal/domain/entity"
	"partyaffairs/internal/types"
)

type BannerRepository interface {
	// GetBanner 获取指定的轮播图信息
	GetBanner(ctx core.Context, id uint32) (*entity.Banner, error)

	// ListBanner 获取轮播图信息列表
	ListBanner(ctx core.Context, req *types.ListBannerRequest) ([]*entity.Banner, uint32, error)

	// CreateBanner 创建轮播图信息
	CreateBanner(ctx core.Context, req *entity.Banner) (uint32, error)

	// UpdateBanner 更新轮播图信息
	UpdateBanner(ctx core.Context, req *entity.Banner) error

	// DeleteBanner 删除轮播图信息
	DeleteBanner(ctx core.Context, id uint32) error
}
