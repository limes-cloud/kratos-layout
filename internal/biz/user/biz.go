package user

import (
	"github.com/limes-cloud/kratosx"
	ktypes "github.com/limes-cloud/kratosx/types"
	"gorm.io/gorm"

	"github.com/go-kratos/kratos-layout/api/errors"
	"github.com/go-kratos/kratos-layout/internal/conf"
	"github.com/go-kratos/kratos-layout/internal/pkg/util"
)

type UseCase struct {
	conf *conf.Config
	repo Repo
}

func NewUseCase(config *conf.Config, repo Repo) *UseCase {
	return &UseCase{conf: config, repo: repo}
}

func (u *UseCase) Get(ctx kratosx.Context, id uint32) (*User, error) {
	user, err := u.repo.Get(ctx, id)
	if err != nil {
		return nil, errors.NotFound()
	}
	return user, nil
}

func (u *UseCase) Page(ctx kratosx.Context, req *PageRequest) ([]*User, uint32, error) {
	list, total, err := u.repo.Page(ctx, &ktypes.PageOptions{
		Page:     req.Page,
		PageSize: req.PageSize,
		Scopes: func(db *gorm.DB) *gorm.DB {
			if req.Email != nil {
				db = db.Where("email=?", *req.Email)
			}
			return db
		},
	})
	if err != nil {
		return nil, 0, errors.Database()
	}
	return list, total, nil
}

func (u *UseCase) Add(ctx kratosx.Context, user *User) (uint32, error) {
	user.Password = util.ParsePwd(user.Password)
	id, err := u.repo.Create(ctx, user)
	if err != nil {
		return 0, errors.Database()
	}
	return id, nil
}

func (u *UseCase) Update(ctx kratosx.Context, user *User) error {
	if err := u.repo.Update(ctx, user); err != nil {
		return errors.Database()
	}
	return nil
}

func (u *UseCase) Delete(ctx kratosx.Context, id uint32) error {
	if err := u.repo.Delete(ctx, id); err != nil {
		return errors.Database()
	}
	return nil
}
