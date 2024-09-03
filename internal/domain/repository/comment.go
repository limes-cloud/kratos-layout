package repository

import (
	"github.com/limes-cloud/kratosx"

	"poverty/internal/domain/entity"
	"poverty/internal/types"
)

type CommentRepository interface {
	// GetComment 获取指定的留言信息
	GetComment(ctx kratosx.Context, id uint32) (*entity.Comment, error)

	// ListComment 获取留言信息列表
	ListComment(ctx kratosx.Context, req *types.ListCommentRequest) ([]*entity.Comment, uint32, error)

	// CreateComment 创建留言信息
	CreateComment(ctx kratosx.Context, req *entity.Comment) (uint32, error)

	// UpdateComment 更新留言信息
	UpdateComment(ctx kratosx.Context, req *entity.Comment) error

	// DeleteComment 删除留言信息
	DeleteComment(ctx kratosx.Context, id uint32) error
}
