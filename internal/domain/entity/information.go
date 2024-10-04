package entity

import (
	"github.com/limes-cloud/kratosx/types"
)

type InformationClassify struct {
	Name   string `json:"name" gorm:"column:name"`
	Weight uint32 `json:"weight" gorm:"column:weight"`
	types.BaseModel
}

type Information struct {
	types.BaseModel
	ClassifyId  uint32               `gorm:"column:classify_id" json:"classifyId"`
	Title       string               `gorm:"column:title" json:"title"`             // 资讯标题
	Description string               `gorm:"column:description" json:"description"` // 资讯简介
	Cover       string               `gorm:"column:cover" json:"cover"`             // 资讯封面
	Unit        string               `gorm:"column:unit" json:"unit"`               // 发布单位
	Content     string               `gorm:"column:content" json:"content"`         // 资讯内容
	IsTop       *bool                `gorm:"column:is_top" json:"isTop"`            // 是否置顶
	Status      *bool                `gorm:"column:status" json:"status"`           // 资讯状态
	Read        int64                `gorm:"column:read" json:"read"`               // 阅读人数
	Classify    *InformationClassify `json:"classify" gorm:"foreignKey:classify_id"`
}
