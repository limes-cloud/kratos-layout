package model

import (
	"github.com/limes-cloud/kratosx/types"
)

type Dictionary struct {
	Keyword     string  `json:"keyword" gorm:"column:keyword"`
	Name        string  `json:"name" gorm:"column:name"`
	Description *string `json:"description" gorm:"column:description"`
	types.BaseModel
}
