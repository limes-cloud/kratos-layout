package repository

import (
	"github.com/limes-cloud/kratosx"

	"partyaffairs/internal/domain/entity"
	"partyaffairs/internal/types"
)

type ResourceRepository interface {
	// ListResourceClassify 获取资料分组列表
	ListResourceClassify(ctx kratosx.Context) ([]*entity.ResourceClassify, error)

	// CreateResourceClassify 创建资料分组
	CreateResourceClassify(ctx kratosx.Context, req *entity.ResourceClassify) (uint32, error)

	// UpdateResourceClassify 更新资料分组
	UpdateResourceClassify(ctx kratosx.Context, req *entity.ResourceClassify) error

	// DeleteResourceClassify 删除资料分组
	DeleteResourceClassify(ctx kratosx.Context, id uint32) error

	// GetResource 获取指定的资料信息
	GetResource(ctx kratosx.Context, id uint32) (*entity.Resource, error)

	// ListResource 获取资料信息列表
	ListResource(ctx kratosx.Context, req *types.ListResourceRequest) ([]*entity.Resource, uint32, error)

	// CreateResource 创建资料信息
	CreateResource(ctx kratosx.Context, req *entity.Resource) (uint32, error)

	// UpdateResource 更新资料信息
	UpdateResource(ctx kratosx.Context, req *entity.Resource) error

	// DeleteResource 删除资料信息
	DeleteResource(ctx kratosx.Context, id uint32) error
}
