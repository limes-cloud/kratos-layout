package service

import (
	"context"
	"github.com/limes-cloud/kratos"
	"github.com/limes-cloud/kratos-layout/internal/logic"

	v1 "github.com/limes-cloud/kratos-layout/api/v1"
	"github.com/limes-cloud/kratos-layout/config"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedGreeterServer
	logic *logic.Logic
}

func NewGreeterService(conf *config.Config) *GreeterService {
	return &GreeterService{
		logic: logic.NewLogic(conf),
	}
}

// SayHello implements helloworld.GreeterServer.
func (s *GreeterService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	return s.logic.SayHello(kratos.MustContext(ctx), in)
}
