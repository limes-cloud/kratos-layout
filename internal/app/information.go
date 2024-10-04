package app

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx"
	ktypes "github.com/limes-cloud/kratosx/types"

	pb "partyaffairs/api/partyaffairs/information/v1"
	"partyaffairs/internal/conf"
	"partyaffairs/internal/domain/entity"
	"partyaffairs/internal/domain/service"
	"partyaffairs/internal/infra/dbs"
	"partyaffairs/internal/infra/rpc"
	"partyaffairs/internal/types"
)

type Information struct {
	pb.UnimplementedInformationServer
	srv *service.InformationService
}

func NewInformation(conf *conf.Config) *Information {
	return &Information{
		srv: service.NewInformationService(conf, dbs.NewInformation(), rpc.NewFile()),
	}
}

func init() {
	register(func(c *conf.Config, hs *http.Server) {
		srv := NewInformation(c)
		pb.RegisterInformationHTTPServer(hs, srv)
	})
}

// ListInformationClassify 获取任务分组列表
func (s *Information) ListInformationClassify(c context.Context, _ *pb.ListInformationClassifyRequest) (*pb.ListInformationClassifyReply, error) {
	list, err := s.srv.ListInformationClassify(kratosx.MustContext(c))
	if err != nil {
		return nil, err
	}
	reply := pb.ListInformationClassifyReply{}
	for _, item := range list {
		reply.List = append(reply.List, &pb.ListInformationClassifyReply_InformationClassify{
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
func (s *Information) CreateInformationClassify(c context.Context, req *pb.CreateInformationClassifyRequest) (*pb.CreateInformationClassifyReply, error) {
	id, err := s.srv.CreateInformationClassify(kratosx.MustContext(c), &entity.InformationClassify{
		Name:   req.Name,
		Weight: req.Weight,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateInformationClassifyReply{Id: id}, nil
}

// UpdateInformationClassify 更新任务分组
func (s *Information) UpdateInformationClassify(c context.Context, req *pb.UpdateInformationClassifyRequest) (*pb.UpdateInformationClassifyReply, error) {
	if err := s.srv.UpdateInformationClassify(kratosx.MustContext(c), &entity.InformationClassify{
		BaseModel: ktypes.BaseModel{Id: req.Id},
		Name:      req.Name,
		Weight:    req.Weight,
	}); err != nil {
		return nil, err
	}
	return &pb.UpdateInformationClassifyReply{}, nil
}

// DeleteInformationClassify 删除任务分组
func (s *Information) DeleteInformationClassify(c context.Context, req *pb.DeleteInformationClassifyRequest) (*pb.DeleteInformationClassifyReply, error) {
	err := s.srv.DeleteInformationClassify(kratosx.MustContext(c), req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteInformationClassifyReply{}, nil
}

// GetInformation 获取指定的资讯信息
func (s *Information) GetInformation(c context.Context, req *pb.GetInformationRequest) (*pb.GetInformationReply, error) {
	information, err := s.srv.GetInformation(kratosx.MustContext(c), req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetInformationReply{
		Id:          information.Id,
		ClassifyId:  information.ClassifyId,
		Title:       information.Title,
		Description: information.Description,
		Cover:       information.Cover,
		Unit:        information.Unit,
		Content:     information.Content,
		IsTop:       information.IsTop,
		Status:      information.Status,
		Read:        uint32(information.Read),
		Classify: &pb.GetInformationReply_Classify{
			Id:   information.Classify.Id,
			Name: information.Classify.Name,
		},
		CreatedAt: uint32(information.CreatedAt),
		UpdatedAt: uint32(information.UpdatedAt),
	}, nil
}

// ListInformation 获取资讯信息列表
func (s *Information) ListInformation(c context.Context, req *pb.ListInformationRequest) (*pb.ListInformationReply, error) {
	list, total, err := s.srv.ListInformation(kratosx.MustContext(c), &types.ListInformationRequest{
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

	reply := pb.ListInformationReply{Total: total}
	for _, item := range list {
		reply.List = append(reply.List, &pb.ListInformationReply_Information{
			Id:          item.Id,
			ClassifyId:  item.ClassifyId,
			Title:       item.Title,
			Description: item.Description,
			Cover:       item.Cover,
			Unit:        item.Unit,
			IsTop:       item.IsTop,
			Status:      item.Status,
			Read:        uint32(item.Read),
			Classify: &pb.ListInformationReply_Classify{
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
func (s *Information) CreateInformation(c context.Context, req *pb.CreateInformationRequest) (*pb.CreateInformationReply, error) {
	id, err := s.srv.CreateInformation(kratosx.MustContext(c), &entity.Information{
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

	return &pb.CreateInformationReply{Id: id}, nil
}

// UpdateInformation 更新资讯信息
func (s *Information) UpdateInformation(c context.Context, req *pb.UpdateInformationRequest) (*pb.UpdateInformationReply, error) {
	if err := s.srv.UpdateInformation(kratosx.MustContext(c), &entity.Information{
		BaseModel:   ktypes.BaseModel{Id: req.Id},
		ClassifyId:  req.ClassifyId,
		Title:       req.Title,
		Description: req.Description,
		Cover:       req.Cover,
		Unit:        req.Unit,
		Content:     req.Content,
		IsTop:       req.IsTop,
		Status:      req.Status,
	}); err != nil {
		return nil, err
	}
	return &pb.UpdateInformationReply{}, nil
}

// DeleteInformation 删除资讯信息
func (s *Information) DeleteInformation(c context.Context, req *pb.DeleteInformationRequest) (*pb.DeleteInformationReply, error) {
	return &pb.DeleteInformationReply{}, s.srv.DeleteInformation(kratosx.MustContext(c), req.Id)
}
