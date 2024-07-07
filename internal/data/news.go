package data

import (
	"github.com/limes-cloud/kratosx"
	"gorm.io/gorm"

	"partyaffairs/internal/biz"
	"partyaffairs/internal/biz/types"
)

type newsRepo struct {
}

func NewNewsRepo() biz.NewsRepo {
	return &newsRepo{}
}

func (u *newsRepo) AllClassify(ctx kratosx.Context) ([]*biz.NewsClassify, error) {
	var list []*biz.NewsClassify
	return list, ctx.DB().Model(biz.NewsClassify{}).Order("weight desc").Find(&list).Error
}

func (u *newsRepo) AddClassify(ctx kratosx.Context, news *biz.NewsClassify) (uint32, error) {
	return news.Id, ctx.DB().Create(news).Error
}

func (u *newsRepo) UpdateClassify(ctx kratosx.Context, news *biz.NewsClassify) error {
	return ctx.DB().Where("id=?", news.Id).Updates(news).Error
}

func (u *newsRepo) DeleteClassify(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(biz.NewsClassify{}, "id=?", id).Error
}

func (u *newsRepo) PageContent(ctx kratosx.Context, req *types.PageNewsContentRequest) ([]*biz.NewsContent, uint32, error) {
	var list []*biz.NewsContent
	var total int64

	db := ctx.DB().Model(biz.NewsContent{}).
		Select("id", "title", "unit", "cover", "desc", "read", "classify_id", "is_top", "created_at").
		Preload("Classify")

	if req.ClassifyID != nil {
		db.Where("classify_id=?", *req.ClassifyID)
	}
	if req.Title != nil {
		db.Where("title like ?", *req.Title+"%")
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, uint32(total), err
	}

	return list, uint32(total), db.Offset(int((req.Page - 1) * req.PageSize)).Order("is_top desc").Limit(int(req.PageSize)).Find(&list).Error
}

func (u *newsRepo) GetContent(ctx kratosx.Context, id uint32) (*biz.NewsContent, error) {
	nc := biz.NewsContent{}
	return &nc, ctx.DB().Preload("Classify").Preload("Comments").First(&nc, "id=?", id).Error
}

func (u *newsRepo) AddContent(ctx kratosx.Context, news *biz.NewsContent) (uint32, error) {
	return news.Id, ctx.DB().Create(news).Error
}

func (u *newsRepo) AddContentReadCount(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Model(biz.NewsContent{}).Where("id=?", id).UpdateColumn("read", gorm.Expr("`read`+?", 1)).Error
}

func (u *newsRepo) UpdateContent(ctx kratosx.Context, news *biz.NewsContent) error {
	return ctx.DB().Where("id=?", news.Id).Updates(news).Error
}

func (u *newsRepo) DeleteContent(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(biz.NewsClassify{}, "id=?", id).Error
}

func (u *newsRepo) GetComment(ctx kratosx.Context, id uint32) (*biz.NewsComment, error) {
	nc := biz.NewsComment{}
	return &nc, ctx.DB().First(&nc, "id=?", id).Error
}

func (u *newsRepo) AddComment(ctx kratosx.Context, nc *biz.NewsComment) (uint32, error) {
	return nc.Id, ctx.DB().Create(nc).Error
}

func (u *newsRepo) DeleteComment(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(biz.NewsComment{}, "id=?", id).Error
}

func (u *newsRepo) PageComment(ctx kratosx.Context, req *types.PageNewsCommentRequest) ([]*biz.NewsComment, uint32, error) {
	var list []*biz.NewsComment
	var total int64
	db := ctx.DB().Model(biz.NewsComment{}).Where("content_id=?", req.ContentID)
	if req.Text != nil {
		db.Where("text like ?", "%"+*req.Text+"%")
	}
	if err := db.Count(&total).Error; err != nil {
		return nil, uint32(total), err
	}
	return list, uint32(total), db.Offset(int((req.Page - 1) * req.PageSize)).Order("id desc").Limit(int(req.PageSize)).Find(&list).Error
}
