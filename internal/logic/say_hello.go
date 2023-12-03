package logic

import (
	"github.com/limes-cloud/kratos"
	v1 "github.com/limes-cloud/kratos-layout/api/v1"
	"github.com/limes-cloud/kratos-layout/config"
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
func (l *Logic) SayHello(ctx kratos.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	ctx.Logger().Info("info", "say hello request")
	return &v1.HelloReply{
		Message: l.conf.SayText + ":" + in.Name,
	}, nil
}
