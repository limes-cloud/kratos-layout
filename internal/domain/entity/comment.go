package entity

import (
	"github.com/limes-cloud/kratosx/types"
)

type Comment struct {
	types.CreateModel
	UserId     uint32 `gorm:"column:user_id" json:"userId"`  // 用户id
	Content    string `gorm:"column:content" json:"content"` // 留言
	UserName   string `gorm:"-" son:"userName"`
	UserAvatar string `gorm:"-" son:"userAvatar"`
}
