package repository

import (
	"github.com/limes-cloud/kratosx"

	"layout/internal/domain/entity"
	"layout/internal/types"
)

type Dictionary interface {
	// GetDictionary 获取指定的字典目录
	GetDictionary(ctx kratosx.Context, id uint32) (*entity.Dictionary, error)

	// ListDictionary 获取字典目录列表
	ListDictionary(ctx kratosx.Context, req *types.ListDictionaryRequest) ([]*entity.Dictionary, uint32, error)

	// CreateDictionary 创建字典目录
	CreateDictionary(ctx kratosx.Context, req *entity.Dictionary) (uint32, error)

	// UpdateDictionary 更新字典目录
	UpdateDictionary(ctx kratosx.Context, req *entity.Dictionary) error

	// DeleteDictionary 删除字典目录
	DeleteDictionary(ctx kratosx.Context, id uint32) error

	// GetDictionaryByKeyword 获取指定的字典目录
	GetDictionaryByKeyword(ctx kratosx.Context, keyword string) (*entity.Dictionary, error)
}
