package repository

import (
	"partyaffairs/internal/core"
	"partyaffairs/internal/domain/entity"
	"partyaffairs/internal/types"
)

type InterflowRepository interface {
	// ListInterflowClassify 获取资讯分组列表
	ListInterflowClassify(ctx core.Context) ([]*entity.InterflowClassify, error)

	// CreateInterflowClassify 创建资讯分组
	CreateInterflowClassify(ctx core.Context, req *entity.InterflowClassify) (uint32, error)

	// UpdateInterflowClassify 更新资讯分组
	UpdateInterflowClassify(ctx core.Context, req *entity.InterflowClassify) error

	// DeleteInterflowClassify 删除资讯分组
	DeleteInterflowClassify(ctx core.Context, id uint32) error

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
