package repository

import (
	"partyaffairs/internal/core"
	"partyaffairs/internal/domain/entity"
	"partyaffairs/internal/types"
)

type InformationRepository interface {
	// ListInformationClassify 获取资讯分组列表
	ListInformationClassify(ctx core.Context) ([]*entity.InformationClassify, error)

	// CreateInformationClassify 创建资讯分组
	CreateInformationClassify(ctx core.Context, req *entity.InformationClassify) (uint32, error)

	// UpdateInformationClassify 更新资讯分组
	UpdateInformationClassify(ctx core.Context, req *entity.InformationClassify) error

	// DeleteInformationClassify 删除资讯分组
	DeleteInformationClassify(ctx core.Context, id uint32) error

	// GetInformation 获取指定的资讯信息
	GetInformation(ctx core.Context, id uint32) (*entity.Information, error)

	// ListInformation 获取资讯信息列表
	ListInformation(ctx core.Context, req *types.ListInformationRequest) ([]*entity.Information, uint32, error)

	// CreateInformation 创建资讯信息
	CreateInformation(ctx core.Context, req *entity.Information) (uint32, error)

	// UpdateInformation 更新资讯信息
	UpdateInformation(ctx core.Context, req *entity.Information) error

	// DeleteInformation 删除资讯信息
	DeleteInformation(ctx core.Context, id uint32) error

	// IncrReadCount 增加阅读次数
	IncrReadCount(ctx core.Context, id uint32) error
}
