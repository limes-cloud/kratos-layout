package service

import (
	v1 "partyaffairs/api/partyaffairs/v1"
	"partyaffairs/internal/biz"
	"partyaffairs/internal/conf"
	"partyaffairs/internal/data"
)

// Service is a admin service.
type Service struct {
	v1.UnimplementedPartyAffairsServer
	conf     *conf.Config
	notice   *biz.NoticeUseCase
	banner   *biz.BannerUseCase
	news     *biz.NewsUseCase
	resource *biz.ResourceUseCase
	task     *biz.TaskUseCase
	video    *biz.VideoUseCase
}

func New(conf *conf.Config) *Service {
	return &Service{
		conf:     conf,
		notice:   biz.NewNoticeUseCase(conf, data.NewNoticeRepo()),
		banner:   biz.NewBannerUseCase(conf, data.NewBannerRepo()),
		news:     biz.NewNewsUseCase(conf, data.NewNewsRepo()),
		resource: biz.NewResourceUseCase(conf, data.NewResourceRepo()),
		task:     biz.NewTaskUseCase(conf, data.NewTaskRepo()),
		video:    biz.NewVideoUseCase(conf, data.NewVideoRepo()),
	}
}
