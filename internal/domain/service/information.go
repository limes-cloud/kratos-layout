package service

import (
	"partyaffairs/api/errors"
	"partyaffairs/internal/core"
	"partyaffairs/internal/domain/entity"
	"partyaffairs/internal/domain/repository"
	"partyaffairs/internal/types"
)

type InformationService struct {
	repo repository.InformationRepository
}

func NewInformationService(repo repository.InformationRepository, file repository.FileRepository) *InformationService {
	return &InformationService{repo: repo}
}

// ListInformationClassify 获取资讯分组列表
func (srv *InformationService) ListInformationClassify(ctx core.Context) ([]*entity.InformationClassify, error) {
	list, err := srv.repo.ListInformationClassify(ctx)
	if err != nil {
		return nil, errors.ListError(err.Error())
	}
	return list, nil
}

// CreateInformationClassify 创建资讯分组
func (srv *InformationService) CreateInformationClassify(ctx core.Context, tg *entity.InformationClassify) (uint32, error) {
	id, err := srv.repo.CreateInformationClassify(ctx, tg)
	if err != nil {
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateInformationClassify 更新资讯分组
func (srv *InformationService) UpdateInformationClassify(ctx core.Context, tg *entity.InformationClassify) error {
	if err := srv.repo.UpdateInformationClassify(ctx, tg); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteInformationClassify 删除资讯分组
func (srv *InformationService) DeleteInformationClassify(ctx core.Context, id uint32) error {
	err := srv.repo.DeleteInformationClassify(ctx, id)
	if err != nil {
		return errors.DeleteError(err.Error())
	}
	return nil
}

// GetInformation 获取指定的资讯信息
func (srv *InformationService) GetInformation(ctx core.Context, id uint32) (*entity.Information, error) {
	ent, err := srv.repo.GetInformation(ctx, id)
	if err != nil {
		return nil, errors.GetError(err.Error())
	}
	_ = srv.repo.IncrReadCount(ctx, id)
	return ent, nil
}

// ListInformation 获取资讯信息列表
func (srv *InformationService) ListInformation(ctx core.Context, req *types.ListInformationRequest) ([]*entity.Information, uint32, error) {
	list, total, err := srv.repo.ListInformation(ctx, req)
	if err != nil {
		return nil, 0, errors.ListError(err.Error())
	}
	if err != nil {
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}

// CreateInformation 创建资讯信息
func (srv *InformationService) CreateInformation(ctx core.Context, req *entity.Information) (uint32, error) {
	id, err := srv.repo.CreateInformation(ctx, req)
	if err != nil {
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateInformation 更新资讯信息
func (srv *InformationService) UpdateInformation(ctx core.Context, ent *entity.Information) error {
	if err := srv.repo.UpdateInformation(ctx, ent); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteInformation 删除资讯信息
func (srv *InformationService) DeleteInformation(ctx core.Context, id uint32) error {
	if err := srv.repo.DeleteInformation(ctx, id); err != nil {
		return errors.DeleteError(err.Error())
	}
	return nil
}
