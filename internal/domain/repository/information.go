package repository

import (
	"github.com/limes-cloud/kratosx"

	"partyaffairs/internal/domain/entity"
	"partyaffairs/internal/types"
)

type InformationRepository interface {
	// ListInformationClassify 获取资讯分组列表
	ListInformationClassify(ctx kratosx.Context) ([]*entity.InformationClassify, error)

	// CreateInformationClassify 创建资讯分组
	CreateInformationClassify(ctx kratosx.Context, req *entity.InformationClassify) (uint32, error)

	// UpdateInformationClassify 更新资讯分组
	UpdateInformationClassify(ctx kratosx.Context, req *entity.InformationClassify) error

	// DeleteInformationClassify 删除资讯分组
	DeleteInformationClassify(ctx kratosx.Context, id uint32) error

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
