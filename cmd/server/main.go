package main

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/config"
	userpb "github.com/limes-cloud/user-center/api/user/v1"
	_ "go.uber.org/automaxprocs"

	internalconf "github.com/go-kratos/kratos-layout/internal/conf"
	"github.com/go-kratos/kratos-layout/internal/service"
)

func main() {
	app := kratosx.New(
		kratosx.Config(file.NewSource("internal/conf/config.yaml")),
		kratosx.RegistrarServer(RegisterServer),
	)

	kCtx := kratosx.MustContext(context.Background())
	conn, err := kCtx.GrpcConn("UserCenter")
	if err != nil {
		panic(err)
	}
	userClient := userpb.NewServiceClient(conn)
	user, err := userClient.GetSimpleUser(kCtx, &userpb.GetSimpleUserRequest{Id: 1})

	fmt.Println(user, err)
	if err := app.Run(); err != nil {
		log.Fatal(err.Error())
	}
}

func RegisterServer(c config.Config, hs *http.Server, gs *grpc.Server) {
	conf := &internalconf.Config{}
	if err := c.Value("file").Scan(conf); err != nil {
		panic("file config format error:" + err.Error())
	}
	c.Watch("file", func(value config.Value) {
		if err := value.Scan(conf); err != nil {
			log.Error("file 配置变更失败")
		} else {
			log.Error("file 配置变更成功")
		}
	})

	service.New(conf, hs, gs)
}
