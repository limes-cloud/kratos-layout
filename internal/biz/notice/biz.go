package notice

import (
	"github.com/limes-cloud/kratosx"

	"github.com/go-kratos/kratos-layout/api/errors"
	"github.com/go-kratos/kratos-layout/internal/conf"
)

type UseCase struct {
	conf *conf.Config
	repo Repo
}

func NewUseCase(config *conf.Config, repo Repo) *UseCase {
	return &UseCase{conf: config, repo: repo}
}

func (u *UseCase) Get(ctx kratosx.Context, id uint32) (*Notice, error) {
	user, err := u.repo.Get(ctx, id)
	if err != nil {
		return nil, errors.NotFoundError()
	}
	return user, nil
}

func (u *UseCase) Page(ctx kratosx.Context, req *PageRequest) ([]*Notice, uint32, error) {
	list, total, err := u.repo.Page(ctx, req)
	if err != nil {
		return nil, 0, errors.DatabaseError()
	}
	return list, total, nil
}

func (u *UseCase) Add(ctx kratosx.Context, user *Notice) (uint32, error) {
	id, err := u.repo.Create(ctx, user)
	if err != nil {
		return 0, errors.DatabaseError()
	}
	return id, nil
}

func (u *UseCase) Update(ctx kratosx.Context, user *Notice) error {
	if err := u.repo.Update(ctx, user); err != nil {
		return errors.DatabaseError()
	}
	return nil
}

func (u *UseCase) Delete(ctx kratosx.Context, id uint32) error {
	if err := u.repo.Delete(ctx, id); err != nil {
		return errors.DatabaseError()
	}
	return nil
}
