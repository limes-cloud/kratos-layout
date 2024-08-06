package rpc

import (
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"
	user "github.com/limes-cloud/usercenter/api/usercenter/user/v1"
	"poverty/api/poverty/errors"
	"poverty/internal/domain/entity"
	"poverty/internal/domain/repository"
)

const (
	UserCenter = "UserCenter"
)

type UserInfra struct {
}

func NewUserInfra() repository.UserRepository {
	return &UserInfra{}
}

func (i UserInfra) client(ctx kratosx.Context) (user.UserClient, error) {
	conn, err := kratosx.MustContext(ctx).GrpcConn(UserCenter)
	if err != nil {
		return nil, errors.ResourceServerError()
	}
	return user.NewUserClient(conn), nil
}

func (i UserInfra) GetUser(ctx kratosx.Context, id uint32) (*entity.User, error) {
	client, err := i.client(ctx)
	if err != nil {
		return nil, err
	}
	reply, err := client.GetUser(ctx, &user.GetUserRequest{Id: &id})
	if err != nil {
		return nil, err
	}
	return &entity.User{
		Id:        reply.Id,
		Phone:     reply.Phone,
		Email:     reply.Email,
		Username:  reply.Username,
		NickName:  reply.NickName,
		RealName:  reply.RealName,
		Avatar:    reply.Avatar,
		AvatarUrl: reply.AvatarUrl,
	}, nil
}

func (i UserInfra) ListUserByName(ctx kratosx.Context, name string) ([]*entity.User, error) {
	client, err := i.client(ctx)
	if err != nil {
		return nil, err
	}
	reply, err := client.ListUser(ctx, &user.ListUserRequest{Page: 1, PageSize: 50, RealName: &name})
	if err != nil {
		return nil, err
	}
	var ents []*entity.User
	if err := valx.Transform(reply.List, &ents); err != nil {
		return nil, err
	}
	return ents, nil
}
