package entity

import (
	"github.com/limes-cloud/kratosx/types"
)

type Activity struct {
	types.BaseModel
	Title       string `gorm:"column:title" json:"title"`             // 活动标题
	Description string `gorm:"column:description" json:"description"` // 活动简介
	Cover       string `gorm:"column:cover" json:"cover"`             // 活动封面
	Unit        string `gorm:"column:unit" json:"unit"`               // 发布单位
	Content     string `gorm:"column:content" json:"content"`         // 活动内容
	IsTop       *bool  `gorm:"column:is_top" json:"isTop"`            // 是否置顶
	Status      *bool  `gorm:"column:status" json:"status"`           // 活动状态
	Read        int64  `gorm:"column:read" json:"read"`               // 阅读人数
}
