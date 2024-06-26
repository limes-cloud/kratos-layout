package main

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/config"
	_ "go.uber.org/automaxprocs"

	"layout/internal/conf"
	"layout/internal/service"
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
	cfg := &conf.Config{}
	if err := c.Value("business").Scan(cfg); err != nil {
		panic("business config format error:" + err.Error())
	}
	c.ScanWatch("business", func(value config.Value) {
		if err := value.Scan(cfg); err != nil {
			log.Error("business 配置变更失败")
		} else {
			log.Error("file 配置变更成功")
		}
	})

	service.New(cfg, hs, gs)
}
