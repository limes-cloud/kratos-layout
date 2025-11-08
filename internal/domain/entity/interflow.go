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

type InterflowClassify struct {
	Name        string  `json:"name" gorm:"column:name"`
	Weight      uint32  `json:"weight" gorm:"column:weight"`
	Description string  `json:"description" gorm:"column:description"`
	Person      string  `json:"person" gorm:"column:person"`
	Users       []*User `json:"users" gorm:"-"`
	model.BaseTenantModel
}
