package dbs

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"partyaffairs/internal/core"
	"partyaffairs/internal/domain/entity"
	"partyaffairs/internal/types"
)

type Chat struct {
}

func NewChat() *Chat {
	return &Chat{}
}

// GetChatRecord 获取指定的会话记录信息数据
func (r *Chat) GetChatRecord(ctx core.Context, id uint32) (*entity.ChatRecord, error) {
	var (
		ent = entity.ChatRecord{}
		fs  = []string{"*"}
	)
	db := ctx.DB().Select(fs)
	return &ent, db.First(&ent, id).Error
}

// ListChatRecord 获取会话记录信息列表
func (r *Chat) ListChatRecord(ctx core.Context, req *types.ListChatRecordRequest) ([]*entity.ChatRecord, uint32, error) {
	var (
		list  []*entity.ChatRecord
		fs    = []string{"*"}
		total int64
	)

	db := ctx.DB().Model(entity.ChatRecord{}).Select(fs)

	if req.UserIds != nil {
		db = db.Where("user_id in ?", req.UserIds)
	}
	if req.UserId != nil {
		db = db.Where("user_id = ?", *req.UserId)
	}
	if req.SessionId != nil {
		db = db.Where("session_id = ?", *req.SessionId)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	db = db.Offset(int((req.Page - 1) * req.PageSize)).Limit(int(req.PageSize))

	if req.OrderBy == nil || *req.OrderBy == "" {
		req.OrderBy = proto.String("id")
	}
	if req.Order == nil || *req.Order == "" {
		req.Order = proto.String("asc")
	}

	db = db.Order(fmt.Sprintf("%s %s", *req.OrderBy, *req.Order))
	if *req.OrderBy != "id" {
		db = db.Order("id asc")
	}

	return list, uint32(total), db.Find(&list).Error
}

// CreateChatRecord 创建会话记录信息数据
func (r *Chat) CreateChatRecord(ctx core.Context, ent *entity.ChatRecord) (uint32, error) {
	return ent.Id, ctx.DB().Create(ent).Error
}

// UpdateChatRecord 更新会话记录信息数据
func (r *Chat) UpdateChatRecord(ctx core.Context, ent *entity.ChatRecord) error {
	return ctx.DB().Updates(ent).Error
}

// DeleteChatRecord 删除会话记录信息数据
func (r *Chat) DeleteChatRecord(ctx core.Context, id uint32) error {
	return ctx.DB().Delete(&entity.ChatRecord{}, id).Error
}
