package entity

import "github.com/limes-cloud/kratosx/model"

type ResourceClassify struct {
	Name string `json:"name" gorm:"column:name"`
	Logo string `json:"logo" gorm:"column:logo"`
	model.BaseTenantModel
}

type Resource struct {
	Title         string            `json:"title" gorm:"column:title"`
	Description   string            `json:"description" gorm:"column:description"`
	Key           string            `json:"key" gorm:"column:key"`
	DownloadCount uint32            `json:"downloadCount" gorm:"column:download_count"`
	ClassifyId    uint32            `json:"classifyId" gorm:"column:classify_id"`
	Classify      *ResourceClassify `json:"classify" gorm:"foreignKey:classify_id"`
	model.BaseTenantModel
}
