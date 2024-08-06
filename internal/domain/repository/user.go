package repository

import (
	"github.com/limes-cloud/kratosx"
	"poverty/internal/domain/entity"
)

type UserRepository interface {
	GetUser(ctx kratosx.Context, id uint32) (*entity.User, error)
	ListUserByName(ctx kratosx.Context, name string) ([]*entity.User, error)
}
