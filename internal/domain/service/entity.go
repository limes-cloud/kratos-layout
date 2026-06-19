package service

import (
	"fmt"
	"strings"

	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/library/db/gormtranserror"
	"gorm.io/gorm"

	"layout/api/errors"
	"layout/internal/domain/entity"
	"layout/internal/domain/repository"
	"layout/internal/types"
)

type Entity struct {
	repo repository.Entity
}

func NewEntity(repo repository.Entity) *Entity {
	return &Entity{repo: repo}
}

const (
	RegisterEntityModeAppend = "append"
	RegisterEntityModeSync   = "sync"
)

type registerEntityState struct {
	AppId     uint32
	SeenEnt   map[string]struct{}
	SeenField map[uint32]map[string]struct{}
	SeenRule  map[uint32]map[string]struct{}
	Reply     *types.RegisterEntityReply
}

func (srv *Entity) LoadEntity(ctx kratosx.Context) ([]*entity.Entity, error) {
	var list []*entity.Entity
	return list, ctx.DB().Preload("Fields").Find(&list).Error
}

func (srv *Entity) ImportEntity(ctx kratosx.Context, list []*entity.Entity) error {
	return ctx.Transaction(func(ctx kratosx.Context) error {
		for _, ent := range list {
			fields := ent.Fields
			ent.Fields = nil
			id, err := srv.repo.CreateEntity(ctx, ent)
			if err != nil {
				return err
			}
			for ind, field := range fields {
				field.EntityId = id
				field.Index = ind
				if _, err := srv.repo.CreateEntityField(ctx, field); err != nil {
					return err
				}
			}
		}
		return nil
	})
}

func (srv *Entity) ListEntity(ctx kratosx.Context, req *types.ListEntityRequest) ([]*entity.Entity, uint32, error) {
	list, total, err := srv.repo.ListEntity(ctx, req)
	if err != nil {
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}

func (srv *Entity) CreateEntity(ctx kratosx.Context, req *entity.Entity) (uint32, error) {
	id, err := srv.repo.CreateEntity(ctx, req)
	if err != nil {
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

func (srv *Entity) UpdateEntity(ctx kratosx.Context, req *entity.Entity) error {
	if err := srv.repo.UpdateEntity(ctx, req); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}

func (srv *Entity) DeleteEntity(ctx kratosx.Context, id uint32) error {
	if err := srv.repo.DeleteEntity(ctx, id); err != nil {
		return errors.DeleteError(err.Error())
	}
	return nil
}

func (srv *Entity) ListEntityField(ctx kratosx.Context, req *types.ListEntityFieldRequest) ([]*entity.EntityField, uint32, error) {
	list, total, err := srv.repo.ListEntityField(ctx, req)
	if err != nil {
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}

func (srv *Entity) CreateEntityField(ctx kratosx.Context, req *entity.EntityField) (uint32, error) {
	id, err := srv.repo.CreateEntityField(ctx, req)
	if err != nil {
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

func (srv *Entity) UpdateEntityField(ctx kratosx.Context, req *entity.EntityField) error {
	if err := srv.repo.UpdateEntityField(ctx, req); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}

func (srv *Entity) DeleteEntityField(ctx kratosx.Context, id uint32) error {
	if err := srv.repo.DeleteEntityField(ctx, id); err != nil {
		return errors.DeleteError(err.Error())
	}
	return nil
}

func (srv *Entity) ListEntityRule(ctx kratosx.Context, req *types.ListEntityRuleRequest) ([]*entity.EntityRule, uint32, error) {
	list, total, err := srv.repo.ListEntityRule(ctx, req)
	if err != nil {
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}

func (srv *Entity) CreateEntityRule(ctx kratosx.Context, req *entity.EntityRule) (uint32, error) {
	id, err := srv.repo.CreateEntityRule(ctx, req)
	if err != nil {
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

func (srv *Entity) UpdateEntityRule(ctx kratosx.Context, req *entity.EntityRule) error {
	if err := srv.repo.UpdateEntityRule(ctx, req); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}

func (srv *Entity) DeleteEntityRule(ctx kratosx.Context, id uint32) error {
	if err := srv.repo.DeleteEntityRule(ctx, id); err != nil {
		return errors.DeleteError(err.Error())
	}
	return nil
}

func (srv *Entity) RegisterEntity(ctx kratosx.Context, req *types.RegisterEntityRequest) (*types.RegisterEntityReply, error) {
	mode := strings.ToLower(strings.TrimSpace(req.Mode))
	if mode == "" {
		mode = RegisterEntityModeAppend
	}
	if mode != RegisterEntityModeAppend && mode != RegisterEntityModeSync {
		return nil, errors.ParamsError("mode仅支持append或sync")
	}

	state := &registerEntityState{
		SeenEnt:   map[string]struct{}{},
		SeenField: map[uint32]map[string]struct{}{},
		SeenRule:  map[uint32]map[string]struct{}{},
		Reply:     &types.RegisterEntityReply{},
	}
	if err := ctx.Transaction(func(ctx kratosx.Context) error {
		for _, item := range req.Entities {
			if err := srv.registerEntityNode(ctx, state, item); err != nil {
				return err
			}
		}
		if mode != RegisterEntityModeSync {
			return nil
		}
		return srv.deleteUnregisteredEntities(ctx, state)
	}); err != nil {
		return nil, errors.UpdateError(err.Error())
	}
	return state.Reply, nil
}

func (srv *Entity) registerEntityNode(ctx kratosx.Context, state *registerEntityState, node *types.RegisterEntityNode) error {
	if node == nil {
		return nil
	}
	ent := &entity.Entity{
		AppId:    state.AppId,
		Database: strings.TrimSpace(node.Database),
		Name:     strings.TrimSpace(node.Name),
		Comment:  strings.TrimSpace(node.Comment),
	}
	if ent.Database == "" || ent.Name == "" {
		return fmt.Errorf("实体缺少database或name")
	}
	entKey := strings.ToLower(ent.Database) + ":" + strings.ToLower(ent.Name)
	if _, ok := state.SeenEnt[entKey]; ok {
		return fmt.Errorf("实体重复: %s", entKey)
	}
	state.SeenEnt[entKey] = struct{}{}

	old, err := srv.repo.GetEntityByName(ctx, state.AppId, ent.Database, ent.Name)
	if gormtranserror.Is(err, gorm.ErrRecordNotFound) {
		id, err := srv.repo.CreateEntity(ctx, ent)
		if err != nil {
			return err
		}
		ent.Id = id
		state.Reply.Created++
	} else if err != nil {
		return err
	} else {
		ent.Id = old.Id
		if err := srv.repo.UpdateRegisteredEntity(ctx, ent); err != nil {
			return err
		}
		state.Reply.Updated++
	}
	if err := srv.registerEntityFields(ctx, state, ent.Id, node.Fields); err != nil {
		return err
	}
	return srv.registerEntityRules(ctx, state, ent.Id, node.Rules)
}

func (srv *Entity) registerEntityFields(ctx kratosx.Context, state *registerEntityState, entityId uint32, list []*types.RegisterEntityFieldNode) error {
	for ind, item := range list {
		if item == nil {
			continue
		}
		name := strings.TrimSpace(item.Name)
		if name == "" {
			return fmt.Errorf("实体字段名称不能为空")
		}
		if _, ok := state.SeenField[entityId]; !ok {
			state.SeenField[entityId] = map[string]struct{}{}
		}
		if _, ok := state.SeenField[entityId][name]; ok {
			return fmt.Errorf("实体字段重复: %s", name)
		}
		state.SeenField[entityId][name] = struct{}{}

		field := &entity.EntityField{
			EntityId: entityId,
			Name:     name,
			Comment:  strings.TrimSpace(item.Comment),
			Index:    item.Index,
		}
		if field.Index == 0 {
			field.Index = ind
		}

		old, err := srv.repo.GetEntityField(ctx, entityId, name)
		if gormtranserror.Is(err, gorm.ErrRecordNotFound) {
			if _, err := srv.repo.CreateEntityField(ctx, field); err != nil {
				return err
			}
			state.Reply.Created++
			continue
		}
		if err != nil {
			return err
		}
		field.Id = old.Id
		if err := srv.repo.UpdateRegisteredEntityField(ctx, field); err != nil {
			return err
		}
		state.Reply.Updated++
	}
	return nil
}

func (srv *Entity) registerEntityRules(ctx kratosx.Context, state *registerEntityState, entityId uint32, list []*types.RegisterEntityRuleNode) error {
	for _, item := range list {
		if item == nil {
			continue
		}
		name := strings.TrimSpace(item.Name)
		if name == "" {
			return fmt.Errorf("实体规则名称不能为空")
		}
		if _, ok := state.SeenRule[entityId]; !ok {
			state.SeenRule[entityId] = map[string]struct{}{}
		}
		if _, ok := state.SeenRule[entityId][name]; ok {
			return fmt.Errorf("实体规则重复: %s", name)
		}
		state.SeenRule[entityId][name] = struct{}{}

		rule := &entity.EntityRule{
			EntityId:    entityId,
			Name:        name,
			Status:      item.Status,
			Expression:  strings.TrimSpace(item.Expression),
			Description: strings.TrimSpace(item.Description),
		}
		old, err := srv.repo.GetEntityRule(ctx, entityId, name)
		if gormtranserror.Is(err, gorm.ErrRecordNotFound) {
			if _, err := srv.repo.CreateEntityRule(ctx, rule); err != nil {
				return err
			}
			state.Reply.Created++
			continue
		}
		if err != nil {
			return err
		}
		rule.Id = old.Id
		if err := srv.repo.UpdateRegisteredEntityRule(ctx, rule); err != nil {
			return err
		}
		state.Reply.Updated++
	}
	return nil
}

func (srv *Entity) deleteUnregisteredEntities(ctx kratosx.Context, state *registerEntityState) error {
	list, _, err := srv.repo.ListEntity(ctx, &types.ListEntityRequest{AppId: &state.AppId})
	if err != nil {
		return err
	}
	var ids []uint32
	for _, item := range list {
		key := strings.ToLower(strings.TrimSpace(item.Database)) + ":" + strings.ToLower(strings.TrimSpace(item.Name))
		if _, ok := state.SeenEnt[key]; ok {
			continue
		}
		ids = append(ids, item.Id)
	}
	if err := srv.repo.DeleteEntities(ctx, ids); err != nil {
		return err
	}
	state.Reply.Deleted += uint32(len(ids))
	return nil
}
