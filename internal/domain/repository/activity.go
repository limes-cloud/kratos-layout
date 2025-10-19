package repository

import (
	"partyaffairs/internal/core"
	"partyaffairs/internal/domain/entity"
	"partyaffairs/internal/types"
)

type ActivityRepository interface {
	// GetActivity 获取指定的活动信息
	GetActivity(ctx core.Context, id uint32) (*entity.Activity, error)

	// ListActivity 获取活动信息列表
	ListActivity(ctx core.Context, req *types.ListActivityRequest) ([]*entity.Activity, uint32, error)

	// CreateActivity 创建活动信息
	CreateActivity(ctx core.Context, req *entity.Activity) (uint32, error)

	// UpdateActivity 更新活动信息
	UpdateActivity(ctx core.Context, req *entity.Activity) error

	// DeleteActivity 删除活动信息
	DeleteActivity(ctx core.Context, id uint32) error
}
