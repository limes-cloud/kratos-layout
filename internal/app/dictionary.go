package app

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/value"

	"layout/api/dictionary"
	"layout/api/errors"
	"layout/internal/domain/entity"
	"layout/internal/domain/service"
	"layout/internal/infra/dbs"
	"layout/internal/types"
)

type Dictionary struct {
	dictionary.UnimplementedDictionaryServer
	srv *service.Dictionary
}

func NewDictionary() *Dictionary {
	return &Dictionary{
		srv: service.NewDictionary(dbs.NewDictionary()),
	}
}

func init() {
	register(func(hs *http.Server, gs *grpc.Server) {
		app := NewDictionary()
		dictionary.RegisterDictionaryHTTPServer(hs, app)
		dictionary.RegisterDictionaryServer(gs, app)
	})
}

// GetDictionary 获取指定的字典目录
func (s *Dictionary) GetDictionary(c context.Context, req *dictionary.GetDictionaryRequest) (*dictionary.GetDictionaryReply, error) {
	ctx := kratosx.MustContext(c)
	result, err := s.srv.GetDictionary(ctx, &types.GetDictionaryRequest{
		Id:      req.Id,
		Keyword: req.Keyword,
	})
	if err != nil {
		return nil, err
	}

	reply := dictionary.GetDictionaryReply{}
	if err = value.Transform(result, &reply); err != nil {
		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	return &reply, nil
}

// ListDictionary 获取字典目录列表
func (s *Dictionary) ListDictionary(c context.Context, req *dictionary.ListDictionaryRequest) (*dictionary.ListDictionaryReply, error) {
	ctx := kratosx.MustContext(c)

	result, total, err := s.srv.ListDictionary(ctx, &types.ListDictionaryRequest{
		Page:     req.Page,
		PageSize: req.PageSize,
		Order:    req.Order,
		OrderBy:  req.OrderBy,
		Keyword:  req.Keyword,
		Name:     req.Name,
	})
	if err != nil {
		return nil, err
	}

	reply := dictionary.ListDictionaryReply{Total: total}
	if err = value.Transform(result, &reply.List); err != nil {
		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	return &reply, nil
}

// CreateDictionary 创建字典目录
func (s *Dictionary) CreateDictionary(c context.Context, req *dictionary.CreateDictionaryRequest) (*dictionary.CreateDictionaryReply, error) {
	var (
		ctx = kratosx.MustContext(c)
		ent = entity.Dictionary{}
	)

	if err := value.Transform(req, &ent); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	id, err := s.srv.CreateDictionary(ctx, &ent)
	if err != nil {
		return nil, err
	}

	return &dictionary.CreateDictionaryReply{Id: id}, nil
}

// UpdateDictionary 更新字典目录
func (s *Dictionary) UpdateDictionary(c context.Context, req *dictionary.UpdateDictionaryRequest) (*dictionary.UpdateDictionaryReply, error) {
	var (
		ctx = kratosx.MustContext(c)
		ent = entity.Dictionary{}
	)

	if err := value.Transform(req, &ent); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	if err := s.srv.UpdateDictionary(ctx, &ent); err != nil {
		return nil, err
	}

	return &dictionary.UpdateDictionaryReply{}, nil
}

// DeleteDictionary 删除字典目录
func (s *Dictionary) DeleteDictionary(c context.Context, req *dictionary.DeleteDictionaryRequest) (*dictionary.DeleteDictionaryReply, error) {
	err := s.srv.DeleteDictionary(kratosx.MustContext(c), req.Id)
	if err != nil {
		return nil, err
	}
	return &dictionary.DeleteDictionaryReply{}, nil
}
