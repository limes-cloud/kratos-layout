package dbs

import (
	"fmt"

	"github.com/limes-cloud/kratosx"
	"google.golang.org/protobuf/proto"

	"poverty/internal/domain/entity"
	"poverty/internal/types"
)

type CommentInfra struct {
}

func NewCommentInfra() *CommentInfra {
	return &CommentInfra{}
}

// GetComment 获取指定的留言信息数据
func (r *CommentInfra) GetComment(ctx kratosx.Context, id uint32) (*entity.Comment, error) {
	var (
		ent = entity.Comment{}
		fs  = []string{"*"}
	)
	db := ctx.DB().Select(fs)
	return &ent, db.First(&ent, id).Error
}

// ListComment 获取留言信息列表
func (r *CommentInfra) ListComment(ctx kratosx.Context, req *types.ListCommentRequest) ([]*entity.Comment, uint32, error) {
	var (
		list  []*entity.Comment
		fs    = []string{"*"}
		total int64
	)

	db := ctx.DB().Model(entity.Comment{})
	db = db.Select(fs)

	if req.UserIds != nil {
		db = db.Where("user_id in ?", req.UserIds)
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

// CreateComment 创建留言信息数据
func (r *CommentInfra) CreateComment(ctx kratosx.Context, ent *entity.Comment) (uint32, error) {
	return ent.Id, ctx.DB().Create(ent).Error
}

// UpdateComment 更新留言信息数据
func (r *CommentInfra) UpdateComment(ctx kratosx.Context, ent *entity.Comment) error {
	return ctx.DB().Updates(ent).Error
}

// DeleteComment 删除留言信息数据
func (r *CommentInfra) DeleteComment(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(&entity.Comment{}, id).Error
}
