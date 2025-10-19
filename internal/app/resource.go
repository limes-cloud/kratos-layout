package app

import (
	"context"
	"github.com/limes-cloud/kratosx/model"
	"partyaffairs/api/resource"
	"partyaffairs/internal/core"

	"github.com/go-kratos/kratos/v2/transport/http"
	"partyaffairs/internal/domain/entity"
	"partyaffairs/internal/domain/service"
	"partyaffairs/internal/infra/dbs"
	"partyaffairs/internal/infra/rpc"
	"partyaffairs/internal/types"
)

type Resource struct {
	resource.UnimplementedResourceServer
	srv *service.ResourceService
}

func NewResource() *Resource {
	return &Resource{
		srv: service.NewResourceService(dbs.NewResource(), rpc.NewFile()),
	}
}

func init() {
	register(func(hs *http.Server) {
		srv := NewResource()
		resource.RegisterResourceHTTPServer(hs, srv)
	})
}

// ListResourceClassify 获取任务分组列表
func (s *Resource) ListResourceClassify(c context.Context, _ *resource.ListResourceClassifyRequest) (*resource.ListResourceClassifyReply, error) {
	list, err := s.srv.ListResourceClassify(core.MustContext(c))
	if err != nil {
		return nil, err
	}
	reply := resource.ListResourceClassifyReply{}
	for _, item := range list {
		reply.List = append(reply.List, &resource.ListResourceClassifyReply_ResourceClassify{
			Id:        item.Id,
			Name:      item.Name,
			Weight:    item.Weight,
			CreatedAt: uint32(item.CreatedAt),
			UpdatedAt: uint32(item.UpdatedAt),
		})
	}
	return &reply, nil
}

// CreateResourceClassify 创建任务分组
func (s *Resource) CreateResourceClassify(c context.Context, req *resource.CreateResourceClassifyRequest) (*resource.CreateResourceClassifyReply, error) {
	id, err := s.srv.CreateResourceClassify(core.MustContext(c), &entity.ResourceClassify{
		Name:   req.Name,
		Weight: req.Weight,
	})
	if err != nil {
		return nil, err
	}
	return &resource.CreateResourceClassifyReply{Id: id}, nil
}

// UpdateResourceClassify 更新任务分组
func (s *Resource) UpdateResourceClassify(c context.Context, req *resource.UpdateResourceClassifyRequest) (*resource.UpdateResourceClassifyReply, error) {
	if err := s.srv.UpdateResourceClassify(core.MustContext(c), &entity.ResourceClassify{
		BaseTenantModel: model.BaseTenantModel{Id: req.Id},

		Name:   req.Name,
		Weight: req.Weight,
	}); err != nil {
		return nil, err
	}
	return &resource.UpdateResourceClassifyReply{}, nil
}

// DeleteResourceClassify 删除任务分组
func (s *Resource) DeleteResourceClassify(c context.Context, req *resource.DeleteResourceClassifyRequest) (*resource.DeleteResourceClassifyReply, error) {
	err := s.srv.DeleteResourceClassify(core.MustContext(c), req.Id)
	if err != nil {
		return nil, err
	}
	return &resource.DeleteResourceClassifyReply{}, nil
}

// GetResource 获取指定的资讯信息
func (s *Resource) GetResource(c context.Context, req *resource.GetResourceRequest) (*resource.GetResourceReply, error) {
	res, err := s.srv.GetResource(core.MustContext(c), req.Id)
	if err != nil {
		return nil, err
	}

	return &resource.GetResourceReply{
		Id:            res.Id,
		ClassifyId:    res.ClassifyId,
		Title:         res.Title,
		Description:   res.Description,
		Key:           res.Key,
		DownloadCount: res.DownloadCount,
		Classify: &resource.GetResourceReply_Classify{
			Id:   res.Classify.Id,
			Name: res.Classify.Name,
		},
		CreatedAt: uint32(res.CreatedAt),
		UpdatedAt: uint32(res.UpdatedAt),
	}, nil
}

// ListResource 获取资讯信息列表
func (s *Resource) ListResource(c context.Context, req *resource.ListResourceRequest) (*resource.ListResourceReply, error) {
	list, total, err := s.srv.ListResource(core.MustContext(c), &types.ListResourceRequest{
		Page:       req.Page,
		PageSize:   req.PageSize,
		Title:      req.Title,
		ClassifyId: req.ClassifyId,
	})
	if err != nil {
		return nil, err
	}

	reply := resource.ListResourceReply{Total: total}
	for _, item := range list {
		reply.List = append(reply.List, &resource.ListResourceReply_Resource{
			Id:            item.Id,
			ClassifyId:    item.ClassifyId,
			Title:         item.Title,
			Description:   item.Description,
			Key:           item.Key,
			DownloadCount: item.DownloadCount,
			Classify: &resource.ListResourceReply_Classify{
				Id:   item.Classify.Id,
				Name: item.Classify.Name,
			},
			CreatedAt: uint32(item.CreatedAt),
			UpdatedAt: uint32(item.UpdatedAt),
		})
	}
	return &reply, nil
}

// CreateResource 创建资讯信息
func (s *Resource) CreateResource(c context.Context, req *resource.CreateResourceRequest) (*resource.CreateResourceReply, error) {
	id, err := s.srv.CreateResource(core.MustContext(c), &entity.Resource{
		ClassifyId:  req.ClassifyId,
		Title:       req.Title,
		Description: req.Description,
		Key:         req.Key,
	})
	if err != nil {
		return nil, err
	}

	return &resource.CreateResourceReply{Id: id}, nil
}

// UpdateResource 更新资讯信息
func (s *Resource) UpdateResource(c context.Context, req *resource.UpdateResourceRequest) (*resource.UpdateResourceReply, error) {
	if err := s.srv.UpdateResource(core.MustContext(c), &entity.Resource{
		BaseTenantModel: model.BaseTenantModel{Id: req.Id},
		ClassifyId:      req.ClassifyId,
		Title:           req.Title,
		Description:     req.Description,
		Key:             req.Key,
	}); err != nil {
		return nil, err
	}
	return &resource.UpdateResourceReply{}, nil
}

// DeleteResource 删除资讯信息
func (s *Resource) DeleteResource(c context.Context, req *resource.DeleteResourceRequest) (*resource.DeleteResourceReply, error) {
	return &resource.DeleteResourceReply{}, s.srv.DeleteResource(core.MustContext(c), req.Id)
}
