package entity

import ktypes "github.com/limes-cloud/kratosx/types"

type ResourceClassify struct {
	Name   string `json:"name" gorm:"column:name"`
	Weight uint32 `json:"weight" gorm:"column:weight"`
	ktypes.BaseModel
}

type Resource struct {
	Title         string            `json:"title" gorm:"column:title"`
	Description   string            `json:"description" gorm:"column:description"`
	Url           string            `json:"url" gorm:"column:url"`
	DownloadCount uint32            `json:"downloadCount" gorm:"column:download_count"`
	ClassifyId    uint32            `json:"classifyId" gorm:"column:classify_id"`
	Classify      *ResourceClassify `json:"classify" gorm:"foreignKey:classify_id"`
	ktypes.BaseModel
}
