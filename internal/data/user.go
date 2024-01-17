package data

import (
	"context"
	"github.com/go-kratos/kratos-layout/internal/biz"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/types"
)

type userRepo struct {
}

func NewUserRepo() biz.UserRepo {
	db := kratosx.MustContext(context.Background()).DB()
	_ = db.Set(
		"gorm:table_options",
		"COMMENT='用户信息' ENGINE=InnoDB CHARSET=utf8mb4",
	).AutoMigrate(biz.User{})

	return &userRepo{}
}

func (u userRepo) Create(ctx kratosx.Context, user *biz.User) (uint32, error) {
	return user.ID, ctx.DB().Create(user).Error
}

func (u userRepo) Update(ctx kratosx.Context, user *biz.User) error {
	return ctx.DB().Model(biz.User{}).Updates(user).Error
}

func (u userRepo) Delete(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(biz.User{}, "id=?", id).Error
}

func (u userRepo) Get(ctx kratosx.Context, id uint32) (*biz.User, error) {
	var ins biz.User
	return &ins, ctx.DB().First(&ins, "id=?", id).Error
}

func (u userRepo) Page(ctx kratosx.Context, options *types.PageOptions) ([]*biz.User, uint32, error) {
	var list []*biz.User
	var total int64

	db := ctx.DB().Model(biz.User{})
	if options.Scopes != nil {
		db = db.Scopes(options.Scopes)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, uint32(total), err
	}

	db = db.Offset(int((options.Page - 1) * options.PageSize)).Limit(int(options.PageSize))

	return list, uint32(total), db.Find(&list).Error
}
