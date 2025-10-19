package app

import (
	"context"
	"github.com/limes-cloud/kratosx/model"
	"partyaffairs/api/information"
	"partyaffairs/internal/core"

	"github.com/go-kratos/kratos/v2/transport/http"
	"partyaffairs/internal/domain/entity"
	"partyaffairs/internal/domain/service"
	"partyaffairs/internal/infra/dbs"
	"partyaffairs/internal/infra/rpc"
	"partyaffairs/internal/types"
)

type Information struct {
	information.UnimplementedInformationServer
	srv *service.InformationService
}

func NewInformation() *Information {
	return &Information{
		srv: service.NewInformationService(dbs.NewInformation(), rpc.NewFile()),
	}
}

func init() {
	register(func(hs *http.Server) {
		srv := NewInformation()
		information.RegisterInformationHTTPServer(hs, srv)
	})
}

// ListInformationClassify 获取任务分组列表
func (s *Information) ListInformationClassify(c context.Context, _ *information.ListInformationClassifyRequest) (*information.ListInformationClassifyReply, error) {
	list, err := s.srv.ListInformationClassify(core.MustContext(c))
	if err != nil {
		return nil, err
	}
	reply := information.ListInformationClassifyReply{}
	for _, item := range list {
		reply.List = append(reply.List, &information.ListInformationClassifyReply_InformationClassify{
			Id:        item.Id,
			Name:      item.Name,
			Weight:    item.Weight,
			CreatedAt: uint32(item.CreatedAt),
			UpdatedAt: uint32(item.UpdatedAt),
		})
	}
	return &reply, nil
}

// CreateInformationClassify 创建任务分组
func (s *Information) CreateInformationClassify(c context.Context, req *information.CreateInformationClassifyRequest) (*information.CreateInformationClassifyReply, error) {
	id, err := s.srv.CreateInformationClassify(core.MustContext(c), &entity.InformationClassify{
		Name:   req.Name,
		Weight: req.Weight,
	})
	if err != nil {
		return nil, err
	}
	return &information.CreateInformationClassifyReply{Id: id}, nil
}

// UpdateInformationClassify 更新任务分组
func (s *Information) UpdateInformationClassify(c context.Context, req *information.UpdateInformationClassifyRequest) (*information.UpdateInformationClassifyReply, error) {
	if err := s.srv.UpdateInformationClassify(core.MustContext(c), &entity.InformationClassify{
		BaseTenantModel: model.BaseTenantModel{Id: req.Id},
		Name:            req.Name,
		Weight:          req.Weight,
	}); err != nil {
		return nil, err
	}
	return &information.UpdateInformationClassifyReply{}, nil
}

// DeleteInformationClassify 删除任务分组
func (s *Information) DeleteInformationClassify(c context.Context, req *information.DeleteInformationClassifyRequest) (*information.DeleteInformationClassifyReply, error) {
	err := s.srv.DeleteInformationClassify(core.MustContext(c), req.Id)
	if err != nil {
		return nil, err
	}
	return &information.DeleteInformationClassifyReply{}, nil
}

// GetInformation 获取指定的资讯信息
func (s *Information) GetInformation(c context.Context, req *information.GetInformationRequest) (*information.GetInformationReply, error) {
	res, err := s.srv.GetInformation(core.MustContext(c), req.Id)
	if err != nil {
		return nil, err
	}
	return &information.GetInformationReply{
		Id:          res.Id,
		ClassifyId:  res.ClassifyId,
		Title:       res.Title,
		Description: res.Description,
		Cover:       res.Cover,
		Unit:        res.Unit,
		Content:     res.Content,
		IsTop:       res.IsTop,
		Status:      res.Status,
		Read:        uint32(res.Read),
		Classify: &information.GetInformationReply_Classify{
			Id:   res.Classify.Id,
			Name: res.Classify.Name,
		},
		CreatedAt: uint32(res.CreatedAt),
		UpdatedAt: uint32(res.UpdatedAt),
	}, nil
}

// ListInformation 获取资讯信息列表
func (s *Information) ListInformation(c context.Context, req *information.ListInformationRequest) (*information.ListInformationReply, error) {
	list, total, err := s.srv.ListInformation(core.MustContext(c), &types.ListInformationRequest{
		Page:       req.Page,
		PageSize:   req.PageSize,
		Title:      req.Title,
		IsTop:      req.IsTop,
		Status:     req.Status,
		ClassifyId: req.ClassifyId,
	})
	if err != nil {
		return nil, err
	}

	reply := information.ListInformationReply{Total: total}
	for _, item := range list {
		reply.List = append(reply.List, &information.ListInformationReply_Information{
			Id:          item.Id,
			ClassifyId:  item.ClassifyId,
			Title:       item.Title,
			Description: item.Description,
			Cover:       item.Cover,
			Unit:        item.Unit,
			IsTop:       item.IsTop,
			Status:      item.Status,
			Read:        uint32(item.Read),
			Classify: &information.ListInformationReply_Classify{
				Id:   item.Classify.Id,
				Name: item.Classify.Name,
			},
			CreatedAt: uint32(item.CreatedAt),
			UpdatedAt: uint32(item.UpdatedAt),
		})
	}
	return &reply, nil
}

// CreateInformation 创建资讯信息
func (s *Information) CreateInformation(c context.Context, req *information.CreateInformationRequest) (*information.CreateInformationReply, error) {
	id, err := s.srv.CreateInformation(core.MustContext(c), &entity.Information{
		ClassifyId:  req.ClassifyId,
		Title:       req.Title,
		Description: req.Description,
		Cover:       req.Cover,
		Unit:        req.Unit,
		Content:     req.Content,
		IsTop:       req.IsTop,
		Status:      req.Status,
	})
	if err != nil {
		return nil, err
	}

	return &information.CreateInformationReply{Id: id}, nil
}

// UpdateInformation 更新资讯信息
func (s *Information) UpdateInformation(c context.Context, req *information.UpdateInformationRequest) (*information.UpdateInformationReply, error) {
	if err := s.srv.UpdateInformation(core.MustContext(c), &entity.Information{
		BaseTenantModel: model.BaseTenantModel{Id: req.Id},
		ClassifyId:      req.ClassifyId,
		Title:           req.Title,
		Description:     req.Description,
		Cover:           req.Cover,
		Unit:            req.Unit,
		Content:         req.Content,
		IsTop:           req.IsTop,
		Status:          req.Status,
	}); err != nil {
		return nil, err
	}
	return &information.UpdateInformationReply{}, nil
}

// DeleteInformation 删除资讯信息
func (s *Information) DeleteInformation(c context.Context, req *information.DeleteInformationRequest) (*information.DeleteInformationReply, error) {
	return &information.DeleteInformationReply{}, s.srv.DeleteInformation(core.MustContext(c), req.Id)
}
