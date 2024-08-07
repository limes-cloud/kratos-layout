package repository

import (
	"github.com/limes-cloud/kratosx"

	"poverty/internal/domain/entity"
	"poverty/internal/types"
)

type ResourceRepository interface {
	// GetResource 获取指定的文件信息
	GetResource(ctx kratosx.Context, id uint32) (*entity.Resource, error)

	// ListResource 获取文件信息列表
	ListResource(ctx kratosx.Context, req *types.ListResourceRequest) ([]*entity.Resource, uint32, error)

	// CreateResource 创建文件信息
	CreateResource(ctx kratosx.Context, req *entity.Resource) (uint32, error)

	// UpdateResource 更新文件信息
	UpdateResource(ctx kratosx.Context, req *entity.Resource) error

	// DeleteResource 删除文件信息
	DeleteResource(ctx kratosx.Context, id uint32) error
}
