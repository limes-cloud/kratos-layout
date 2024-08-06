package dbs

import (
	"fmt"

	"github.com/limes-cloud/kratosx"
	"google.golang.org/protobuf/proto"

	"poverty/internal/domain/entity"
	"poverty/internal/types"
)

type ActivityInfra struct {
}

func NewActivityInfra() *ActivityInfra {
	return &ActivityInfra{}
}

// GetActivity 获取指定的活动信息数据
func (r *ActivityInfra) GetActivity(ctx kratosx.Context, id uint32) (*entity.Activity, error) {
	var (
		ent = entity.Activity{}
		fs  = []string{"*"}
	)
	db := ctx.DB().Select(fs)
	return &ent, db.First(&ent, id).Error
}

// ListActivity 获取活动信息列表
func (r *ActivityInfra) ListActivity(ctx kratosx.Context, req *types.ListActivityRequest) ([]*entity.Activity, uint32, error) {
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

	if req.OrderBy == nil || *req.OrderBy == "" {
		req.OrderBy = proto.String("id")
	}
	if req.Order == nil || *req.Order == "" {
		req.Order = proto.String("desc")
	}

	db = db.Order(fmt.Sprintf("%s %s", *req.OrderBy, *req.Order))
	if *req.OrderBy != "id" {
		db = db.Order("id desc")
	}

	return list, uint32(total), db.Find(&list).Error
}

// CreateActivity 创建活动信息数据
func (r *ActivityInfra) CreateActivity(ctx kratosx.Context, ent *entity.Activity) (uint32, error) {
	return ent.Id, ctx.DB().Create(ent).Error
}

// UpdateActivity 更新活动信息数据
func (r *ActivityInfra) UpdateActivity(ctx kratosx.Context, ent *entity.Activity) error {
	return ctx.DB().Updates(ent).Error
}

// DeleteActivity 删除活动信息数据
func (r *ActivityInfra) DeleteActivity(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(&entity.Activity{}, id).Error
}
