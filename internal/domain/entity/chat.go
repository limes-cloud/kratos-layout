package entity

import (
	"github.com/limes-cloud/kratosx/types"
)

type ChatRecord struct {
	types.CreateModel
	UserId     uint32 `gorm:"column:user_id" json:"userId"`       // 用户id
	SessionId  string `gorm:"column:session_id" json:"sessionId"` // 会话id
	Message    string `gorm:"column:message" json:"message"`      // 会话信息
	Type       string `gorm:"column:type" json:"type"`            // 会话类型
	UserName   string `gorm:"-" son:"userName"`
	UserAvatar string `gorm:"-" son:"userAvatar"`
}
