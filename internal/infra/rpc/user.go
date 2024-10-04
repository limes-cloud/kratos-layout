package rpc

import (
	user "github.com/limes-cloud/application/api/application/user/v1"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"
	"google.golang.org/protobuf/proto"

	"partyaffairs/api/partyaffairs/errors"
	"partyaffairs/internal/domain/entity"
	"partyaffairs/internal/types"
)

const (
	Application = "Application"
)

var (
	ClientApp = "PartyAffairs"
)

type User struct {
}

func NewUser() *User {
	return &User{}
}

func (i User) client(ctx kratosx.Context) (user.UserClient, error) {
	conn, err := kratosx.MustContext(ctx).GrpcConn(Application)
	if err != nil {
		return nil, errors.ResourceServerError()
	}
	return user.NewUserClient(conn), nil
}

func (i User) GetUser(ctx kratosx.Context, id uint32) (*entity.User, error) {
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

func (i User) ListUser(ctx kratosx.Context, req *types.ListUserRequest) ([]*entity.User, uint32, error) {
	client, err := i.client(ctx)
	if err != nil {
		return nil, 0, err
	}
	reply, err := client.ListUser(ctx, &user.ListUserRequest{
		Page:     req.Page,
		PageSize: req.PageSize,
		Status:   proto.Bool(true),
		App:      &ClientApp,
		NotInIds: req.NotIn,
	})
	if err != nil {
		return nil, 0, err
	}
	var ents []*entity.User
	if err := valx.Transform(reply.List, &ents); err != nil {
		return nil, 0, err
	}
	return ents, reply.Total, nil
}
