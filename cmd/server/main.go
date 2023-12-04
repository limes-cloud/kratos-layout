package main

import (
	v1 "github.com/go-kratos/kratos-layout/api/v1"
	systemConfig "github.com/go-kratos/kratos-layout/config"
	"github.com/go-kratos/kratos-layout/internal/service"
	kratosConfig "github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/config"
	_ "go.uber.org/automaxprocs"
)

func main() {
	app := kratosx.New(
		kratosx.RegistrarServer(RegisterServer),
	)

	if err := app.Run(); err != nil {
		log.Fatal(err.Error())
	}
}

func RegisterServer(c config.Config, hs *http.Server, gs *grpc.Server) {
	conf := &systemConfig.Config{}
	if err := c.Value("business").Scan(conf); err != nil {
		panic("business config format error:" + err.Error())
	}
	c.Watch("business", func(value kratosConfig.Value) {
		if err := value.Scan(conf); err != nil {
			log.Error("business 配置变更失败")
		}
	})

	srv := service.New(conf)
	v1.RegisterServiceHTTPServer(hs, srv)
	v1.RegisterServiceServer(gs, srv)
}
