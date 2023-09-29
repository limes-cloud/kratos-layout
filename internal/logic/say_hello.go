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
	ctx.Logger().Infof("hello  world")
	//tstr, err := ctx.JWT().NewToken(map[string]string{
	//	"user": "fangweiye",
	//})

	//m := map[string]string{}
	//err = ctx.JWT().ParseToken(ctx, &m)
	//fmt.Println(m, err)
	//user :=  model.User{}
	//user.Create()
	ctx.Logger().Warn("test")
	return &v1.HelloReply{
		Message: "11",
	}, nil
}
