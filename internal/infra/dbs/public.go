package dbs

import (
	"gorm.io/gorm"
	"partyaffairs/internal/core"
	"partyaffairs/internal/domain/entity"
	"partyaffairs/internal/types"
)

type Public struct {
}

func NewPublic() *Public {
	return &Public{}
}

// ListPublicClassify 获取列表
func (t *Public) ListPublicClassify(ctx core.Context) ([]*entity.PublicClassify, error) {
	var (
		list []*entity.PublicClassify
		fs   = []string{"*"}
	)
	return list, ctx.DB().Model(entity.PublicClassify{}).Select(fs).Find(&list).Error
}

// CreatePublicClassify 创建数据
func (t *Public) CreatePublicClassify(ctx core.Context, tg *entity.PublicClassify) (uint32, error) {
	return tg.Id, ctx.DB().Create(tg).Error
}

// UpdatePublicClassify 更新数据
func (t *Public) UpdatePublicClassify(ctx core.Context, tg *entity.PublicClassify) error {
	return ctx.DB().Updates(tg).Error
}

// DeletePublicClassify 删除数据
func (t *Public) DeletePublicClassify(ctx core.Context, id uint32) error {
	return ctx.DB().Where("id = ?", id).Delete(&entity.PublicClassify{}).Error
}

// GetPublic 获取指定的资讯信息数据
func (r *Public) GetPublic(ctx core.Context, id uint32) (*entity.Public, error) {
	var (
		ent = entity.Public{}
		fs  = []string{"*"}
	)
	db := ctx.DB().Select(fs).Preload("Classify")
	return &ent, db.First(&ent, id).Error
}

// ListPublic 获取资讯信息列表
func (r *Public) ListPublic(ctx core.Context, req *types.ListPublicRequest) ([]*entity.Public, uint32, error) {
	var (
		list  []*entity.Public
		fs    = []string{"*"}
		total int64
	)

	db := ctx.DB().Model(entity.Public{})
	db = db.Select(fs).Preload("Classify")

	if req.Title != nil {
		db = db.Where("title LIKE ?", *req.Title+"%")
	}
	if req.IsTop != nil {
		db = db.Where("is_top = ?", *req.IsTop)
	}
	if req.Status != nil {
		db = db.Where("status = ?", *req.Status)
	}
	if req.ClassifyId != nil {
		db = db.Where("classify_id = ?", *req.ClassifyId)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	db = db.Order("is_top desc, id desc")
	db = db.Offset(int((req.Page - 1) * req.PageSize)).Limit(int(req.PageSize))

	return list, uint32(total), db.Find(&list).Error
}

// IncrReadCount 增加阅读次数
func (r *Public) IncrReadCount(ctx core.Context, id uint32) error {
	return ctx.DB().
		Model(&entity.Public{}).
		Where("id = ?", id).
		Update("read", gorm.Expr("`read` + ?", 1)).Error
}

// CreatePublic 创建资讯信息数据
func (r *Public) CreatePublic(ctx core.Context, ent *entity.Public) (uint32, error) {
	return ent.Id, ctx.DB().Create(ent).Error
}

// UpdatePublic 更新资讯信息数据
func (r *Public) UpdatePublic(ctx core.Context, ent *entity.Public) error {
	return ctx.DB().Updates(ent).Error
}

// DeletePublic 删除资讯信息数据
func (r *Public) DeletePublic(ctx core.Context, id uint32) error {
	return ctx.DB().Delete(&entity.Public{}, id).Error
}
