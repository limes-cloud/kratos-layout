package notice

import (
	ktypes "github.com/limes-cloud/kratosx/types"

	"github.com/go-kratos/kratos-layout/internal/biz/user"
)

type Notice struct {
	Title      string      `json:"title" gorm:"not null;size:128;comment:公告标题"`
	Desc       string      `json:"desc" gorm:"not null;size:512;comment:公告描述"`
	Unit       string      `json:"unit" gorm:"not null;size:128;comment:发布单位"`
	Content    string      `json:"content" gorm:"not null;type:text;comment:轮播权重"`
	IsTop      *bool       `json:"is_top" gorm:"default false;comment:是否置顶"`
	NoticeUser *NoticeUser `json:"notice_user" gorm:"constraint:onDelete:cascade"`
	ktypes.BaseModel
}

type NoticeUser struct {
	UserID   uint32     `json:"user_id" gorm:"uniqueIndex:un;not null;comment:用户id"`
	NoticeID uint32     `json:"notice_id" gorm:"uniqueIndex:un;not null;comment:通知id"`
	User     *user.User `json:"user" gorm:"constraint:onDelete:cascade"`
	ktypes.CreateModel
}
