package logic

import "github.com/limes-cloud/kratos-layout/config"

type Logic struct {
	conf *config.Config
}

func NewLogic(conf *config.Config) *Logic {
	return &Logic{
		conf: conf,
	}
}
