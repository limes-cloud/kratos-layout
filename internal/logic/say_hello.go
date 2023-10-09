package logic

import (
	"github.com/limes-cloud/kratos"
	v1 "github.com/limes-cloud/kratos-layout/api/v1"
)

// SayHello implements helloworld.GreeterServer.
func (l *Logic) SayHello(ctx kratos.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	if l.conf.OpenWeb {
		ctx.Logger().Warn("close web")
	}

	return &v1.HelloReply{
		Message: in.Name,
	}, nil
}
