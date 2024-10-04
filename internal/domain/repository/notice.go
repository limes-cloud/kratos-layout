package repository

import (
	"github.com/limes-cloud/kratosx"

	"partyaffairs/internal/domain/entity"
	"partyaffairs/internal/types"
)

type NoticeRepository interface {
	// GetNotice 获取指定的通知信息
	GetNotice(ctx kratosx.Context, id uint32) (*entity.Notice, error)

	// ListNotice 获取通知信息列表
	ListNotice(ctx kratosx.Context, req *types.ListNoticeRequest) ([]*entity.Notice, uint32, error)

	// CreateNotice 创建通知信息
	CreateNotice(ctx kratosx.Context, req *entity.Notice) (uint32, error)

	// UpdateNotice 更新通知信息
	UpdateNotice(ctx kratosx.Context, req *entity.Notice) error

	// DeleteNotice 删除通知信息
	DeleteNotice(ctx kratosx.Context, id uint32) error
}
