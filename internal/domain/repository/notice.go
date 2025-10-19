package repository

import (
	"partyaffairs/internal/core"
	"partyaffairs/internal/domain/entity"
	"partyaffairs/internal/types"
)

type NoticeRepository interface {
	ReadUserNotice(ctx core.Context, uid, nid uint32) error

	// GetNotice 获取指定的通知信息
	GetNotice(ctx core.Context, id uint32) (*entity.Notice, error)

	// ListNotice 获取通知信息列表
	ListNotice(ctx core.Context, req *types.ListNoticeRequest) ([]*entity.Notice, uint32, error)

	// CreateNotice 创建通知信息
	CreateNotice(ctx core.Context, req *entity.Notice) (uint32, error)

	// UpdateNotice 更新通知信息
	UpdateNotice(ctx core.Context, req *entity.Notice) error

	// DeleteNotice 删除通知信息
	DeleteNotice(ctx core.Context, id uint32) error
}
