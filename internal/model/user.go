package model

import (
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/types"
)

type User struct {
	types.BaseModel
	Name string `json:"name" gorm:"unique;not null;size:32;comment:用户姓名"`
}

type UserClosure struct {
	ID             uint32 `json:"id" gorm:"primary_key;auto_increment;comment:主键ID"`
	Ancestor       uint32 `json:"ancestor" gorm:"NOT NULL;COMMENT:用户id"`
	Descendant     uint32 `json:"descendant" gorm:"NOT NULL;COMMENT:用户id"`
	AncestorUser   User   `json:"ancestor_user" gorm:"foreignkey:ancestor;references:id"`
	DescendantUser User   `json:"descendant_user" gorm:"foreignkey:descendant;references:id"`
}

func (u *User) Create(ctx kratosx.Context) error {
	// read redis
	return ctx.DB().Create(u).Error
}

func (u *User) Delete(ctx kratosx.Context) error {
	return ctx.DB().Delete(u).Error
}
