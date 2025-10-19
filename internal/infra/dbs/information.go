package dbs

import (
	"gorm.io/gorm"
	"partyaffairs/internal/core"
	"partyaffairs/internal/domain/entity"
	"partyaffairs/internal/types"
)

type Information struct {
}

func NewInformation() *Information {
	return &Information{}
}

// ListInformationClassify 获取列表
func (t *Information) ListInformationClassify(ctx core.Context) ([]*entity.InformationClassify, error) {
	var (
		list []*entity.InformationClassify
		fs   = []string{"*"}
	)
	return list, ctx.DB().Model(entity.InformationClassify{}).Select(fs).Find(&list).Error
}

// CreateInformationClassify 创建数据
func (t *Information) CreateInformationClassify(ctx core.Context, tg *entity.InformationClassify) (uint32, error) {
	return tg.Id, ctx.DB().Create(tg).Error
}

// UpdateInformationClassify 更新数据
func (t *Information) UpdateInformationClassify(ctx core.Context, tg *entity.InformationClassify) error {
	return ctx.DB().Updates(tg).Error
}

// DeleteInformationClassify 删除数据
func (t *Information) DeleteInformationClassify(ctx core.Context, id uint32) error {
	return ctx.DB().Where("id = ?", id).Delete(&entity.InformationClassify{}).Error
}

// GetInformation 获取指定的资讯信息数据
func (r *Information) GetInformation(ctx core.Context, id uint32) (*entity.Information, error) {
	var (
		ent = entity.Information{}
		fs  = []string{"*"}
	)
	db := ctx.DB().Select(fs).Preload("Classify")
	return &ent, db.First(&ent, id).Error
}

// ListInformation 获取资讯信息列表
func (r *Information) ListInformation(ctx core.Context, req *types.ListInformationRequest) ([]*entity.Information, uint32, error) {
	var (
		list  []*entity.Information
		fs    = []string{"*"}
		total int64
	)

	db := ctx.DB().Model(entity.Information{})
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
func (r *Information) IncrReadCount(ctx core.Context, id uint32) error {
	return ctx.DB().
		Model(&entity.Information{}).
		Where("id = ?", id).
		Update("read", gorm.Expr("`read` + ?", 1)).Error
}

// CreateInformation 创建资讯信息数据
func (r *Information) CreateInformation(ctx core.Context, ent *entity.Information) (uint32, error) {
	return ent.Id, ctx.DB().Create(ent).Error
}

// UpdateInformation 更新资讯信息数据
func (r *Information) UpdateInformation(ctx core.Context, ent *entity.Information) error {
	return ctx.DB().Updates(ent).Error
}

// DeleteInformation 删除资讯信息数据
func (r *Information) DeleteInformation(ctx core.Context, id uint32) error {
	return ctx.DB().Delete(&entity.Information{}, id).Error
}
