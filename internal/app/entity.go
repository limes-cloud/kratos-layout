package app

import (
	"context"
	"encoding/json"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/model"
	"github.com/limes-cloud/kratosx/pkg/value"
	"google.golang.org/protobuf/types/known/emptypb"

	apientity "tianji/api/entity"
	"tianji/api/errors"
	"tianji/internal/domain/entity"
	"tianji/internal/domain/service"
	"tianji/internal/infra/dbs"
	"tianji/internal/types"
)

type Entity struct {
	apientity.UnimplementedEntityServer
	srv *service.Entity
}

func NewEntity() *Entity {
	return &Entity{
		srv: service.NewEntity(dbs.NewEntity()),
	}
}

func init() {
	register(func(hs *http.Server, gs *grpc.Server) {
		srv := NewEntity()
		apientity.RegisterEntityHTTPServer(hs, srv)
		apientity.RegisterEntityServer(gs, srv)
	})
}

func (app *Entity) LoadEntity(c context.Context, _ *emptypb.Empty) (*apientity.LoadEntityReply, error) {
	ctx := kratosx.MustContext(c)
	list, err := app.srv.LoadEntity(ctx)
	if err != nil {
		return nil, err
	}
	reply := apientity.LoadEntityReply{Total: uint32(len(list))}
	if err := value.Transform(list, &reply.List); err != nil {
		return nil, errors.TransformError()
	}
	return &reply, nil
}

func (app *Entity) ImportEntity(c context.Context, req *apientity.ImportEntityRequest) (*emptypb.Empty, error) {
	ctx := kratosx.MustContext(c)
	var list []*entity.Entity
	if err := value.Transform(req.List, &list); err != nil {
		return nil, errors.TransformError()
	}
	for _, item := range list {
		item.AppId = req.AppId
	}
	if err := app.srv.ImportEntity(ctx, list); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (app *Entity) ListEntity(c context.Context, req *apientity.ListEntityRequest) (*apientity.ListEntityReply, error) {
	ctx := kratosx.MustContext(c)
	in := types.ListEntityRequest{}
	if err := value.Transform(req, &in); err != nil {
		return nil, errors.TransformError()
	}
	list, total, err := app.srv.ListEntity(ctx, &in)
	if err != nil {
		return nil, err
	}
	reply := apientity.ListEntityReply{Total: total}
	if err := value.Transform(list, &reply.List); err != nil {
		return nil, errors.TransformError()
	}
	return &reply, nil
}

func (app *Entity) CreateEntity(c context.Context, req *apientity.CreateEntityRequest) (*apientity.CreateEntityReply, error) {
	ctx := kratosx.MustContext(c)
	in := entity.Entity{}
	if err := value.Transform(req, &in); err != nil {
		return nil, errors.TransformError()
	}
	id, err := app.srv.CreateEntity(ctx, &in)
	if err != nil {
		return nil, err
	}
	return &apientity.CreateEntityReply{Id: id}, nil
}

func (app *Entity) UpdateEntity(c context.Context, req *apientity.UpdateEntityRequest) (*apientity.UpdateEntityReply, error) {
	ctx := kratosx.MustContext(c)
	in := entity.Entity{}
	if err := value.Transform(req, &in); err != nil {
		return nil, errors.TransformError()
	}
	if err := app.srv.UpdateEntity(ctx, &in); err != nil {
		return nil, err
	}
	return &apientity.UpdateEntityReply{}, nil
}

func (app *Entity) DeleteEntity(c context.Context, req *apientity.DeleteEntityRequest) (*apientity.DeleteEntityReply, error) {
	return &apientity.DeleteEntityReply{}, app.srv.DeleteEntity(kratosx.MustContext(c), req.Id)
}

func (app *Entity) RegisterEntity(c context.Context, req *apientity.RegisterEntityRequest) (*apientity.RegisterEntityReply, error) {
	reply, err := app.srv.RegisterEntity(kratosx.MustContext(c), &types.RegisterEntityRequest{
		AppKeyword: req.AppKeyword,
		Mode:       req.Mode,
		Entities:   transRegisterEntityNodes(req.Entities),
	})
	if err != nil {
		return nil, err
	}
	return &apientity.RegisterEntityReply{
		Created: reply.Created,
		Updated: reply.Updated,
		Deleted: reply.Deleted,
	}, nil
}

func (app *Entity) ListEntityField(c context.Context, req *apientity.ListEntityFieldRequest) (*apientity.ListEntityFieldReply, error) {
	ctx := kratosx.MustContext(c)
	in := types.ListEntityFieldRequest{}
	if err := value.Transform(req, &in); err != nil {
		return nil, errors.TransformError()
	}
	list, total, err := app.srv.ListEntityField(ctx, &in)
	if err != nil {
		return nil, err
	}
	reply := apientity.ListEntityFieldReply{Total: total}
	if err := value.Transform(list, &reply.List); err != nil {
		return nil, errors.TransformError()
	}
	return &reply, nil
}

func (app *Entity) CreateEntityField(c context.Context, req *apientity.CreateEntityFieldRequest) (*apientity.CreateEntityFieldReply, error) {
	ctx := kratosx.MustContext(c)
	in := entity.EntityField{}
	if err := value.Transform(req, &in); err != nil {
		return nil, errors.TransformError()
	}
	id, err := app.srv.CreateEntityField(ctx, &in)
	if err != nil {
		return nil, err
	}
	return &apientity.CreateEntityFieldReply{Id: id}, nil
}

func (app *Entity) UpdateEntityField(c context.Context, req *apientity.UpdateEntityFieldRequest) (*apientity.UpdateEntityFieldReply, error) {
	ctx := kratosx.MustContext(c)
	in := entity.EntityField{}
	if err := value.Transform(req, &in); err != nil {
		return nil, errors.TransformError()
	}
	if err := app.srv.UpdateEntityField(ctx, &in); err != nil {
		return nil, err
	}
	return &apientity.UpdateEntityFieldReply{}, nil
}

func (app *Entity) DeleteEntityField(c context.Context, req *apientity.DeleteEntityFieldRequest) (*apientity.DeleteEntityFieldReply, error) {
	return &apientity.DeleteEntityFieldReply{}, app.srv.DeleteEntityField(kratosx.MustContext(c), req.Id)
}

func (app *Entity) ListEntityRule(c context.Context, req *apientity.ListEntityRuleRequest) (*apientity.ListEntityRuleReply, error) {
	ctx := kratosx.MustContext(c)
	in := types.ListEntityRuleRequest{}
	if err := value.Transform(req, &in); err != nil {
		return nil, errors.TransformError()
	}
	list, total, err := app.srv.ListEntityRule(ctx, &in)
	if err != nil {
		return nil, err
	}
	reply := apientity.ListEntityRuleReply{Total: total}
	for _, item := range list {
		rule := &apientity.EntityRule{
			Id:          item.Id,
			EntityId:    item.EntityId,
			Name:        item.Name,
			Description: item.Description,
			Status:      value.Value(item.Status),
			CreatedAt:   uint32(item.CreatedAt),
			UpdatedAt:   uint32(item.UpdatedAt),
		}
		_ = json.Unmarshal([]byte(item.Expression), &rule.Expression)
		reply.List = append(reply.List, rule)
	}
	return &reply, nil
}

func (app *Entity) CreateEntityRule(c context.Context, req *apientity.CreateEntityRuleRequest) (*apientity.CreateEntityRuleReply, error) {
	id, err := app.srv.CreateEntityRule(kratosx.MustContext(c), &entity.EntityRule{
		EntityId:    req.EntityId,
		Name:        req.Name,
		Status:      req.Status,
		Expression:  value.ObjToString(req.Expression),
		Description: req.Description,
	})
	if err != nil {
		return nil, err
	}
	return &apientity.CreateEntityRuleReply{Id: id}, nil
}

func (app *Entity) UpdateEntityRule(c context.Context, req *apientity.UpdateEntityRuleRequest) (*apientity.UpdateEntityRuleReply, error) {
	if err := app.srv.UpdateEntityRule(kratosx.MustContext(c), &entity.EntityRule{
		BaseModel:   model.BaseModel{Id: req.Id},
		EntityId:    value.Value(req.EntityId),
		Name:        value.Value(req.Name),
		Status:      req.Status,
		Expression:  value.ObjToString(req.Expression),
		Description: value.Value(req.Description),
	}); err != nil {
		return nil, err
	}
	return &apientity.UpdateEntityRuleReply{}, nil
}

func (app *Entity) DeleteEntityRule(c context.Context, req *apientity.DeleteEntityRuleRequest) (*apientity.DeleteEntityRuleReply, error) {
	return &apientity.DeleteEntityRuleReply{}, app.srv.DeleteEntityRule(kratosx.MustContext(c), req.Id)
}

func transRegisterEntityNodes(list []*apientity.RegisterEntityNode) []*types.RegisterEntityNode {
	var res []*types.RegisterEntityNode
	for _, item := range list {
		if item == nil {
			continue
		}
		res = append(res, &types.RegisterEntityNode{
			Database: item.Database,
			Name:     item.Name,
			Comment:  item.Comment,
			Fields:   transRegisterEntityFields(item.Fields),
			Rules:    transRegisterEntityRules(item.Rules),
		})
	}
	return res
}

func transRegisterEntityFields(list []*apientity.RegisterEntityField) []*types.RegisterEntityFieldNode {
	var res []*types.RegisterEntityFieldNode
	for _, item := range list {
		if item == nil {
			continue
		}
		res = append(res, &types.RegisterEntityFieldNode{
			Name:    item.Name,
			Comment: item.Comment,
			Index:   int(item.Index),
		})
	}
	return res
}

func transRegisterEntityRules(list []*apientity.RegisterEntityRule) []*types.RegisterEntityRuleNode {
	var res []*types.RegisterEntityRuleNode
	for _, item := range list {
		if item == nil {
			continue
		}
		res = append(res, &types.RegisterEntityRuleNode{
			Name:        item.Name,
			Status:      item.Status,
			Expression:  item.Expression,
			Description: item.Description,
		})
	}
	return res
}
