package repository

import (
	"github.com/limes-cloud/kratosx"

	"partyaffairs/internal/domain/entity"
	"partyaffairs/internal/types"
)

type UserRepository interface {
	// GetUser 获取用户信息
	GetUser(ctx kratosx.Context, id uint32) (*entity.User, error)

	// ListUser 获取用户分页列表
	ListUser(ctx kratosx.Context, req *types.ListUserRequest) ([]*entity.User, uint32, error)
}
