package dbs

import (
	"github.com/limes-cloud/kratosx"

	"partyaffairs/internal/domain/entity"
	"partyaffairs/internal/types"
)

type Notice struct {
}

func NewNotice() *Notice {
	return &Notice{}
}

// GetNotice 获取指定的通知信息数据
func (r *Notice) GetNotice(ctx kratosx.Context, id uint32) (*entity.Notice, error) {
	var (
		ent = entity.Notice{}
		fs  = []string{"*"}
	)
	db := ctx.DB().Select(fs)
	return &ent, db.First(&ent, id).Error
}

// ListNotice 获取通知信息列表
func (r *Notice) ListNotice(ctx kratosx.Context, req *types.ListNoticeRequest) ([]*entity.Notice, uint32, error) {
	var (
		list  []*entity.Notice
		fs    = []string{"*"}
		total int64
	)

	db := ctx.DB().Model(entity.Notice{})
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

// CreateNotice 创建通知信息数据
func (r *Notice) CreateNotice(ctx kratosx.Context, ent *entity.Notice) (uint32, error) {
	return ent.Id, ctx.DB().Create(ent).Error
}

// UpdateNotice 更新通知信息数据
func (r *Notice) UpdateNotice(ctx kratosx.Context, ent *entity.Notice) error {
	return ctx.DB().Updates(ent).Error
}

// DeleteNotice 删除通知信息数据
func (r *Notice) DeleteNotice(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(&entity.Notice{}, id).Error
}
