package dbs

import (
	"github.com/limes-cloud/kratosx"

	"partyaffairs/internal/domain/entity"
	"partyaffairs/internal/types"
)

type Activity struct {
}

func NewActivity() *Activity {
	return &Activity{}
}

// GetActivity 获取指定的活动信息数据
func (r *Activity) GetActivity(ctx kratosx.Context, id uint32) (*entity.Activity, error) {
	var (
		ent = entity.Activity{}
		fs  = []string{"*"}
	)
	db := ctx.DB().Select(fs)
	return &ent, db.First(&ent, id).Error
}

// ListActivity 获取活动信息列表
func (r *Activity) ListActivity(ctx kratosx.Context, req *types.ListActivityRequest) ([]*entity.Activity, uint32, error) {
	var (
		list  []*entity.Activity
		fs    = []string{"*"}
		total int64
	)

	db := ctx.DB().Model(entity.Activity{})
	db = db.Select(fs)

	if req.Title != nil {
		db = db.Where("title LIKE ?", *req.Title+"%")
	}
	if req.IsTop != nil {
		db = db.Where("is_top = ?", *req.IsTop)
	}
	if req.Status != nil {
		db = db.Where("status = ?", *req.Status)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	db = db.Offset(int((req.Page - 1) * req.PageSize)).Limit(int(req.PageSize))

	return list, uint32(total), db.Find(&list).Error
}

// CreateActivity 创建活动信息数据
func (r *Activity) CreateActivity(ctx kratosx.Context, ent *entity.Activity) (uint32, error) {
	return ent.Id, ctx.DB().Create(ent).Error
}

// UpdateActivity 更新活动信息数据
func (r *Activity) UpdateActivity(ctx kratosx.Context, ent *entity.Activity) error {
	return ctx.DB().Updates(ent).Error
}

// DeleteActivity 删除活动信息数据
func (r *Activity) DeleteActivity(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(&entity.Activity{}, id).Error
}
