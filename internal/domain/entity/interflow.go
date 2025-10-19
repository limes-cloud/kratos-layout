package entity

import "github.com/limes-cloud/kratosx/model"

type Interflow struct {
	model.CreateModel
	ID         uint32 `json:"id" gorm:"column:id"`
	FromUserID uint32 `json:"fromUserId" gorm:"column:from_user_id"`
	ToUserID   uint32 `json:"toUserId" gorm:"column:to_user_id"`
	Content    string `json:"content" gorm:"column:content"`
	Type       string `json:"type" gorm:"column:type"`
}

type InterflowPerson struct {
	model.CreateModel
	UserID   uint32 `json:"userId" gorm:"column:user_id"`
	Username string `json:"username" gorm:"column:-"`
	Nickname string `json:"nickname" gorm:"column:-"`
	Avatar   string `json:"avatar" gorm:"column:-"`
}
