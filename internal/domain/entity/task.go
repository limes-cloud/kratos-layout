package entity

import ktypes "github.com/limes-cloud/kratosx/types"

type Task struct {
	Title       string     `json:"title" gorm:"column:title"`
	Description string     `json:"description" gorm:"column:description"`
	IsUpdate    *bool      `json:"isUpdate" gorm:"column:is_update"`
	Start       uint32     `json:"start" gorm:"column:start"`
	End         uint32     `json:"end" gorm:"column:end"`
	Config      string     `json:"config" gorm:"column:config"`
	TaskValue   *TaskValue `json:"taskValue"`
	ktypes.BaseModel
}

type TaskValue struct {
	TaskId uint32 `json:"taskId"  gorm:"column:task_id"`
	UserId uint32 `json:"userId" gorm:"column:user_id"`
	Value  string `json:"value"  gorm:"column:value"`
	User   *User  `json:"user" gorm:"column:value"`
	ktypes.BaseModel
}
