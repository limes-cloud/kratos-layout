package data

import (
	"github.com/limes-cloud/kratosx"

	"partyaffairs/internal/biz"
)

type bannerRepo struct {
}

func NewBannerRepo() biz.BannerRepo {
	return &bannerRepo{}
}

func (u *bannerRepo) All(ctx kratosx.Context) ([]*biz.Banner, error) {
	var list []*biz.Banner
	return list, ctx.DB().Model(biz.Banner{}).Order("weight desc").Find(&list).Error
}

func (u *bannerRepo) Create(ctx kratosx.Context, banner *biz.Banner) (uint32, error) {
	return banner.Id, ctx.DB().Create(banner).Error
}

func (u *bannerRepo) Update(ctx kratosx.Context, banner *biz.Banner) error {
	return ctx.DB().Where("id=?", banner.Id).Updates(banner).Error
}

func (u *bannerRepo) Delete(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(biz.Banner{}, "id=?", id).Error
}
