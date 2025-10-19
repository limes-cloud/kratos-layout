package repository

import (
	"partyaffairs/internal/core"
	"partyaffairs/internal/domain/entity"
	"partyaffairs/internal/types"
)

type UserRepository interface {
	// GetUser 获取用户信息
	GetUser(ctx core.Context, id uint32) (*entity.User, error)

	// ListUser 获取用户分页列表
	ListUser(ctx core.Context, req *types.ListUserRequest) ([]*entity.User, uint32, error)

	ListUserMap(ctx core.Context, req *types.ListUserRequest) (map[uint32]*entity.User, error)
}
