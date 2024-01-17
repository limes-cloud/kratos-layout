package service

import (
	v1 "github.com/go-kratos/kratos-layout/api/v1"
	"github.com/go-kratos/kratos-layout/config"
	"github.com/go-kratos/kratos-layout/internal/biz"
	"github.com/go-kratos/kratos-layout/internal/data"
)

// Service is a greeter service.
type Service struct {
	v1.UnimplementedServiceServer
	user *biz.UserUseCase
}

func New(conf *config.Config) *Service {
	return &Service{
		user: biz.NewUserUseCase(conf, data.NewUserRepo()),
	}
}
