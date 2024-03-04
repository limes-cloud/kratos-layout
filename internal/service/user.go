package service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jinzhu/copier"
	"github.com/limes-cloud/kratosx"

	"github.com/go-kratos/kratos-layout/api/errors"
	"github.com/go-kratos/kratos-layout/api/userpb"
	userbiz "github.com/go-kratos/kratos-layout/internal/biz/user"
	"github.com/go-kratos/kratos-layout/internal/conf"
	userdata "github.com/go-kratos/kratos-layout/internal/data/user"
)

type UserService struct {
	userpb.UnimplementedServiceServer
	user *userbiz.UseCase
}

func NewUser(conf *conf.Config) *UserService {
	return &UserService{
		user: userbiz.NewUseCase(conf, userdata.NewRepo()),
	}
}

func (s *UserService) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.User, error) {
	user, err := s.user.Get(kratosx.MustContext(ctx), req.Id)
	if err != nil {
		return nil, err
	}
	reply := userpb.User{}
	if err := copier.Copy(&reply, user); err != nil {
		return nil, errors.TransformFormat(err.Error())
	}
	return &reply, nil
}

func (s *UserService) PageUser(ctx context.Context, req *userpb.PageUserRequest) (*userpb.PageUserReply, error) {
	bizReq := userbiz.PageRequest{}
	if err := copier.Copy(&bizReq, req); err != nil {
		return nil, errors.TransformFormat(err.Error())
	}

	list, total, err := s.user.Page(kratosx.MustContext(ctx), &bizReq)
	if err != nil {
		return nil, err
	}

	reply := userpb.PageUserReply{Total: total}
	if err := copier.Copy(&reply.List, list); err != nil {
		return nil, errors.TransformFormat(err.Error())
	}

	return &reply, nil
}

func (s *UserService) AddUser(ctx context.Context, req *userpb.AddUserRequest) (*userpb.AddUserReply, error) {
	bizReq := userbiz.User{}
	if err := copier.Copy(&bizReq, req); err != nil {
		return nil, errors.TransformFormat(err.Error())
	}

	id, err := s.user.Add(kratosx.MustContext(ctx), &bizReq)
	if err != nil {
		return nil, err
	}

	return &userpb.AddUserReply{Id: id}, nil
}

func (s *UserService) UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*empty.Empty, error) {
	bizReq := userbiz.User{}
	if err := copier.Copy(&bizReq, req); err != nil {
		return nil, errors.TransformFormat(err.Error())
	}

	if err := s.user.Update(kratosx.MustContext(ctx), &bizReq); err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *UserService) DeleteUser(ctx context.Context, req *userpb.DeleteUserRequest) (*empty.Empty, error) {
	return nil, s.user.Delete(kratosx.MustContext(ctx), req.Id)
}
