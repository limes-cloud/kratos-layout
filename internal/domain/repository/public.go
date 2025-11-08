package repository

import (
	"partyaffairs/internal/core"
	"partyaffairs/internal/domain/entity"
	"partyaffairs/internal/types"
)

type PublicRepository interface {
	// ListPublicClassify 获取资讯分组列表
	ListPublicClassify(ctx core.Context) ([]*entity.PublicClassify, error)

	// CreatePublicClassify 创建资讯分组
	CreatePublicClassify(ctx core.Context, req *entity.PublicClassify) (uint32, error)

	// UpdatePublicClassify 更新资讯分组
	UpdatePublicClassify(ctx core.Context, req *entity.PublicClassify) error

	// DeletePublicClassify 删除资讯分组
	DeletePublicClassify(ctx core.Context, id uint32) error

	// GetPublic 获取指定的资讯信息
	GetPublic(ctx core.Context, id uint32) (*entity.Public, error)

	// ListPublic 获取资讯信息列表
	ListPublic(ctx core.Context, req *types.ListPublicRequest) ([]*entity.Public, uint32, error)

	// CreatePublic 创建资讯信息
	CreatePublic(ctx core.Context, req *entity.Public) (uint32, error)

	// UpdatePublic 更新资讯信息
	UpdatePublic(ctx core.Context, req *entity.Public) error

	// DeletePublic 删除资讯信息
	DeletePublic(ctx core.Context, id uint32) error

	// IncrReadCount 增加阅读次数
	IncrReadCount(ctx core.Context, id uint32) error
}
