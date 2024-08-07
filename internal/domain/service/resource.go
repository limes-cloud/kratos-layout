package service

import (
	"github.com/limes-cloud/kratosx"

	"poverty/api/poverty/errors"
	"poverty/internal/conf"
	"poverty/internal/domain/entity"
	"poverty/internal/domain/repository"
	"poverty/internal/types"
)

type ResourceService struct {
	conf *conf.Config
	repo repository.ResourceRepository
	file repository.FileRepository
}

func NewResourceService(conf *conf.Config, repo repository.ResourceRepository, file repository.FileRepository) *ResourceService {
	return &ResourceService{conf: conf, repo: repo, file: file}
}

// GetResource 获取指定的文件信息
func (srv *ResourceService) GetResource(ctx kratosx.Context, id uint32) (*entity.Resource, error) {
	ent, err := srv.repo.GetResource(ctx, id)
	if err != nil {
		return nil, errors.GetError(err.Error())
	}
	return ent, nil
}

// ListResource 获取文件信息列表
func (srv *ResourceService) ListResource(ctx kratosx.Context, req *types.ListResourceRequest) ([]*entity.Resource, uint32, error) {
	list, total, err := srv.repo.ListResource(ctx, req)
	if err != nil {
		return nil, 0, errors.ListError(err.Error())
	}
	for ind, item := range list {
		file, err := srv.file.GetFile(ctx, item.Src)
		if err != nil {
			continue
		}
		list[ind].File = file
	}
	return list, total, nil
}

// CreateResource 创建文件信息
func (srv *ResourceService) CreateResource(ctx kratosx.Context, req *entity.Resource) (uint32, error) {
	id, err := srv.repo.CreateResource(ctx, req)
	if err != nil {
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateResource 更新文件信息
func (srv *ResourceService) UpdateResource(ctx kratosx.Context, ent *entity.Resource) error {
	if err := srv.repo.UpdateResource(ctx, ent); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteResource 删除文件信息
func (srv *ResourceService) DeleteResource(ctx kratosx.Context, id uint32) error {
	if err := srv.repo.DeleteResource(ctx, id); err != nil {
		return errors.DeleteError(err.Error())
	}
	return nil
}
