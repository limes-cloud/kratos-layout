package model

import "github.com/limes-cloud/kratos"

type User struct {
	ID          int64   `json:"id" gorm:"primary_key;autoIncrement;size:32;comment:主键ID"`
	Name        string  `json:"name" gorm:"not null;size:32;comment:用户姓名"`
	Nickname    string  `json:"nickname" gorm:"not null;size:128;comment:用户昵称"`
	Sex         *bool   `json:"sex,omitempty" gorm:"not null;comment:用户性别"`
	Phone       string  `json:"phone" gorm:"not null;size:32;comment:用户电话"`
	Password    string  `json:"password,omitempty" gorm:"not null;->:false;<-:create,update;comment:用户密码"`
	Avatar      string  `json:"avatar" gorm:"not null;size:128;comment:用户头像"`
	Email       string  `json:"email,omitempty" gorm:"not null;type:varbinary(128);comment:用户邮箱"`
	Status      *bool   `json:"status,omitempty" gorm:"not null;comment:用户状态"`
	DisableDesc *string `json:"disable_desc" gorm:"not null;size:128;comment:禁用原因"`
	LastLogin   int64   `json:"last_login" gorm:"comment:最后登陆时间"`
	Operator    string  `json:"operator" gorm:"not null;size:128;comment:操作人员名称"`
	OperatorID  int64   `json:"operator_id" gorm:"not null;size:32;comment:操作人员id"`
	CreatedAt   int64   `json:"created_at,omitempty" gorm:"index;comment:创建时间"`
	UpdatedAt   int64   `json:"updated_at,omitempty" gorm:"index;comment:修改时间"`
}

func (u *User) TableName() string {
	return "tb_system_user"
}

func (u *User) Create(ctx kratos.Context) error {
	return ctx.DB().Create(u).Error
}
