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

type InformationService struct {
	conf *conf.Config
	repo repository.InformationRepository
	file repository.FileRepository
}

func NewInformationService(conf *conf.Config, repo repository.InformationRepository, file repository.FileRepository) *InformationService {
	return &InformationService{conf: conf, repo: repo, file: file}
}

// GetInformation 获取指定的资讯信息
func (srv *InformationService) GetInformation(ctx kratosx.Context, id uint32) (*entity.Information, error) {
	ent, err := srv.repo.GetInformation(ctx, id)
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

// ListInformation 获取资讯信息列表
func (srv *InformationService) ListInformation(ctx kratosx.Context, req *types.ListInformationRequest) ([]*entity.Information, uint32, error) {
	list, total, err := srv.repo.ListInformation(ctx, req)
	if err != nil {
		return nil, 0, errors.ListError(err.Error())
	}
	if err != nil {
		return nil, 0, errors.ListError(err.Error())
	}
	for ind, item := range list {
		list[ind].Cover = srv.file.GetFileURL(ctx, item.Cover)
	}
	return list, total, nil
}

// CreateInformation 创建资讯信息
func (srv *InformationService) CreateInformation(ctx kratosx.Context, req *entity.Information) (uint32, error) {
	id, err := srv.repo.CreateInformation(ctx, req)
	if err != nil {
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateInformation 更新资讯信息
func (srv *InformationService) UpdateInformation(ctx kratosx.Context, ent *entity.Information) error {
	if err := srv.repo.UpdateInformation(ctx, ent); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteInformation 删除资讯信息
func (srv *InformationService) DeleteInformation(ctx kratosx.Context, id uint32) error {
	if err := srv.repo.DeleteInformation(ctx, id); err != nil {
		return errors.DeleteError(err.Error())
	}
	return nil
}
