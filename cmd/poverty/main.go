package main

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/config"
	_ "go.uber.org/automaxprocs"

	"poverty/internal/app"
	"poverty/internal/conf"
)

func main() {
	srv := kratosx.New(
		//kratosx.Config(client.NewFromEnv()),
		kratosx.RegistrarServer(RegisterServer),
	)

	if err := srv.Run(); err != nil {
		log.Fatal(err.Error())
	}
}

func RegisterServer(c config.Config, hs *http.Server, gs *grpc.Server) {
	cfg := &conf.Config{}
	c.ScanWatch("business", func(value config.Value) {
		if err := value.Scan(cfg); err != nil {
			log.Error("business 配置变更失败")
		} else {
			log.Error("file 配置变更成功")
		}
	})

	app.Register(cfg, hs, gs)
}
