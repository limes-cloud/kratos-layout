package logic

import (
	v1 "github.com/go-kratos/kratos-layout/api/v1"
	"github.com/go-kratos/kratos-layout/config"
	"github.com/limes-cloud/kratosx"
)

type Logic struct {
	conf *config.Config
}

func NewLogic(conf *config.Config) *Logic {
	return &Logic{
		conf: conf,
	}
}

// SayHello implements SayHello
func (l *Logic) SayHello(ctx kratosx.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	ctx.Logger().Infow("info", "say hello request")
	return &v1.HelloReply{
		Message: l.conf.SayText + ":" + in.Name,
	}, nil
}
