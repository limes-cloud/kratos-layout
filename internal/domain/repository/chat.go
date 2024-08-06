package repository

import (
	"github.com/limes-cloud/kratosx"

	"poverty/internal/domain/entity"
	"poverty/internal/types"
)

type ChatRepository interface {
	// GetChatRecord 获取指定的会话记录信息
	GetChatRecord(ctx kratosx.Context, id uint32) (*entity.ChatRecord, error)

	// ListChatRecord 获取会话记录信息列表
	ListChatRecord(ctx kratosx.Context, req *types.ListChatRecordRequest) ([]*entity.ChatRecord, uint32, error)

	// CreateChatRecord 创建会话记录信息
	CreateChatRecord(ctx kratosx.Context, req *entity.ChatRecord) (uint32, error)

	// UpdateChatRecord 更新会话记录信息
	UpdateChatRecord(ctx kratosx.Context, req *entity.ChatRecord) error

	// DeleteChatRecord 删除会话记录信息
	DeleteChatRecord(ctx kratosx.Context, id uint32) error
}
