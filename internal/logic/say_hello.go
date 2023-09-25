package logic

import (
	"fmt"
	"github.com/limes-cloud/kratos"
	v1 "github.com/limes-cloud/kratos-layout/api/v1"
)

// SayHello implements helloworld.GreeterServer.
func (l *Logic) SayHello(ctx kratos.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	if l.conf.OpenWeb {
		ctx.Logger().Warn("close web")
	}

	tstr, err := ctx.JWT().NewToken(map[string]string{
		"user": "fangweiye",
	})
	if err != nil {
		return nil, v1.ErrorUserNotFound("用户token生成失败")
	}

	m := map[string]string{}
	err = ctx.JWT().ParseToken(ctx, &m)
	fmt.Println(m, err)
	//user :=  model.User{}
	//user.Create()

	return &v1.HelloReply{
		Message: tstr,
	}, nil
}
