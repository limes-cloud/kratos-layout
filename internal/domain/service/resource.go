package service

import (
	"github.com/limes-cloud/kratosx"

	"partyaffairs/api/partyaffairs/errors"
	"partyaffairs/internal/conf"
	"partyaffairs/internal/domain/entity"
	"partyaffairs/internal/domain/repository"
	"partyaffairs/internal/types"
)

type ResourceService struct {
	conf *conf.Config
	repo repository.ResourceRepository
	file repository.FileRepository
}

func NewResourceService(conf *conf.Config, repo repository.ResourceRepository, file repository.FileRepository) *ResourceService {
	return &ResourceService{conf: conf, repo: repo, file: file}
}

// ListResourceClassify 获取资讯分组列表
func (srv *ResourceService) ListResourceClassify(ctx kratosx.Context) ([]*entity.ResourceClassify, error) {
	list, err := srv.repo.ListResourceClassify(ctx)
	if err != nil {
		return nil, errors.ListError(err.Error())
	}
	return list, nil
}

// CreateResourceClassify 创建资讯分组
func (srv *ResourceService) CreateResourceClassify(ctx kratosx.Context, tg *entity.ResourceClassify) (uint32, error) {
	id, err := srv.repo.CreateResourceClassify(ctx, tg)
	if err != nil {
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateResourceClassify 更新资讯分组
func (srv *ResourceService) UpdateResourceClassify(ctx kratosx.Context, tg *entity.ResourceClassify) error {
	if err := srv.repo.UpdateResourceClassify(ctx, tg); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteResourceClassify 删除资讯分组
func (srv *ResourceService) DeleteResourceClassify(ctx kratosx.Context, id uint32) error {
	err := srv.repo.DeleteResourceClassify(ctx, id)
	if err != nil {
		return errors.DeleteError(err.Error())
	}
	return nil
}

// GetResource 获取指定的资讯信息
func (srv *ResourceService) GetResource(ctx kratosx.Context, id uint32) (*entity.Resource, error) {
	ent, err := srv.repo.GetResource(ctx, id)
	if err != nil {
		return nil, errors.GetError(err.Error())
	}
	ent.Url = srv.file.GetFileURL(ctx, ent.Url)
	return ent, nil
}

// ListResource 获取资讯信息列表
func (srv *ResourceService) ListResource(ctx kratosx.Context, req *types.ListResourceRequest) ([]*entity.Resource, uint32, error) {
	list, total, err := srv.repo.ListResource(ctx, req)
	if err != nil {
		return nil, 0, errors.ListError(err.Error())
	}
	if err != nil {
		return nil, 0, errors.ListError(err.Error())
	}
	for ind, item := range list {
		list[ind].Url = srv.file.GetFileURL(ctx, item.Url)
	}
	return list, total, nil
}

// CreateResource 创建资讯信息
func (srv *ResourceService) CreateResource(ctx kratosx.Context, req *entity.Resource) (uint32, error) {
	id, err := srv.repo.CreateResource(ctx, req)
	if err != nil {
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateResource 更新资讯信息
func (srv *ResourceService) UpdateResource(ctx kratosx.Context, ent *entity.Resource) error {
	if err := srv.repo.UpdateResource(ctx, ent); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteResource 删除资讯信息
func (srv *ResourceService) DeleteResource(ctx kratosx.Context, id uint32) error {
	if err := srv.repo.DeleteResource(ctx, id); err != nil {
		return errors.DeleteError(err.Error())
	}
	return nil
}
