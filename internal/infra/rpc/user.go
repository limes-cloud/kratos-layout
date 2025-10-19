package rpc

import (
	"github.com/limes-cloud/kratosx/pkg/value"
	"github.com/limes-cloud/manager/api/user"
	"google.golang.org/protobuf/proto"
	"partyaffairs/internal/core"

	"partyaffairs/api/errors"
	"partyaffairs/internal/domain/entity"
	"partyaffairs/internal/types"
)

type User struct {
}

func NewUser() *User {
	return &User{}
}

func (i User) client(ctx core.Context) (user.UserClient, error) {
	conn, err := core.MustContext(ctx).GrpcConn("Manager")
	if err != nil {
		return nil, errors.ResourceServerError()
	}
	return user.NewUserClient(conn), nil
}

func (i User) GetUser(ctx core.Context, id uint32) (*entity.User, error) {
	client, err := i.client(ctx)
	if err != nil {
		return nil, err
	}
	reply, err := client.GetUser(ctx, &user.GetUserRequest{Id: &id})
	if err != nil {
		return nil, err
	}
	return &entity.User{
		Id:       reply.Id,
		Username: reply.Username,
		Nickname: reply.Nickname,
		Avatar:   reply.Avatar,
	}, nil
}

func (i User) ListUser(ctx core.Context, req *types.ListUserRequest) ([]*entity.User, uint32, error) {
	client, err := i.client(ctx)
	if err != nil {
		return nil, 0, err
	}
	reply, err := client.ListUser(ctx, &user.ListUserRequest{
		Page:     req.Page,
		PageSize: req.PageSize,
		Status:   proto.Bool(true),
		InIds:    req.In,
	})
	if err != nil {
		return nil, 0, err
	}
	var ents []*entity.User
	if err := value.Transform(reply.List, &ents); err != nil {
		return nil, 0, err
	}
	return ents, reply.Total, nil
}

func (i User) ListUserMap(ctx core.Context, req *types.ListUserRequest) (map[uint32]*entity.User, error) {
	list, _, err := i.ListUser(ctx, req)
	if err != nil {
		return nil, err
	}
	var m = make(map[uint32]*entity.User, len(list))
	for _, v := range list {
		m[v.Id] = v
	}
	return m, nil
}
