package repository

import (
	"github.com/limes-cloud/kratosx"

	"tianji/internal/domain/entity"
	"tianji/internal/types"
)

type Entity interface {
	GetEntity(ctx kratosx.Context, id uint32) (*entity.Entity, error)
	GetEntityByName(ctx kratosx.Context, appId uint32, db string, name string) (*entity.Entity, error)
	ListEntity(ctx kratosx.Context, req *types.ListEntityRequest) ([]*entity.Entity, uint32, error)
	CreateEntity(ctx kratosx.Context, req *entity.Entity) (uint32, error)
	UpdateEntity(ctx kratosx.Context, req *entity.Entity) error
	UpdateRegisteredEntity(ctx kratosx.Context, req *entity.Entity) error
	DeleteEntity(ctx kratosx.Context, id uint32) error
	DeleteEntities(ctx kratosx.Context, ids []uint32) error

	GetEntityField(ctx kratosx.Context, entityId uint32, name string) (*entity.EntityField, error)
	ListEntityField(ctx kratosx.Context, req *types.ListEntityFieldRequest) ([]*entity.EntityField, uint32, error)
	CreateEntityField(ctx kratosx.Context, req *entity.EntityField) (uint32, error)
	UpdateEntityField(ctx kratosx.Context, req *entity.EntityField) error
	UpdateRegisteredEntityField(ctx kratosx.Context, req *entity.EntityField) error
	DeleteEntityField(ctx kratosx.Context, id uint32) error
	DeleteEntityFields(ctx kratosx.Context, ids []uint32) error

	GetEntityRule(ctx kratosx.Context, entityId uint32, name string) (*entity.EntityRule, error)
	ListEntityRule(ctx kratosx.Context, req *types.ListEntityRuleRequest) ([]*entity.EntityRule, uint32, error)
	CreateEntityRule(ctx kratosx.Context, req *entity.EntityRule) (uint32, error)
	UpdateEntityRule(ctx kratosx.Context, req *entity.EntityRule) error
	UpdateRegisteredEntityRule(ctx kratosx.Context, req *entity.EntityRule) error
	DeleteEntityRule(ctx kratosx.Context, id uint32) error
	DeleteEntityRules(ctx kratosx.Context, ids []uint32) error
}
