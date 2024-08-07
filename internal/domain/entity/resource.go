package entity

import (
	"github.com/limes-cloud/kratosx/types"
)

type Resource struct {
	types.BaseModel
	Title  string `gorm:"column:title" json:"title"`   // 文件标题
	Src    string `gorm:"column:src" json:"src"`       // 文件url
	Status *bool  `gorm:"column:status" json:"status"` // 状态
	*File
}
