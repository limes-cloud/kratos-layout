package entity

import "github.com/limes-cloud/kratosx/model"

type Task struct {
	Points      *uint32    `json:"points" gorm:"column:points"`
	Title       string     `json:"title" gorm:"column:title"`
	Description string     `json:"description" gorm:"column:description"`
	IsUpdate    *bool      `json:"isUpdate" gorm:"column:is_update"`
	Start       uint32     `json:"start" gorm:"column:start"`
	End         uint32     `json:"end" gorm:"column:end"`
	Config      string     `json:"config" gorm:"column:config"`
	TaskValue   *TaskValue `json:"taskValue"`
	model.BaseTenantModel
}

type TaskValue struct {
	TaskId uint32 `json:"taskId"  gorm:"column:task_id"`
	Value  string `json:"value"  gorm:"column:value"`
	User   *User  `json:"user" gorm:"column:value"`
	model.BaseTenantUserModel
}
