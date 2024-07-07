package data

import (
	"github.com/limes-cloud/kratosx"

	"partyaffairs/internal/biz"
	"partyaffairs/internal/biz/types"
)

type noticeRepo struct {
}

func NewNoticeRepo() biz.NoticeRepo {
	return &noticeRepo{}
}

func (u *noticeRepo) Get(ctx kratosx.Context, id uint32) (*biz.Notice, error) {
	nc := biz.Notice{}
	return &nc, ctx.DB().First(&nc, "id=?", id).Error
}

func (u *noticeRepo) ReadUserIds(ctx kratosx.Context, nid uint32) ([]uint32, error) {
	var ids []uint32
	return ids, ctx.DB().Model(biz.NoticeUser{}).Select("user_id").Where("notice_id=?", nid).Scan(&ids).Error
}

func (u *noticeRepo) Page(ctx kratosx.Context, req *types.PageNoticeRequest) ([]*biz.Notice, uint32, error) {
	var list []*biz.Notice
	var total int64

	db := ctx.DB().Model(biz.Notice{})
	db = db.Select("notice.id", "title", "unit", "desc", "is_top", "notice.created_at", "notice.updated_at")
	if req.UserID != nil {
		db = db.Preload("NoticeUser", ctx.DB().Where("user_id=?", req.UserID))
		db = db.Joins("NoticeUser", ctx.DB().Where("user_id=?", req.UserID))
		if req.NotRead != nil && *req.NotRead {
			db = db.Where("user_id is null")
		}
	}

	if req.Title != nil {
		db.Where("title like ?", *req.Title+"%")
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, uint32(total), err
	}

	return list, uint32(total), db.Offset(int((req.Page - 1) * req.PageSize)).Order("is_top desc").Limit(int(req.PageSize)).Find(&list).Error
}

func (u *noticeRepo) Create(ctx kratosx.Context, notice *biz.Notice) (uint32, error) {
	return notice.Id, ctx.DB().Create(notice).Error
}

func (u *noticeRepo) AddNoticeUser(ctx kratosx.Context, nu *biz.NoticeUser) error {
	return ctx.DB().Create(nu).Error
}

func (u *noticeRepo) Update(ctx kratosx.Context, notice *biz.Notice) error {
	return ctx.DB().Where("id=?", notice.Id).Updates(notice).Error
}

func (u *noticeRepo) Delete(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(biz.Notice{}, "id=?", id).Error
}
