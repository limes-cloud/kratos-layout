package app

import (
	"context"
	"github.com/limes-cloud/kratosx/model"
	"partyaffairs/api/public"
	"partyaffairs/internal/core"

	"github.com/go-kratos/kratos/v2/transport/http"
	"partyaffairs/internal/domain/entity"
	"partyaffairs/internal/domain/service"
	"partyaffairs/internal/infra/dbs"
	"partyaffairs/internal/infra/rpc"
	"partyaffairs/internal/types"
)

type Public struct {
	public.UnimplementedPublicServer
	srv *service.PublicService
}

func NewPublic() *Public {
	return &Public{
		srv: service.NewPublicService(dbs.NewPublic(), rpc.NewFile()),
	}
}

func init() {
	register(func(hs *http.Server) {
		srv := NewPublic()
		public.RegisterPublicHTTPServer(hs, srv)
	})
}

// ListPublicClassify 获取任务分组列表
func (s *Public) ListPublicClassify(c context.Context, _ *public.ListPublicClassifyRequest) (*public.ListPublicClassifyReply, error) {
	list, err := s.srv.ListPublicClassify(core.MustContext(c))
	if err != nil {
		return nil, err
	}
	reply := public.ListPublicClassifyReply{}
	for _, item := range list {
		reply.List = append(reply.List, &public.ListPublicClassifyReply_PublicClassify{
			Id:        item.Id,
			Name:      item.Name,
			Weight:    item.Weight,
			CreatedAt: uint32(item.CreatedAt),
			UpdatedAt: uint32(item.UpdatedAt),
		})
	}
	return &reply, nil
}

// CreatePublicClassify 创建任务分组
func (s *Public) CreatePublicClassify(c context.Context, req *public.CreatePublicClassifyRequest) (*public.CreatePublicClassifyReply, error) {
	id, err := s.srv.CreatePublicClassify(core.MustContext(c), &entity.PublicClassify{
		Name:   req.Name,
		Weight: req.Weight,
	})
	if err != nil {
		return nil, err
	}
	return &public.CreatePublicClassifyReply{Id: id}, nil
}

// UpdatePublicClassify 更新任务分组
func (s *Public) UpdatePublicClassify(c context.Context, req *public.UpdatePublicClassifyRequest) (*public.UpdatePublicClassifyReply, error) {
	if err := s.srv.UpdatePublicClassify(core.MustContext(c), &entity.PublicClassify{
		BaseTenantModel: model.BaseTenantModel{Id: req.Id},
		Name:            req.Name,
		Weight:          req.Weight,
	}); err != nil {
		return nil, err
	}
	return &public.UpdatePublicClassifyReply{}, nil
}

// DeletePublicClassify 删除任务分组
func (s *Public) DeletePublicClassify(c context.Context, req *public.DeletePublicClassifyRequest) (*public.DeletePublicClassifyReply, error) {
	err := s.srv.DeletePublicClassify(core.MustContext(c), req.Id)
	if err != nil {
		return nil, err
	}
	return &public.DeletePublicClassifyReply{}, nil
}

// GetPublic 获取指定的资讯信息
func (s *Public) GetPublic(c context.Context, req *public.GetPublicRequest) (*public.GetPublicReply, error) {
	res, err := s.srv.GetPublic(core.MustContext(c), req.Id)
	if err != nil {
		return nil, err
	}
	return &public.GetPublicReply{
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
		Classify: &public.GetPublicReply_Classify{
			Id:   res.Classify.Id,
			Name: res.Classify.Name,
		},
		CreatedAt: uint32(res.CreatedAt),
		UpdatedAt: uint32(res.UpdatedAt),
	}, nil
}

// ListPublic 获取资讯信息列表
func (s *Public) ListPublic(c context.Context, req *public.ListPublicRequest) (*public.ListPublicReply, error) {
	list, total, err := s.srv.ListPublic(core.MustContext(c), &types.ListPublicRequest{
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

	reply := public.ListPublicReply{Total: total}
	for _, item := range list {
		reply.List = append(reply.List, &public.ListPublicReply_Public{
			Id:          item.Id,
			ClassifyId:  item.ClassifyId,
			Title:       item.Title,
			Description: item.Description,
			Cover:       item.Cover,
			Unit:        item.Unit,
			IsTop:       item.IsTop,
			Status:      item.Status,
			Read:        uint32(item.Read),
			Classify: &public.ListPublicReply_Classify{
				Id:   item.Classify.Id,
				Name: item.Classify.Name,
			},
			CreatedAt: uint32(item.CreatedAt),
			UpdatedAt: uint32(item.UpdatedAt),
		})
	}
	return &reply, nil
}

// CreatePublic 创建资讯信息
func (s *Public) CreatePublic(c context.Context, req *public.CreatePublicRequest) (*public.CreatePublicReply, error) {
	id, err := s.srv.CreatePublic(core.MustContext(c), &entity.Public{
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

	return &public.CreatePublicReply{Id: id}, nil
}

// UpdatePublic 更新资讯信息
func (s *Public) UpdatePublic(c context.Context, req *public.UpdatePublicRequest) (*public.UpdatePublicReply, error) {
	if err := s.srv.UpdatePublic(core.MustContext(c), &entity.Public{
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
	return &public.UpdatePublicReply{}, nil
}

// DeletePublic 删除资讯信息
func (s *Public) DeletePublic(c context.Context, req *public.DeletePublicRequest) (*public.DeletePublicReply, error) {
	return &public.DeletePublicReply{}, s.srv.DeletePublic(core.MustContext(c), req.Id)
}
