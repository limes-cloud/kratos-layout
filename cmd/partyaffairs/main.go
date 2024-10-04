package main

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/configure/api/configure/client"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/config"
	_ "go.uber.org/automaxprocs"

	"partyaffairs/internal/app"
	"partyaffairs/internal/conf"
)

func main() {
	// 启动管理端
	server := kratosx.New(
		kratosx.Config(client.NewFromEnv()),
		kratosx.RegistrarServer(RegisterServer),
	)

	if err := server.Run(); err != nil {
		log.Fatal(err.Error())
	}
}

// RegisterServer 注册服务端
func RegisterServer(c config.Config, hs *http.Server, _ *grpc.Server) {
	cfg := &conf.Config{}
	app.Register(cfg, hs)
}
