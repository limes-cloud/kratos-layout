package data

import (
	"github.com/limes-cloud/kratosx"

	"partyaffairs/internal/biz"
	"partyaffairs/internal/biz/types"
)

type videoRepo struct {
}

func NewVideoRepo() biz.VideoRepo {
	return &videoRepo{}
}

func (u *videoRepo) PageClassify(ctx kratosx.Context, req *types.PageVideoClassifyRequest) ([]*biz.VideoClassify, uint32, error) {
	var list []*biz.VideoClassify
	var total int64

	db := ctx.DB().Model(biz.VideoClassify{})
	if req.Name != nil {
		db.Where("name like ?", *req.Name+"%")
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, uint32(total), err
	}

	return list, uint32(total), db.Offset(int((req.Page - 1) * req.PageSize)).Order("weight desc").Limit(int(req.PageSize)).Find(&list).Error
}

func (u *videoRepo) AddClassify(ctx kratosx.Context, video *biz.VideoClassify) (uint32, error) {
	return video.Id, ctx.DB().Create(video).Error
}

func (u *videoRepo) UpdateClassify(ctx kratosx.Context, video *biz.VideoClassify) error {
	return ctx.DB().Where("id=?", video.Id).Updates(video).Error
}

func (u *videoRepo) DeleteClassify(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(biz.VideoClassify{}, "id=?", id).Error
}

func (u *videoRepo) PageContent(ctx kratosx.Context, req *types.PageVideoContentRequest) ([]*biz.VideoContent, uint32, error) {
	var list []*biz.VideoContent
	var total int64

	db := ctx.DB().Model(biz.VideoContent{}).Preload("Classify")
	if req.UserID != nil {
		db = db.Preload("Process", ctx.DB().Where("user_id=?", *req.UserID))
	}

	if req.ClassifyID != nil && *req.ClassifyID != 0 {
		db.Where("classify_id=?", *req.ClassifyID)
	}
	if req.Title != nil {
		db.Where("title like ?", *req.Title+"%")
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, uint32(total), err
	}

	return list, uint32(total), db.Offset(int((req.Page - 1) * req.PageSize)).Order("weight asc").Limit(int(req.PageSize)).Find(&list).Error
}

func (u *videoRepo) GetContent(ctx kratosx.Context, id uint32) (*biz.VideoContent, error) {
	nc := biz.VideoContent{}
	return &nc, ctx.DB().First(&nc, "id=?", id).Error
}

func (u *videoRepo) AddContent(ctx kratosx.Context, video *biz.VideoContent) (uint32, error) {
	return video.Id, ctx.DB().Create(video).Error
}

func (u *videoRepo) UpdateContent(ctx kratosx.Context, video *biz.VideoContent) error {
	return ctx.DB().Where("id=?", video.Id).Updates(video).Error
}

func (u *videoRepo) DeleteContent(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(biz.VideoContent{}, "id=?", id).Error
}

func (u *videoRepo) UpdateUserVideo(ctx kratosx.Context, video *biz.UserVideoProcess) error {
	return ctx.DB().Updates(video).Error
}

func (u *videoRepo) AddUserVideo(ctx kratosx.Context, video *biz.UserVideoProcess) (uint32, error) {
	return video.Id, ctx.DB().Create(video).Error
}

func (u *videoRepo) GetUserVideo(ctx kratosx.Context, vid, uid uint32) (*biz.UserVideoProcess, error) {
	nc := biz.UserVideoProcess{}
	return &nc, ctx.DB().First(&nc, "video_id=? and user_id=?", vid, uid).Error
}
