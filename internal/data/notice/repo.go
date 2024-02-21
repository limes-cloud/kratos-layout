package notice

import (
	"github.com/limes-cloud/kratosx"

	bizNotice "github.com/go-kratos/kratos-layout/internal/biz/notice"
)

type repo struct {
}

func NewRepo() bizNotice.Repo {
	return &repo{}
}

func (u repo) Create(ctx kratosx.Context, notice *bizNotice.Notice) (uint32, error) {
	return notice.ID, ctx.DB().Create(notice).Error
}

func (u repo) Update(ctx kratosx.Context, notice *bizNotice.Notice) error {
	return ctx.DB().Model(bizNotice.Notice{}).Updates(notice).Error
}

func (u repo) Delete(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(bizNotice.Notice{}, "id=?", id).Error
}

func (u repo) Get(ctx kratosx.Context, id uint32) (*bizNotice.Notice, error) {
	var ins bizNotice.Notice
	return &ins, ctx.DB().First(&ins, "id=?", id).Error
}

func (u repo) Page(ctx kratosx.Context, req *bizNotice.PageRequest) ([]*bizNotice.Notice, uint32, error) {
	var list []*bizNotice.Notice
	var total int64

	db := ctx.DB().Model(bizNotice.Notice{})
	if req.Email != nil {
		db = db.Where("email=?", *req.Email)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, uint32(total), err
	}

	db = db.Offset(int((req.Page - 1) * req.PageSize)).Limit(int(req.PageSize))

	return list, uint32(total), db.Find(&list).Error
}

func (u repo) PageNoticeUser(ctx kratosx.Context, req *bizNotice.PageNoticeUserRequest) ([]*bizNotice.NoticeUser, uint32, error) {
	var list []*bizNotice.NoticeUser
	var total int64

	db := ctx.DB().Model(bizNotice.NoticeUser{}).Where("notice_id=?", req.NoticeID).Preload("User")

	if err := db.Count(&total).Error; err != nil {
		return nil, uint32(total), err
	}

	db = db.Offset(int((req.Page - 1) * req.PageSize)).Limit(int(req.PageSize))

	return list, uint32(total), db.Find(&list).Error
}
