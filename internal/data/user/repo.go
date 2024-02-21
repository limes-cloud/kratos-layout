package user

import (
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/types"

	bizUser "github.com/go-kratos/kratos-layout/internal/biz/user"
)

type repo struct {
}

func NewRepo() bizUser.Repo {
	return &repo{}
}

func (u repo) Create(ctx kratosx.Context, user *bizUser.User) (uint32, error) {
	return user.ID, ctx.DB().Create(user).Error
}

func (u repo) Update(ctx kratosx.Context, user *bizUser.User) error {
	return ctx.DB().Model(bizUser.User{}).Updates(user).Error
}

func (u repo) Delete(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(bizUser.User{}, "id=?", id).Error
}

func (u repo) Get(ctx kratosx.Context, id uint32) (*bizUser.User, error) {
	var ins bizUser.User
	return &ins, ctx.DB().First(&ins, "id=?", id).Error
}

func (u repo) Page(ctx kratosx.Context, options *types.PageOptions) ([]*bizUser.User, uint32, error) {
	var list []*bizUser.User
	var total int64

	db := ctx.DB().Model(bizUser.User{})
	if options.Scopes != nil {
		db = db.Scopes(options.Scopes)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, uint32(total), err
	}

	db = db.Offset(int((options.Page - 1) * options.PageSize)).Limit(int(options.PageSize))

	return list, uint32(total), db.Find(&list).Error
}
