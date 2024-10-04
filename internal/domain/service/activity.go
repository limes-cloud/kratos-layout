package service

import (
	"github.com/limes-cloud/kratosx"

	"partyaffairs/api/partyaffairs/errors"
	"partyaffairs/internal/conf"
	"partyaffairs/internal/domain/entity"
	"partyaffairs/internal/domain/repository"
	"partyaffairs/internal/types"
)

type ActivityService struct {
	conf *conf.Config
	repo repository.ActivityRepository
	file repository.FileRepository
}

func NewActivityService(conf *conf.Config, repo repository.ActivityRepository, file repository.FileRepository) *ActivityService {
	return &ActivityService{conf: conf, repo: repo, file: file}
}

// GetActivity 获取指定的活动信息
func (srv *ActivityService) GetActivity(ctx kratosx.Context, id uint32) (*entity.Activity, error) {
	ent, err := srv.repo.GetActivity(ctx, id)
	if err != nil {
		return nil, errors.GetError(err.Error())
	}
	ent.Cover = srv.file.GetFileURL(ctx, ent.Cover)
	return ent, nil
}

// ListActivity 获取活动信息列表
func (srv *ActivityService) ListActivity(ctx kratosx.Context, req *types.ListActivityRequest) ([]*entity.Activity, uint32, error) {
	list, total, err := srv.repo.ListActivity(ctx, req)
	if err != nil {
		return nil, 0, errors.ListError(err.Error())
	}
	for ind, item := range list {
		list[ind].Cover = srv.file.GetFileURL(ctx, item.Cover)
	}
	return list, total, nil
}

// CreateActivity 创建活动信息
func (srv *ActivityService) CreateActivity(ctx kratosx.Context, req *entity.Activity) (uint32, error) {
	id, err := srv.repo.CreateActivity(ctx, req)
	if err != nil {
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateActivity 更新活动信息
func (srv *ActivityService) UpdateActivity(ctx kratosx.Context, ent *entity.Activity) error {
	if err := srv.repo.UpdateActivity(ctx, ent); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteActivity 删除活动信息
func (srv *ActivityService) DeleteActivity(ctx kratosx.Context, id uint32) error {
	if err := srv.repo.DeleteActivity(ctx, id); err != nil {
		return errors.DeleteError(err.Error())
	}
	return nil
}
