package entity

import "github.com/limes-cloud/kratosx/model"

type Entity struct {
	AppId    uint32         `json:"appId" gorm:"column:app_id"`
	Database string         `json:"database" gorm:"column:database"`
	Name     string         `json:"name" gorm:"column:name"`
	Comment  string         `json:"comment" gorm:"column:comment"`
	Fields   []*EntityField `json:"fields" gorm:"foreignKey:EntityId;references:Id"`
	model.BaseModel
}

type EntityField struct {
	EntityId uint32 `json:"entityId" gorm:"column:entity_id"`
	Name     string `json:"name" gorm:"column:name"`
	Comment  string `json:"comment" gorm:"column:comment"`
	Index    int    `json:"index" gorm:"column:index"`
	Entity   Entity `json:"entity" gorm:"foreignKey:EntityId;references:Id"`
	model.BaseModel
}

type EntityRule struct {
	EntityId    uint32 `json:"entityId" gorm:"column:entity_id"`
	Name        string `json:"name" gorm:"column:name"`
	Status      *bool  `json:"status" gorm:"column:status"`
	Expression  string `json:"expression" gorm:"column:expression"`
	Description string `json:"description" gorm:"column:description"`
	Entity      Entity `json:"entity" gorm:"foreignKey:EntityId;references:Id"`
	model.BaseModel
}
