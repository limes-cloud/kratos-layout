package service

import (
	"partyaffairs/api/errors"
	"partyaffairs/internal/core"
	"partyaffairs/internal/domain/entity"
	"partyaffairs/internal/domain/repository"
	"partyaffairs/internal/types"
)

type PublicService struct {
	repo repository.PublicRepository
}

func NewPublicService(repo repository.PublicRepository, file repository.FileRepository) *PublicService {
	return &PublicService{repo: repo}
}

// ListPublicClassify 获取资讯分组列表
func (srv *PublicService) ListPublicClassify(ctx core.Context) ([]*entity.PublicClassify, error) {
	list, err := srv.repo.ListPublicClassify(ctx)
	if err != nil {
		return nil, errors.ListError(err.Error())
	}
	return list, nil
}

// CreatePublicClassify 创建资讯分组
func (srv *PublicService) CreatePublicClassify(ctx core.Context, tg *entity.PublicClassify) (uint32, error) {
	id, err := srv.repo.CreatePublicClassify(ctx, tg)
	if err != nil {
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdatePublicClassify 更新资讯分组
func (srv *PublicService) UpdatePublicClassify(ctx core.Context, tg *entity.PublicClassify) error {
	if err := srv.repo.UpdatePublicClassify(ctx, tg); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeletePublicClassify 删除资讯分组
func (srv *PublicService) DeletePublicClassify(ctx core.Context, id uint32) error {
	err := srv.repo.DeletePublicClassify(ctx, id)
	if err != nil {
		return errors.DeleteError(err.Error())
	}
	return nil
}

// GetPublic 获取指定的资讯信息
func (srv *PublicService) GetPublic(ctx core.Context, id uint32) (*entity.Public, error) {
	ent, err := srv.repo.GetPublic(ctx, id)
	if err != nil {
		return nil, errors.GetError(err.Error())
	}
	_ = srv.repo.IncrReadCount(ctx, id)
	return ent, nil
}

// ListPublic 获取资讯信息列表
func (srv *PublicService) ListPublic(ctx core.Context, req *types.ListPublicRequest) ([]*entity.Public, uint32, error) {
	list, total, err := srv.repo.ListPublic(ctx, req)
	if err != nil {
		return nil, 0, errors.ListError(err.Error())
	}
	if err != nil {
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}

// CreatePublic 创建资讯信息
func (srv *PublicService) CreatePublic(ctx core.Context, req *entity.Public) (uint32, error) {
	id, err := srv.repo.CreatePublic(ctx, req)
	if err != nil {
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdatePublic 更新资讯信息
func (srv *PublicService) UpdatePublic(ctx core.Context, ent *entity.Public) error {
	if err := srv.repo.UpdatePublic(ctx, ent); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeletePublic 删除资讯信息
func (srv *PublicService) DeletePublic(ctx core.Context, id uint32) error {
	if err := srv.repo.DeletePublic(ctx, id); err != nil {
		return errors.DeleteError(err.Error())
	}
	return nil
}
