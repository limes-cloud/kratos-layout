package service

import (
	"github.com/limes-cloud/kratosx"

	"poverty/api/poverty/errors"
	"poverty/internal/conf"
	"poverty/internal/domain/entity"
	"poverty/internal/domain/repository"
	"poverty/internal/types"
)

type NoticeService struct {
	conf *conf.Config
	repo repository.NoticeRepository
}

func NewNoticeService(conf *conf.Config, repo repository.NoticeRepository) *NoticeService {
	return &NoticeService{conf: conf, repo: repo}
}

// GetNotice 获取指定的通知信息
func (srv *NoticeService) GetNotice(ctx kratosx.Context, id uint32) (*entity.Notice, error) {
	ent, err := srv.repo.GetNotice(ctx, id)
	if err != nil {
		return nil, errors.GetError(err.Error())
	}
	return ent, nil
}

// ListNotice 获取通知信息列表
func (srv *NoticeService) ListNotice(ctx kratosx.Context, req *types.ListNoticeRequest) ([]*entity.Notice, uint32, error) {
	list, total, err := srv.repo.ListNotice(ctx, req)
	if err != nil {
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}

// CreateNotice 创建通知信息
func (srv *NoticeService) CreateNotice(ctx kratosx.Context, req *entity.Notice) (uint32, error) {
	id, err := srv.repo.CreateNotice(ctx, req)
	if err != nil {
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateNotice 更新通知信息
func (srv *NoticeService) UpdateNotice(ctx kratosx.Context, ent *entity.Notice) error {
	if err := srv.repo.UpdateNotice(ctx, ent); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteNotice 删除通知信息
func (srv *NoticeService) DeleteNotice(ctx kratosx.Context, id uint32) error {
	if err := srv.repo.DeleteNotice(ctx, id); err != nil {
		return errors.DeleteError(err.Error())
	}
	return nil
}
