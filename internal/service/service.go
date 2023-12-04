package service

import (
	"context"
	"github.com/go-kratos/kratos-layout/internal/logic"
	"github.com/limes-cloud/kratosx"

	v1 "github.com/go-kratos/kratos-layout/api/v1"
	"github.com/go-kratos/kratos-layout/config"
)

// Service is a greeter service.
type Service struct {
	v1.UnimplementedServiceServer
	logic *logic.Logic
}

func New(conf *config.Config) *Service {
	return &Service{
		logic: logic.NewLogic(conf),
	}
}

func (s *Service) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	return s.logic.SayHello(kratosx.MustContext(ctx), in)
}
