package service

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/limes-cloud/kratosx"

	"poverty/api/poverty/errors"
	"poverty/internal/conf"
	"poverty/internal/domain/entity"
	"poverty/internal/domain/repository"
	"poverty/internal/types"
)

type ActivityService struct {
	conf *conf.Config
	repo repository.ActivityRepository
	file repository.FileRepository
}

func NewActivityService(conf *conf.Config, repo repository.ActivityRepository, file repository.FileRepository) *ActivityService {
	return &ActivityService{conf: conf, repo: repo, file: file}
}

// GetActivity 获取指定的活动信息
func (srv *ActivityService) GetActivity(ctx kratosx.Context, id uint32) (*entity.Activity, error) {
	ent, err := srv.repo.GetActivity(ctx, id)
	if err != nil {
		return nil, errors.GetError(err.Error())
	}
	ent.Cover = srv.file.GetFileURL(ctx, ent.Cover)

	// 获取通知内容，判断是否具有图片
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(ent.Content))
	if err != nil {
		return nil, errors.GetError(err.Error())
	}
	doc.Find("img").Each(func(i int, selection *goquery.Selection) {
		origin, _ := selection.Attr("data-origin")
		md5, _ := selection.Attr("data-md5")
		if origin == "resource" && md5 != "" {
			// 获取图片连接地址
			src := srv.file.GetFileURL(ctx, md5)
			if src == "" {
				selection.Remove()
			} else {
				selection.SetAttr("src", src)
			}
		}
	})

	out, err := doc.Find("body").Html()
	if err != nil {
		return nil, errors.GetError(err.Error())
	}
	ent.Content = out
	return ent, nil
}

// ListActivity 获取活动信息列表
func (srv *ActivityService) ListActivity(ctx kratosx.Context, req *types.ListActivityRequest) ([]*entity.Activity, uint32, error) {
	list, total, err := srv.repo.ListActivity(ctx, req)
	if err != nil {
		return nil, 0, errors.ListError(err.Error())
	}
	for ind, item := range list {
		list[ind].Cover = srv.file.GetFileURL(ctx, item.Cover)
	}
	return list, total, nil
}

// CreateActivity 创建活动信息
func (srv *ActivityService) CreateActivity(ctx kratosx.Context, req *entity.Activity) (uint32, error) {
	id, err := srv.repo.CreateActivity(ctx, req)
	if err != nil {
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateActivity 更新活动信息
func (srv *ActivityService) UpdateActivity(ctx kratosx.Context, ent *entity.Activity) error {
	if err := srv.repo.UpdateActivity(ctx, ent); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteActivity 删除活动信息
func (srv *ActivityService) DeleteActivity(ctx kratosx.Context, id uint32) error {
	if err := srv.repo.DeleteActivity(ctx, id); err != nil {
		return errors.DeleteError(err.Error())
	}
	return nil
}
