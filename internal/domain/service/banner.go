package service

import (
	"github.com/limes-cloud/kratosx"

	"partyaffairs/api/partyaffairs/errors"
	"partyaffairs/internal/conf"
	"partyaffairs/internal/domain/entity"
	"partyaffairs/internal/domain/repository"
	"partyaffairs/internal/types"
)

type BannerService struct {
	conf *conf.Config
	repo repository.BannerRepository
	file repository.FileRepository
}

func NewBannerService(conf *conf.Config, repo repository.BannerRepository, file repository.FileRepository) *BannerService {
	return &BannerService{conf: conf, repo: repo, file: file}
}

// ListBanner 获取轮播图信息列表
func (srv *BannerService) ListBanner(ctx kratosx.Context, req *types.ListBannerRequest) ([]*entity.Banner, uint32, error) {
	list, total, err := srv.repo.ListBanner(ctx, req)
	if err != nil {
		return nil, 0, errors.ListError(err.Error())
	}
	for ind, item := range list {
		list[ind].Src = srv.file.GetFileURL(ctx, item.Src)
	}
	return list, total, nil
}

// CreateBanner 创建轮播图信息
func (srv *BannerService) CreateBanner(ctx kratosx.Context, req *entity.Banner) (uint32, error) {
	id, err := srv.repo.CreateBanner(ctx, req)
	if err != nil {
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateBanner 更新轮播图信息
func (srv *BannerService) UpdateBanner(ctx kratosx.Context, ent *entity.Banner) error {
	if err := srv.repo.UpdateBanner(ctx, ent); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteBanner 删除轮播图信息
func (srv *BannerService) DeleteBanner(ctx kratosx.Context, id uint32) error {
	if err := srv.repo.DeleteBanner(ctx, id); err != nil {
		return errors.DeleteError(err.Error())
	}
	return nil
}
