package dbs

import (
	"github.com/limes-cloud/kratosx/model/page"
	"github.com/limes-cloud/kratosx/pkg/value"
	"partyaffairs/internal/core"
	"partyaffairs/internal/domain/entity"
	"partyaffairs/internal/types"
)

type Notice struct {
}

func NewNotice() *Notice {
	return &Notice{}
}

// GetNotice 获取指定的通知信息数据
func (r *Notice) GetNotice(ctx core.Context, id uint32) (*entity.Notice, error) {
	var (
		ent = entity.Notice{}
		fs  = []string{"*"}
	)
	db := ctx.DB().Select(fs)
	return &ent, db.First(&ent, id).Error
}

// ReadUserNotice 读取用户通知信息
func (r *Notice) ReadUserNotice(ctx core.Context, uid, nid uint32) error {
	return ctx.DB().Model(&entity.NoticeUser{}).
		Where("notice_id = ?", nid).
		Where("user_id = ?", uid).
		Update("is_read", 1).Error
}

// ListNotice 获取通知信息列表
func (r *Notice) ListNotice(ctx core.Context, req *types.ListNoticeRequest) ([]*entity.Notice, uint32, error) {
	var (
		list  []*entity.Notice
		total int64
	)

	db := ctx.DB().Model(entity.Notice{})
	if req.NotRead != nil && *req.NotRead {
		sql := "left join notice_user on notice_user.notice_id = notice.id and notice_user.user_id = ?"
		db = db.Joins(sql, ctx.Auth().UserId)
		db = db.Where("notice_user.is_read = 0")
	}
	if req.Title != nil {
		db = db.Where("title LIKE ?", *req.Title+"%")
	}
	if req.IsTop != nil {
		db = db.Where("is_top = ?", *req.IsTop)
	}
	if req.Status != nil {
		db = db.Where("status = ?", *req.Status)
	}

	if req.Search != nil {
		req.Order = value.Pointer("desc")
		req.OrderBy = value.Pointer("is_top")
		if err := db.Count(&total).Error; err != nil {
			return nil, 0, err
		}

		db = page.SearchScopes(db, req.Search)
	}

	return list, uint32(total), db.Find(&list).Error
}

// CreateNotice 创建通知信息数据
func (r *Notice) CreateNotice(ctx core.Context, ent *entity.Notice) (uint32, error) {
	return ent.Id, ctx.DB().Create(ent).Error
}

// UpdateNotice 更新通知信息数据
func (r *Notice) UpdateNotice(ctx core.Context, ent *entity.Notice) error {
	return ctx.DB().Updates(ent).Error
}

// DeleteNotice 删除通知信息数据
func (r *Notice) DeleteNotice(ctx core.Context, id uint32) error {
	return ctx.DB().Delete(&entity.Notice{}, id).Error
}
