package biz

import (
	"github.com/limes-cloud/kratosx"
	ktypes "github.com/limes-cloud/kratosx/types"
	"github.com/limes-cloud/usercenter/api/usercenter/auth"

	"partyaffairs/api/errors"
	"partyaffairs/internal/biz/types"
	"partyaffairs/internal/conf"
)

type VideoClassify struct {
	Name   string  `json:"name" gorm:"not null;size:128;comment:分类名称"`
	Cover  string  `json:"cover"  gorm:"not null;size:128;comment:分类封面"`
	IsTop  *bool   `json:"is_top" gorm:"comment:是否指定"`
	IsTask *bool   `json:"is_task" gorm:"comment:是否任务"`
	Start  *uint32 `json:"start" gorm:"comment:任务开始时间"`
	End    *uint32 `json:"end" gorm:"comment:任务结束时间"`
	Weight uint32  `json:"weight" gorm:"comment:分类权重"`
	Desc   string  `json:"desc" gorm:"not null;size:256;comment:分类描述"`
	ktypes.BaseModel
}

type VideoContent struct {
	Title      string            `json:"title" gorm:"not null;size:128;comment:视频标题"`
	Desc       string            `json:"desc" gorm:"not null;size:256;comment:视频描述"`
	URL        string            `json:"url" gorm:"not null;size:256;comment:视频url"`
	Weight     uint32            `json:"weight" gorm:"comment:序号"`
	Duration   uint32            `json:"duration"  gorm:"not null;comment:持续时间"`
	ClassifyID uint32            `json:"classify_id" gorm:"not null;comment:资源分类"`
	Classify   *VideoClassify    `json:"classify" gorm:"foreignKey:classify_id"`
	Process    *UserVideoProcess `json:"process" gorm:"foreignKey:video_id;constraint:onDelete:cascade"`
	ktypes.BaseModel
}

type UserVideoProcess struct {
	VideoID uint32 `json:"video_id" gorm:"uniqueIndex:vu;not null;comment:视频ID"`
	UserID  uint32 `json:"user_id" gorm:"uniqueIndex:vu;not null;comment:用户ID"`
	Time    uint32 `json:"current" gorm:"not null;comment:当前进度时间"`
	Finish  bool   `json:"finish" gorm:"default:false;comment:是否完成"`
	ktypes.BaseModel
}

type VideoRepo interface {
	PageClassify(ctx kratosx.Context, in *types.PageVideoClassifyRequest) ([]*VideoClassify, uint32, error)
	AddClassify(ctx kratosx.Context, c *VideoClassify) (uint32, error)
	UpdateClassify(ctx kratosx.Context, c *VideoClassify) error
	DeleteClassify(ctx kratosx.Context, uint322 uint32) error
	GetContent(ctx kratosx.Context, id uint32) (*VideoContent, error)
	PageContent(ctx kratosx.Context, req *types.PageVideoContentRequest) ([]*VideoContent, uint32, error)
	AddContent(ctx kratosx.Context, c *VideoContent) (uint32, error)
	UpdateContent(ctx kratosx.Context, c *VideoContent) error
	DeleteContent(ctx kratosx.Context, id uint32) error
	GetUserVideo(ctx kratosx.Context, vid, uid uint32) (*UserVideoProcess, error)
	AddUserVideo(ctx kratosx.Context, c *UserVideoProcess) (uint32, error)
	UpdateUserVideo(ctx kratosx.Context, uv *UserVideoProcess) error
}

type VideoUseCase struct {
	config *conf.Config
	repo   VideoRepo
}

func NewVideoUseCase(config *conf.Config, repo VideoRepo) *VideoUseCase {
	return &VideoUseCase{config: config, repo: repo}
}

// PageClassify 获取全部资源分类
func (u *VideoUseCase) PageClassify(ctx kratosx.Context, in *types.PageVideoClassifyRequest) ([]*VideoClassify, uint32, error) {
	nc, total, err := u.repo.PageClassify(ctx, in)
	if err != nil {
		return nil, 0, errors.DatabaseError()
	}
	return nc, total, nil
}

// AddClassify 添加资源分类信息
func (u *VideoUseCase) AddClassify(ctx kratosx.Context, nc *VideoClassify) (uint32, error) {
	id, err := u.repo.AddClassify(ctx, nc)
	if err != nil {
		return 0, errors.DatabaseError(err.Error())
	}
	return id, nil
}

// UpdateClassify 更新资源分类信息
func (u *VideoUseCase) UpdateClassify(ctx kratosx.Context, nc *VideoClassify) error {
	if err := u.repo.UpdateClassify(ctx, nc); err != nil {
		return errors.DatabaseError(err.Error())
	}
	return nil
}

// DeleteClassify 删除资源分类信息
func (u *VideoUseCase) DeleteClassify(ctx kratosx.Context, id uint32) error {
	if err := u.repo.DeleteClassify(ctx, id); err != nil {
		return errors.DatabaseError(err.Error())
	}
	return nil
}

// GetContent 获取分页资源
func (u *VideoUseCase) GetContent(ctx kratosx.Context, id uint32) (*VideoContent, error) {
	nc, err := u.repo.GetContent(ctx, id)
	if err != nil {
		return nil, errors.DatabaseError()
	}
	return nc, nil
}

// PageContent 获取分页资源
func (u *VideoUseCase) PageContent(ctx kratosx.Context, req *types.PageVideoContentRequest) ([]*VideoContent, uint32, error) {
	info, err := auth.Get(ctx)
	if err == nil {
		req.UserID = &info.UserId
	}
	nc, total, err := u.repo.PageContent(ctx, req)
	if err != nil {
		return nil, 0, errors.DatabaseError()
	}
	return nc, total, nil
}

// AddContent 添加资源
func (u *VideoUseCase) AddContent(ctx kratosx.Context, nc *VideoContent) (uint32, error) {
	id, err := u.repo.AddContent(ctx, nc)
	if err != nil {
		return 0, errors.DatabaseError(err.Error())
	}
	return id, nil
}

// UpdateContent 更新资源
func (u *VideoUseCase) UpdateContent(ctx kratosx.Context, nc *VideoContent) error {
	if err := u.repo.UpdateContent(ctx, nc); err != nil {
		return errors.DatabaseError(err.Error())
	}
	return nil
}

// DeleteContent 删除资源
func (u *VideoUseCase) DeleteContent(ctx kratosx.Context, id uint32) error {
	if err := u.repo.DeleteContent(ctx, id); err != nil {
		return errors.DatabaseError(err.Error())
	}
	return nil
}

// UpdateUserVideoProcess 更新用户观看视频进度
func (u *VideoUseCase) UpdateUserVideoProcess(ctx kratosx.Context, vid uint32, ct uint32) error {
	info, err := auth.Get(ctx)
	if err != nil {
		return errors.AuthInfoError()
	}

	video, err := u.repo.GetContent(ctx, vid)
	if err != nil {
		return errors.DatabaseError()
	}

	uvp := &UserVideoProcess{
		VideoID: vid,
		UserID:  info.UserId,
		Time:    ct,
	}
	if video.Duration-ct <= 5 {
		uvp.Finish = true
	}

	ouv, err := u.repo.GetUserVideo(ctx, vid, info.UserId)
	if err != nil { // 不存在则创建
		if _, err := u.repo.AddUserVideo(ctx, uvp); err != nil {
			return errors.DatabaseError()
		}
		return nil
	}
	uvp.Id = ouv.Id
	if ouv.Time > ct {
		return nil
	}
	if err := u.repo.UpdateUserVideo(ctx, uvp); err != nil {
		return errors.DatabaseError()
	}
	return nil
}
