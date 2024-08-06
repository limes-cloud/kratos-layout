package repository

import (
	"github.com/limes-cloud/kratosx"

	"poverty/internal/domain/entity"
	"poverty/internal/types"
)

type InformationRepository interface {
	// GetInformation 获取指定的资讯信息
	GetInformation(ctx kratosx.Context, id uint32) (*entity.Information, error)

	// ListInformation 获取资讯信息列表
	ListInformation(ctx kratosx.Context, req *types.ListInformationRequest) ([]*entity.Information, uint32, error)

	// CreateInformation 创建资讯信息
	CreateInformation(ctx kratosx.Context, req *entity.Information) (uint32, error)

	// UpdateInformation 更新资讯信息
	UpdateInformation(ctx kratosx.Context, req *entity.Information) error

	// DeleteInformation 删除资讯信息
	DeleteInformation(ctx kratosx.Context, id uint32) error
}
