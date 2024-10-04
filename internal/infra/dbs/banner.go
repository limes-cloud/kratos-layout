package dbs

import (
	"github.com/limes-cloud/kratosx"

	"partyaffairs/internal/domain/entity"
	"partyaffairs/internal/types"
)

type Banner struct {
}

func NewBanner() *Banner {
	return &Banner{}
}

// GetBanner 获取指定的轮播图信息数据
func (r *Banner) GetBanner(ctx kratosx.Context, id uint32) (*entity.Banner, error) {
	var (
		ent = entity.Banner{}
		fs  = []string{"*"}
	)
	db := ctx.DB().Select(fs)
	return &ent, db.First(&ent, id).Error
}

// ListBanner 获取轮播图信息列表
func (r *Banner) ListBanner(ctx kratosx.Context, req *types.ListBannerRequest) ([]*entity.Banner, uint32, error) {
	var (
		list  []*entity.Banner
		fs    = []string{"*"}
		total int64
	)

	db := ctx.DB().Model(entity.Banner{})
	db = db.Select(fs)

	if req.Title != nil {
		db = db.Where("title LIKE ?", *req.Title+"%")
	}
	if req.Status != nil {
		db = db.Where("status = ?", *req.Status)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	db = db.Offset(int((req.Page - 1) * req.PageSize)).Limit(int(req.PageSize))
	db = db.Order("weight desc, id asc")
	return list, uint32(total), db.Find(&list).Error
}

// CreateBanner 创建轮播图信息数据
func (r *Banner) CreateBanner(ctx kratosx.Context, ent *entity.Banner) (uint32, error) {
	return ent.Id, ctx.DB().Create(ent).Error
}

// UpdateBanner 更新轮播图信息数据
func (r *Banner) UpdateBanner(ctx kratosx.Context, ent *entity.Banner) error {
	return ctx.DB().Updates(ent).Error
}

// DeleteBanner 删除轮播图信息数据
func (r *Banner) DeleteBanner(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(&entity.Banner{}, id).Error
}
