package repository

import (
	"partyaffairs/internal/core"
	"partyaffairs/internal/domain/entity"
	"partyaffairs/internal/types"
)

type Chat interface {
	// GetChatRecord 获取指定的会话记录信息
	GetChatRecord(ctx core.Context, id uint32) (*entity.ChatRecord, error)

	// ListChatRecord 获取会话记录信息列表
	ListChatRecord(ctx core.Context, req *types.ListChatRecordRequest) ([]*entity.ChatRecord, uint32, error)

	// CreateChatRecord 创建会话记录信息
	CreateChatRecord(ctx core.Context, req *entity.ChatRecord) (uint32, error)

	// UpdateChatRecord 更新会话记录信息
	UpdateChatRecord(ctx core.Context, req *entity.ChatRecord) error

	// DeleteChatRecord 删除会话记录信息
	DeleteChatRecord(ctx core.Context, id uint32) error
}
