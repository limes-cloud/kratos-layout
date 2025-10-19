package entity

import (
	"github.com/limes-cloud/kratosx/model"
)

type ChatRecord struct {
	SessionId string `gorm:"column:session_id" json:"sessionId"` // 会话id
	Message   string `gorm:"column:message" json:"message"`      // 会话信息
	Type      string `gorm:"column:type" json:"type"`            // 会话类型
	Username  string `gorm:"-" son:"username"`
	Avatar    string `gorm:"-" son:"avatar"`
	model.CreateTenantUserModel
}
