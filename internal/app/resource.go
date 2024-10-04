package app

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx"
	ktypes "github.com/limes-cloud/kratosx/types"

	pb "partyaffairs/api/partyaffairs/resource/v1"
	"partyaffairs/internal/conf"
	"partyaffairs/internal/domain/entity"
	"partyaffairs/internal/domain/service"
	"partyaffairs/internal/infra/dbs"
	"partyaffairs/internal/infra/rpc"
	"partyaffairs/internal/types"
)

type Resource struct {
	pb.UnimplementedResourceServer
	srv *service.ResourceService
}

func NewResource(conf *conf.Config) *Resource {
	return &Resource{
		srv: service.NewResourceService(conf, dbs.NewResource(), rpc.NewFile()),
	}
}

func init() {
	register(func(c *conf.Config, hs *http.Server) {
		srv := NewResource(c)
		pb.RegisterResourceHTTPServer(hs, srv)
	})
}

// ListResourceClassify 获取任务分组列表
func (s *Resource) ListResourceClassify(c context.Context, _ *pb.ListResourceClassifyRequest) (*pb.ListResourceClassifyReply, error) {
	list, err := s.srv.ListResourceClassify(kratosx.MustContext(c))
	if err != nil {
		return nil, err
	}
	reply := pb.ListResourceClassifyReply{}
	for _, item := range list {
		reply.List = append(reply.List, &pb.ListResourceClassifyReply_ResourceClassify{
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
func (s *Resource) CreateResourceClassify(c context.Context, req *pb.CreateResourceClassifyRequest) (*pb.CreateResourceClassifyReply, error) {
	id, err := s.srv.CreateResourceClassify(kratosx.MustContext(c), &entity.ResourceClassify{
		Name:   req.Name,
		Weight: req.Weight,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateResourceClassifyReply{Id: id}, nil
}

// UpdateResourceClassify 更新任务分组
func (s *Resource) UpdateResourceClassify(c context.Context, req *pb.UpdateResourceClassifyRequest) (*pb.UpdateResourceClassifyReply, error) {
	if err := s.srv.UpdateResourceClassify(kratosx.MustContext(c), &entity.ResourceClassify{
		BaseModel: ktypes.BaseModel{Id: req.Id},
		Name:      req.Name,
		Weight:    req.Weight,
	}); err != nil {
		return nil, err
	}
	return &pb.UpdateResourceClassifyReply{}, nil
}

// DeleteResourceClassify 删除任务分组
func (s *Resource) DeleteResourceClassify(c context.Context, req *pb.DeleteResourceClassifyRequest) (*pb.DeleteResourceClassifyReply, error) {
	err := s.srv.DeleteResourceClassify(kratosx.MustContext(c), req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteResourceClassifyReply{}, nil
}

// GetResource 获取指定的资讯信息
func (s *Resource) GetResource(c context.Context, req *pb.GetResourceRequest) (*pb.GetResourceReply, error) {
	resource, err := s.srv.GetResource(kratosx.MustContext(c), req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetResourceReply{
		Id:            resource.Id,
		ClassifyId:    resource.ClassifyId,
		Title:         resource.Title,
		Description:   resource.Description,
		Url:           resource.Url,
		DownloadCount: resource.DownloadCount,
		Classify: &pb.GetResourceReply_Classify{
			Id:   resource.Classify.Id,
			Name: resource.Classify.Name,
		},
		CreatedAt: uint32(resource.CreatedAt),
		UpdatedAt: uint32(resource.UpdatedAt),
	}, nil
}

// ListResource 获取资讯信息列表
func (s *Resource) ListResource(c context.Context, req *pb.ListResourceRequest) (*pb.ListResourceReply, error) {
	list, total, err := s.srv.ListResource(kratosx.MustContext(c), &types.ListResourceRequest{
		Page:       req.Page,
		PageSize:   req.PageSize,
		Title:      req.Title,
		ClassifyId: req.ClassifyId,
	})
	if err != nil {
		return nil, err
	}

	reply := pb.ListResourceReply{Total: total}
	for _, item := range list {
		reply.List = append(reply.List, &pb.ListResourceReply_Resource{
			Id:            item.Id,
			ClassifyId:    item.ClassifyId,
			Title:         item.Title,
			Description:   item.Description,
			Url:           item.Url,
			DownloadCount: item.DownloadCount,
			Classify: &pb.ListResourceReply_Classify{
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
func (s *Resource) CreateResource(c context.Context, req *pb.CreateResourceRequest) (*pb.CreateResourceReply, error) {
	id, err := s.srv.CreateResource(kratosx.MustContext(c), &entity.Resource{
		ClassifyId:  req.ClassifyId,
		Title:       req.Title,
		Description: req.Description,
		Url:         req.Url,
	})
	if err != nil {
		return nil, err
	}

	return &pb.CreateResourceReply{Id: id}, nil
}

// UpdateResource 更新资讯信息
func (s *Resource) UpdateResource(c context.Context, req *pb.UpdateResourceRequest) (*pb.UpdateResourceReply, error) {
	if err := s.srv.UpdateResource(kratosx.MustContext(c), &entity.Resource{
		BaseModel:   ktypes.BaseModel{Id: req.Id},
		ClassifyId:  req.ClassifyId,
		Title:       req.Title,
		Description: req.Description,
		Url:         req.Url,
	}); err != nil {
		return nil, err
	}
	return &pb.UpdateResourceReply{}, nil
}

// DeleteResource 删除资讯信息
func (s *Resource) DeleteResource(c context.Context, req *pb.DeleteResourceRequest) (*pb.DeleteResourceReply, error) {
	return &pb.DeleteResourceReply{}, s.srv.DeleteResource(kratosx.MustContext(c), req.Id)
}
