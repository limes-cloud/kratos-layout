package user

import ktypes "github.com/limes-cloud/kratosx/types"

type User struct {
	ktypes.CreateModel
	Name     string `json:"name" gorm:"size:32;not null;comment:姓名"`
	Gender   string `json:"gender" gorm:"size:3;not null;comment:性别"`
	Email    string `json:"email" gorm:"size:128;not null;comment:邮箱"`
	Password string `json:"password" gorm:"size:256;not null;comment:密码"`
}
