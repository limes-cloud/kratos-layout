package main

import (
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/config"
	_ "go.uber.org/automaxprocs"

	internalconf "github.com/go-kratos/kratos-layout/internal/conf"
	"github.com/go-kratos/kratos-layout/internal/service"
)

func main() {
	app := kratosx.New(
		kratosx.Config(file.NewSource("internal/conf/config.yaml")),
		kratosx.RegistrarServer(RegisterServer),
	)

	if err := app.Run(); err != nil {
		log.Fatal(err.Error())
	}
}

func RegisterServer(c config.Config, hs *http.Server, gs *grpc.Server) {
	conf := &internalconf.Config{}
	if err := c.Value("business").Scan(conf); err != nil {
		panic("business config format error:" + err.Error())
	}
	c.Watch("business", func(value config.Value) {
		if err := value.Scan(conf); err != nil {
			log.Error("business 配置变更失败")
		}
	})

	service.New(conf, hs, gs)
}
