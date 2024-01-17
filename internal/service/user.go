package service

import (
	"context"
	v1 "github.com/go-kratos/kratos-layout/api/v1"
	"github.com/go-kratos/kratos-layout/internal/biz"
	"github.com/go-kratos/kratos-layout/internal/biz/types"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jinzhu/copier"
	"github.com/limes-cloud/kratosx"
)

func (s *Service) GetUser(ctx context.Context, req *v1.GetUserRequest) (*v1.GetUserReply, error) {
	user, err := s.user.Get(kratosx.MustContext(ctx), req.Id)
	if err != nil {
		return nil, err
	}
	reply := v1.GetUserReply{}
	if err := copier.Copy(&reply, user); err != nil {
		return nil, v1.TransformErrorFormat(err.Error())
	}
	return &reply, nil
}

func (s *Service) PageUser(ctx context.Context, req *v1.PageUserRequest) (*v1.PageUserReply, error) {
	bizReq := types.PageUserRequest{}
	if err := copier.Copy(&bizReq, req); err != nil {
		return nil, v1.TransformErrorFormat(err.Error())
	}

	list, total, err := s.user.Page(kratosx.MustContext(ctx), &bizReq)
	if err != nil {
		return nil, err
	}

	reply := v1.PageUserReply{Total: total}
	if err := copier.Copy(&reply.List, list); err != nil {
		return nil, v1.TransformErrorFormat(err.Error())
	}

	return &reply, nil
}

func (s *Service) AddUser(ctx context.Context, req *v1.AddUserRequest) (*v1.AddUserReply, error) {
	bizReq := biz.User{}
	if err := copier.Copy(&bizReq, req); err != nil {
		return nil, v1.TransformErrorFormat(err.Error())
	}

	id, err := s.user.Add(kratosx.MustContext(ctx), &bizReq)
	if err != nil {
		return nil, err
	}

	return &v1.AddUserReply{Id: id}, nil
}

func (s *Service) UpdateUser(ctx context.Context, req *v1.UpdateUserRequest) (*empty.Empty, error) {
	bizReq := biz.User{}
	if err := copier.Copy(&bizReq, req); err != nil {
		return nil, v1.TransformErrorFormat(err.Error())
	}

	if err := s.user.Update(kratosx.MustContext(ctx), &bizReq); err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *Service) DeleteUser(ctx context.Context, req *v1.DeleteUserRequest) (*empty.Empty, error) {
	return nil, s.user.Delete(kratosx.MustContext(ctx), req.Id)
}
