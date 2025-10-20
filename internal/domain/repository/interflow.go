package repository

import (
	"partyaffairs/internal/core"
	"partyaffairs/internal/domain/entity"
	"partyaffairs/internal/types"
)

type InterflowRepository interface {
	// ListInterflowPerson 获取通知信息列表
	ListInterflowPerson(ctx core.Context) ([]*entity.InterflowPerson, error)

	// CreateInterflowPerson 创建通知信息
	CreateInterflowPerson(ctx core.Context, req *entity.InterflowPerson) (uint32, error)

	// DeleteInterflowPerson 删除通知信息
	DeleteInterflowPerson(ctx core.Context, id uint32) error

	ListInterflowFormUserIds(ctx core.Context, uid uint32) ([]uint32, error)

	// ListInterflow 获取通知信息列表
	ListInterflow(ctx core.Context, req *types.ListInterflowRequest) ([]*entity.Interflow, uint32, error)

	// CreateInterflow 创建通知信息
	CreateInterflow(ctx core.Context, req *entity.Interflow) (uint32, error)

	// UpdateInterflow 更新通知信息
	UpdateInterflow(ctx core.Context, req *entity.Interflow) error

	// DeleteInterflow 删除通知信息
	DeleteInterflow(ctx core.Context, id uint32) error
}
