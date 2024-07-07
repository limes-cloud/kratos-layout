package data

import (
	"github.com/limes-cloud/kratosx"

	"partyaffairs/internal/biz"
	"partyaffairs/internal/biz/types"
)

type resourceRepo struct {
}

func NewResourceRepo() biz.ResourceRepo {
	return &resourceRepo{}
}

func (u *resourceRepo) AllClassify(ctx kratosx.Context) ([]*biz.ResourceClassify, error) {
	var list []*biz.ResourceClassify
	return list, ctx.DB().Model(biz.ResourceClassify{}).Order("weight desc").Find(&list).Error
}

func (u *resourceRepo) AddClassify(ctx kratosx.Context, resource *biz.ResourceClassify) (uint32, error) {
	return resource.Id, ctx.DB().Create(resource).Error
}

func (u *resourceRepo) UpdateClassify(ctx kratosx.Context, resource *biz.ResourceClassify) error {
	return ctx.DB().Where("id=?", resource.Id).Updates(resource).Error
}

func (u *resourceRepo) DeleteClassify(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(biz.ResourceClassify{}, "id=?", id).Error
}

func (u *resourceRepo) PageContent(ctx kratosx.Context, req *types.PageResourceContentRequest) ([]*biz.ResourceContent, uint32, error) {
	var list []*biz.ResourceContent
	var total int64

	db := ctx.DB().Model(biz.ResourceContent{}).Preload("ResourceClassify")

	if req.ClassifyID != nil && *req.ClassifyID != 0 {
		db.Where("classify_id=?", *req.ClassifyID)
	}
	if req.Title != nil {
		db.Where("title like ?", *req.Title+"%")
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, uint32(total), err
	}

	return list, uint32(total), db.Offset(int((req.Page - 1) * req.PageSize)).Limit(int(req.PageSize)).Find(&list).Error
}

func (u *resourceRepo) GetContent(ctx kratosx.Context, id uint32) (*biz.ResourceContent, error) {
	nc := biz.ResourceContent{}
	return &nc, ctx.DB().First(&nc, "id=?", id).Error
}

func (u *resourceRepo) AddContent(ctx kratosx.Context, resource *biz.ResourceContent) (uint32, error) {
	return resource.Id, ctx.DB().Create(resource).Error
}

func (u *resourceRepo) UpdateContent(ctx kratosx.Context, resource *biz.ResourceContent) error {
	return ctx.DB().Where("id=?", resource.Id).Updates(resource).Error
}

func (u *resourceRepo) DeleteContent(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(biz.ResourceClassify{}, "id=?", id).Error
}
