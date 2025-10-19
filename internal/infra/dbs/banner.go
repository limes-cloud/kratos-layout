package dbs

import (
	"github.com/limes-cloud/kratosx/model/page"
	"github.com/limes-cloud/kratosx/pkg/value"
	"partyaffairs/internal/core"
	"partyaffairs/internal/domain/entity"
	"partyaffairs/internal/types"
)

type Banner struct {
}

func NewBanner() *Banner {
	return &Banner{}
}

// GetBanner 获取指定的轮播图信息数据
func (r *Banner) GetBanner(ctx core.Context, id uint32) (*entity.Banner, error) {
	var ent = entity.Banner{}
	return &ent, ctx.DB().First(&ent, id).Error
}

// ListBanner 获取轮播图信息列表
func (r *Banner) ListBanner(ctx core.Context, req *types.ListBannerRequest) ([]*entity.Banner, uint32, error) {
	var (
		list  []*entity.Banner
		total int64
	)

	db := ctx.DB().Model(entity.Banner{})
	if req.Title != nil {
		db = db.Where("title LIKE ?", *req.Title+"%")
	}
	if req.Status != nil {
		db = db.Where("status = ?", *req.Status)
	}
	if req.Search != nil {
		req.Order = value.Pointer("desc")
		req.OrderBy = value.Pointer("weight")

		if err := db.Count(&total).Error; err != nil {
			return nil, 0, err
		}
		db = page.SearchScopes(db, req.Search)
	}
	return list, uint32(total), db.Find(&list).Error
}

// CreateBanner 创建轮播图信息数据
func (r *Banner) CreateBanner(ctx core.Context, ent *entity.Banner) (uint32, error) {
	return ent.Id, ctx.DB().Create(ent).Error
}

// UpdateBanner 更新轮播图信息数据
func (r *Banner) UpdateBanner(ctx core.Context, ent *entity.Banner) error {
	return ctx.DB().Updates(ent).Error
}

// DeleteBanner 删除轮播图信息数据
func (r *Banner) DeleteBanner(ctx core.Context, id uint32) error {
	return ctx.DB().Delete(&entity.Banner{}, id).Error
}
