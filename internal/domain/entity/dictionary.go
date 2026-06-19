package entity

import "github.com/limes-cloud/kratosx/model"

type Dictionary struct {
	Keyword     string  `json:"keyword" gorm:"column:keyword"`
	Name        string  `json:"name" gorm:"column:name"`
	Description *string `json:"description" gorm:"column:description"`
	model.BaseModel
}
