package entity

import "github.com/limes-cloud/kratosx/model"

type Banner struct {
	model.BaseTenantModel
	Title  string  `gorm:"column:title" json:"title"`   // 轮播标题
	Key    string  `gorm:"column:key" json:"key"`       // 轮播key
	Path   *string `gorm:"column:path" json:"path"`     // 跳转路径
	Weight *int32  `gorm:"column:weight" json:"weight"` // 权重
	Status *bool   `gorm:"column:status" json:"status"` // 状态
}
