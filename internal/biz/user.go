package biz

import (
	v1 "github.com/go-kratos/kratos-layout/api/v1"
	"github.com/go-kratos/kratos-layout/config"
	"github.com/go-kratos/kratos-layout/internal/biz/types"
	"github.com/go-kratos/kratos-layout/pkg/util"
	"github.com/limes-cloud/kratosx"
	ktypes "github.com/limes-cloud/kratosx/types"
	"gorm.io/gorm"
)

type User struct {
	ktypes.CreateModel
	Name     string `json:"name" gorm:"size:32;not null;comment:姓名"`
	Gender   string `json:"gender" gorm:"size:3;not null;comment:性别"`
	Email    string `json:"email" gorm:"size:128;not null;comment:邮箱"`
	Password string `json:"password" gorm:"size:256;not null;comment:密码"`
}

type UserRepo interface {
	Get(ctx kratosx.Context, id uint32) (*User, error)
	Page(ctx kratosx.Context, options *ktypes.PageOptions) ([]*User, uint32, error)
	Create(ctx kratosx.Context, c *User) (uint32, error)
	Update(ctx kratosx.Context, c *User) error
	Delete(ctx kratosx.Context, uint322 uint32) error
}

type UserUseCase struct {
	config *config.Config
	repo   UserRepo
}

func NewUserUseCase(config *config.Config, repo UserRepo) *UserUseCase {
	return &UserUseCase{config: config, repo: repo}
}

func (u *UserUseCase) Get(ctx kratosx.Context, id uint32) (*User, error) {
	user, err := u.repo.Get(ctx, id)
	if err != nil {
		return nil, v1.NotFoundError()
	}
	return user, nil
}

func (u *UserUseCase) Page(ctx kratosx.Context, req *types.PageUserRequest) ([]*User, uint32, error) {
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
		return nil, 0, v1.DatabaseError()
	}
	return list, total, nil
}

func (u *UserUseCase) Add(ctx kratosx.Context, user *User) (uint32, error) {
	user.Password = util.ParsePwd(user.Password)
	id, err := u.repo.Create(ctx, user)
	if err != nil {
		return 0, v1.DatabaseError()
	}
	return id, nil
}

func (u *UserUseCase) Update(ctx kratosx.Context, user *User) error {
	if err := u.repo.Update(ctx, user); err != nil {
		return v1.DatabaseError()
	}
	return nil
}

func (u *UserUseCase) Delete(ctx kratosx.Context, id uint32) error {
	if err := u.repo.Delete(ctx, id); err != nil {
		return v1.DatabaseError()
	}
	return nil
}
