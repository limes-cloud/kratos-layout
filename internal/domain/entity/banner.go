package entity

import "github.com/limes-cloud/kratosx/types"

type Banner struct {
	types.BaseModel
	Title  string  `gorm:"column:title" json:"title"`   // 轮播标题
	Src    string  `gorm:"column:src" json:"src"`       // 轮播url
	Path   *string `gorm:"column:path" json:"path"`     // 跳转路径
	Weight *int32  `gorm:"column:weight" json:"weight"` // 权重
	Status *bool   `gorm:"column:status" json:"status"` // 状态
}
