package entity

import "github.com/limes-cloud/kratosx/model"

type Notice struct {
	model.BaseTenantModel
	Title       string `gorm:"column:title" json:"title"`             // 通知标题
	Description string `gorm:"column:description" json:"description"` // 通知简介
	Unit        string `gorm:"column:unit" json:"unit"`               // 发布单位
	Content     string `gorm:"column:content" json:"content"`         // 通知内容
	IsTop       *bool  `gorm:"column:is_top" json:"isTop"`            // 是否置顶
	Status      *bool  `gorm:"column:status" json:"status"`           // 通知状态
}

type NoticeUser struct {
	NoticeId uint32 `gorm:"column:notice_id" json:"noticeId"`
	IsRead   *bool  `gorm:"column:is_read" json:"isRead"`
	model.BaseTenantUserModel
}
