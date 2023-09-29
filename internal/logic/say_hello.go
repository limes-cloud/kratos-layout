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

	sender, err := ctx.Email()
	if err != nil {
		return nil, err
	}
	if err = sender.Send("1280291001@qq.com", map[string]interface{}{
		"name": "方伟业",
	}); err != nil {
		return nil, err
	}

	ctx.Logger().Warn("hello world")
	return &v1.HelloReply{
		Message: in.Name,
	}, nil
}
