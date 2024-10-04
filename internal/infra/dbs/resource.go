package dbs

import (
	"github.com/limes-cloud/kratosx"

	"partyaffairs/internal/domain/entity"
	"partyaffairs/internal/types"
)

type Resource struct {
}

func NewResource() *Resource {
	return &Resource{}
}

// ListResourceClassify 获取列表
func (t *Resource) ListResourceClassify(ctx kratosx.Context) ([]*entity.ResourceClassify, error) {
	var (
		list []*entity.ResourceClassify
		fs   = []string{"*"}
	)
	return list, ctx.DB().Model(entity.ResourceClassify{}).Select(fs).Find(&list).Error
}

// CreateResourceClassify 创建数据
func (t *Resource) CreateResourceClassify(ctx kratosx.Context, tg *entity.ResourceClassify) (uint32, error) {
	return tg.Id, ctx.DB().Create(tg).Error
}

// UpdateResourceClassify 更新数据
func (t *Resource) UpdateResourceClassify(ctx kratosx.Context, tg *entity.ResourceClassify) error {
	return ctx.DB().Updates(tg).Error
}

// DeleteResourceClassify 删除数据
func (t *Resource) DeleteResourceClassify(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Where("id = ?", id).Delete(&entity.ResourceClassify{}).Error
}

// GetResource 获取指定的资讯信息数据
func (r *Resource) GetResource(ctx kratosx.Context, id uint32) (*entity.Resource, error) {
	var (
		ent = entity.Resource{}
		fs  = []string{"*"}
	)
	db := ctx.DB().Select(fs).Preload("Classify")
	return &ent, db.First(&ent, id).Error
}

// ListResource 获取资讯信息列表
func (r *Resource) ListResource(ctx kratosx.Context, req *types.ListResourceRequest) ([]*entity.Resource, uint32, error) {
	var (
		list  []*entity.Resource
		fs    = []string{"*"}
		total int64
	)

	db := ctx.DB().Model(entity.Resource{})
	db = db.Select(fs).Preload("Classify")

	if req.Title != nil {
		db = db.Where("title LIKE ?", *req.Title+"%")
	}
	if req.ClassifyId != nil {
		db = db.Where("classify_id = ?", *req.ClassifyId)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	db = db.Offset(int((req.Page - 1) * req.PageSize)).Limit(int(req.PageSize))

	return list, uint32(total), db.Find(&list).Error
}

// CreateResource 创建资讯信息数据
func (r *Resource) CreateResource(ctx kratosx.Context, ent *entity.Resource) (uint32, error) {
	return ent.Id, ctx.DB().Create(ent).Error
}

// UpdateResource 更新资讯信息数据
func (r *Resource) UpdateResource(ctx kratosx.Context, ent *entity.Resource) error {
	return ctx.DB().Updates(ent).Error
}

// DeleteResource 删除资讯信息数据
func (r *Resource) DeleteResource(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(&entity.Resource{}, id).Error
}
