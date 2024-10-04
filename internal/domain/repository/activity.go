package repository

import (
	"github.com/limes-cloud/kratosx"

	"partyaffairs/internal/domain/entity"
	"partyaffairs/internal/types"
)

type ActivityRepository interface {
	// GetActivity 获取指定的活动信息
	GetActivity(ctx kratosx.Context, id uint32) (*entity.Activity, error)

	// ListActivity 获取活动信息列表
	ListActivity(ctx kratosx.Context, req *types.ListActivityRequest) ([]*entity.Activity, uint32, error)

	// CreateActivity 创建活动信息
	CreateActivity(ctx kratosx.Context, req *entity.Activity) (uint32, error)

	// UpdateActivity 更新活动信息
	UpdateActivity(ctx kratosx.Context, req *entity.Activity) error

	// DeleteActivity 删除活动信息
	DeleteActivity(ctx kratosx.Context, id uint32) error
}
