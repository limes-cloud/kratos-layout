package dbs

import (
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/model/page"

	"tianji/internal/domain/entity"
	"tianji/internal/types"
)

type Entity struct{}

func NewEntity() *Entity {
	return &Entity{}
}

func (infra *Entity) GetEntity(ctx kratosx.Context, id uint32) (*entity.Entity, error) {
	var ent entity.Entity
	return &ent, ctx.DB().First(&ent, id).Error
}

func (infra *Entity) GetEntityByName(ctx kratosx.Context, appId uint32, db string, name string) (*entity.Entity, error) {
	var ent entity.Entity
	return &ent, ctx.DB().
		Where("app_id = ? and `database` = ? and name = ?", appId, db, name).
		First(&ent).Error
}

func (infra *Entity) ListEntity(ctx kratosx.Context, req *types.ListEntityRequest) ([]*entity.Entity, uint32, error) {
	var (
		list  []*entity.Entity
		total int64
	)

	db := ctx.DB().Model(entity.Entity{})
	if req.AppId != nil {
		db = db.Where("app_id = ?", *req.AppId)
	}
	if req.Name != nil {
		db = db.Where("name LIKE ?", *req.Name+"%")
	}
	if req.Database != nil {
		db = db.Where("`database` = ?", *req.Database)
	}
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	db = page.SearchScopes(db, &req.Search)
	return list, uint32(total), db.Find(&list).Error
}

func (infra *Entity) CreateEntity(ctx kratosx.Context, req *entity.Entity) (uint32, error) {
	if err := ctx.DB().Create(req).Error; err != nil {
		return 0, err
	}
	return req.Id, nil
}

func (infra *Entity) UpdateEntity(ctx kratosx.Context, req *entity.Entity) error {
	return ctx.DB().Updates(req).Error
}

func (infra *Entity) UpdateRegisteredEntity(ctx kratosx.Context, req *entity.Entity) error {
	return ctx.DB().Model(&entity.Entity{}).
		Where("id = ?", req.Id).
		Omit("app_id", "database", "name").
		Updates(req).Error
}

func (infra *Entity) DeleteEntity(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Where("id = ?", id).Delete(&entity.Entity{}).Error
}

func (infra *Entity) DeleteEntities(ctx kratosx.Context, ids []uint32) error {
	if len(ids) == 0 {
		return nil
	}
	return ctx.DB().Where("id in ?", ids).Delete(&entity.Entity{}).Error
}

func (infra *Entity) GetEntityField(ctx kratosx.Context, entityId uint32, name string) (*entity.EntityField, error) {
	var ent entity.EntityField
	return &ent, ctx.DB().
		Where("entity_id = ? and name = ?", entityId, name).
		First(&ent).Error
}

func (infra *Entity) ListEntityField(ctx kratosx.Context, req *types.ListEntityFieldRequest) ([]*entity.EntityField, uint32, error) {
	var (
		list  []*entity.EntityField
		total int64
	)

	db := ctx.DB().Model(entity.EntityField{}).Where("entity_id = ?", req.EntityId)
	if req.Name != nil {
		db = db.Where("name LIKE ?", *req.Name+"%")
	}
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	return list, uint32(total), db.Order("`index` asc, id asc").Find(&list).Error
}

func (infra *Entity) CreateEntityField(ctx kratosx.Context, req *entity.EntityField) (uint32, error) {
	if err := ctx.DB().Create(req).Error; err != nil {
		return 0, err
	}
	return req.Id, nil
}

func (infra *Entity) UpdateEntityField(ctx kratosx.Context, req *entity.EntityField) error {
	return ctx.DB().Updates(req).Error
}

func (infra *Entity) UpdateRegisteredEntityField(ctx kratosx.Context, req *entity.EntityField) error {
	return ctx.DB().Model(&entity.EntityField{}).
		Where("id = ?", req.Id).
		Omit("entity_id", "name").
		Updates(req).Error
}

func (infra *Entity) DeleteEntityField(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Where("id = ?", id).Delete(&entity.EntityField{}).Error
}

func (infra *Entity) DeleteEntityFields(ctx kratosx.Context, ids []uint32) error {
	if len(ids) == 0 {
		return nil
	}
	return ctx.DB().Where("id in ?", ids).Delete(&entity.EntityField{}).Error
}

func (infra *Entity) GetEntityRule(ctx kratosx.Context, entityId uint32, name string) (*entity.EntityRule, error) {
	var ent entity.EntityRule
	return &ent, ctx.DB().
		Where("entity_id = ? and name = ?", entityId, name).
		First(&ent).Error
}

func (infra *Entity) ListEntityRule(ctx kratosx.Context, req *types.ListEntityRuleRequest) ([]*entity.EntityRule, uint32, error) {
	var (
		list  []*entity.EntityRule
		total int64
	)

	db := ctx.DB().Model(entity.EntityRule{})
	if req.EntityId != nil {
		db = db.Where("entity_id = ?", *req.EntityId)
	}
	if req.Name != nil {
		db = db.Where("name LIKE ?", *req.Name+"%")
	}
	if req.Status != nil {
		db = db.Where("status = ?", *req.Status)
	}
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if req.Search != nil {
		db = page.SearchScopes(db, req.Search)
	} else {
		db = db.Order("id asc")
	}
	return list, uint32(total), db.Find(&list).Error
}

func (infra *Entity) CreateEntityRule(ctx kratosx.Context, req *entity.EntityRule) (uint32, error) {
	if err := ctx.DB().Create(req).Error; err != nil {
		return 0, err
	}
	return req.Id, nil
}

func (infra *Entity) UpdateEntityRule(ctx kratosx.Context, req *entity.EntityRule) error {
	return ctx.DB().Updates(req).Error
}

func (infra *Entity) UpdateRegisteredEntityRule(ctx kratosx.Context, req *entity.EntityRule) error {
	return ctx.DB().Model(&entity.EntityRule{}).
		Where("id = ?", req.Id).
		Omit("entity_id", "name").
		Updates(req).Error
}

func (infra *Entity) DeleteEntityRule(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Where("id = ?", id).Delete(&entity.EntityRule{}).Error
}

func (infra *Entity) DeleteEntityRules(ctx kratosx.Context, ids []uint32) error {
	if len(ids) == 0 {
		return nil
	}
	return ctx.DB().Where("id in ?", ids).Delete(&entity.EntityRule{}).Error
}
