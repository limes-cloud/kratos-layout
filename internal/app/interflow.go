package app

import (
	"context"
	"github.com/limes-cloud/kratosx/model"
	"github.com/limes-cloud/kratosx/model/page"
	"github.com/limes-cloud/kratosx/pkg/value"
	"github.com/limes-cloud/manager/api/errors"
	"partyaffairs/api/interflow"
	"partyaffairs/internal/core"
	"partyaffairs/internal/infra/rpc"

	"github.com/go-kratos/kratos/v2/transport/http"
	"partyaffairs/internal/domain/entity"
	"partyaffairs/internal/domain/service"
	"partyaffairs/internal/infra/dbs"
	"partyaffairs/internal/types"
)

type Interflow struct {
	interflow.UnimplementedInterflowServer
	srv *service.InterflowService
}

func NewInterflow() *Interflow {
	return &Interflow{
		srv: service.NewInterflowService(dbs.NewInterflow(), rpc.NewUser()),
	}
}

func init() {
	register(func(hs *http.Server) {
		srv := NewInterflow()
		interflow.RegisterInterflowHTTPServer(hs, srv)
	})
}

func (s *Interflow) ListInterflowHistory(c context.Context, _ *interflow.ListInterflowHistoryRequest) (*interflow.ListInterflowHistoryReply, error) {
	list, err := s.srv.ListInterflowHistory(core.MustContext(c))
	if err != nil {
		return nil, err
	}

	reply := interflow.ListInterflowHistoryReply{}
	if err := value.Transform(list, &reply.List); err != nil {
		return nil, errors.TransformError()
	}

	return &reply, nil
}

// ListInterflowClassify 获取任务分组列表
func (s *Interflow) ListInterflowClassify(c context.Context, _ *interflow.ListInterflowClassifyRequest) (*interflow.ListInterflowClassifyReply, error) {
	list, err := s.srv.ListInterflowClassify(core.MustContext(c))
	if err != nil {
		return nil, err
	}
	reply := interflow.ListInterflowClassifyReply{}
	if err := value.Transform(list, &reply.List); err != nil {
		return nil, err
	}
	return &reply, nil
}

// CreateInterflowClassify 创建任务分组
func (s *Interflow) CreateInterflowClassify(c context.Context, req *interflow.CreateInterflowClassifyRequest) (*interflow.CreateInterflowClassifyReply, error) {
	id, err := s.srv.CreateInterflowClassify(core.MustContext(c), &entity.InterflowClassify{
		Name:        req.Name,
		Weight:      req.Weight,
		Description: req.Description,
		Person:      req.Person,
	})
	if err != nil {
		return nil, err
	}
	return &interflow.CreateInterflowClassifyReply{Id: id}, nil
}

// UpdateInterflowClassify 更新任务分组
func (s *Interflow) UpdateInterflowClassify(c context.Context, req *interflow.UpdateInterflowClassifyRequest) (*interflow.UpdateInterflowClassifyReply, error) {
	if err := s.srv.UpdateInterflowClassify(core.MustContext(c), &entity.InterflowClassify{
		BaseTenantModel: model.BaseTenantModel{Id: req.Id},
		Name:            req.Name,
		Weight:          req.Weight,
		Description:     req.Description,
		Person:          req.Person,
	}); err != nil {
		return nil, err
	}
	return &interflow.UpdateInterflowClassifyReply{}, nil
}

// DeleteInterflowClassify 删除任务分组
func (s *Interflow) DeleteInterflowClassify(c context.Context, req *interflow.DeleteInterflowClassifyRequest) (*interflow.DeleteInterflowClassifyReply, error) {
	err := s.srv.DeleteInterflowClassify(core.MustContext(c), req.Id)
	if err != nil {
		return nil, err
	}
	return &interflow.DeleteInterflowClassifyReply{}, nil
}

// ListInterflow 获取资讯信息列表
func (s *Interflow) ListInterflow(c context.Context, req *interflow.ListInterflowRequest) (*interflow.ListInterflowReply, error) {
	list, total, err := s.srv.ListInterflow(core.MustContext(c), &types.ListInterflowRequest{
		Search: &page.Search{
			Page:     req.Page,
			PageSize: req.PageSize,
		},
		FromUserId: req.GetFromUserId(),
		ToUserId:   req.GetToUserId(),
		FromCur:    req.FromCur,
	})
	if err != nil {
		return nil, err
	}

	reply := interflow.ListInterflowReply{Total: total}
	if err := value.Transform(list, &reply.List); err != nil {
		return nil, errors.TransformError()
	}

	return &reply, nil
}

// CreateInterflow 创建资讯信息
func (s *Interflow) CreateInterflow(c context.Context, req *interflow.CreateInterflowRequest) (*interflow.CreateInterflowReply, error) {
	id, err := s.srv.CreateInterflow(core.MustContext(c), &entity.Interflow{
		ToUserID: req.ToUserId,
		Type:     req.Type,
		Content:  req.Content,
	})
	if err != nil {
		return nil, err
	}

	return &interflow.CreateInterflowReply{Id: id}, nil
}
