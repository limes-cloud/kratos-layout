package entity

import (
	"github.com/limes-cloud/kratosx/types"
)

type Information struct {
	types.BaseModel
	Title       string `gorm:"column:title" json:"title"`             // 资讯标题
	Description string `gorm:"column:description" json:"description"` // 资讯简介
	Cover       string `gorm:"column:cover" json:"cover"`             // 资讯封面
	Unit        string `gorm:"column:unit" json:"unit"`               // 发布单位
	Content     string `gorm:"column:content" json:"content"`         // 资讯内容
	IsTop       *bool  `gorm:"column:is_top" json:"isTop"`            // 是否置顶
	Status      *bool  `gorm:"column:status" json:"status"`           // 资讯状态
	Read        int64  `gorm:"column:read" json:"read"`               // 阅读人数
}
