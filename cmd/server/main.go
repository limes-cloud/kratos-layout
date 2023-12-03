package main

import (
	"github.com/limes-cloud/kratos-layout/internal/handler"
	"github.com/limes-cloud/kratos/contrib/config/configure"
	"os"

	"github.com/limes-cloud/kratos"
	v1 "github.com/limes-cloud/kratos-layout/api/v1"
	srcConf "github.com/limes-cloud/kratos-layout/config"
	"github.com/limes-cloud/kratos/config"
	"github.com/limes-cloud/kratos/config/file"
	"github.com/limes-cloud/kratos/log"
	"github.com/limes-cloud/kratos/middleware/tracing"
	"github.com/limes-cloud/kratos/transport/grpc"
	"github.com/limes-cloud/kratos/transport/http"
	_ "go.uber.org/automaxprocs"
)

var (
	Name    = os.Getenv("AppName")
	Version = os.Getenv("Version")
	id, _   = os.Hostname()
)

func main() {
	app := kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		// 从配置中心加载配置
		kratos.Config(configure.NewFromEnv()),
		// 从文件加载配置
		kratos.Config(file.NewSource("config/config.yaml")),
		kratos.RegistrarServer(RegisterServer),
		kratos.LoggerWith(kratos.LogField{
			"id":      id,
			"name":    Name,
			"version": Version,
			"trace":   tracing.TraceID(),
			"span":    tracing.SpanID(),
		}),
	)

	if err := app.Run(); err != nil {
		log.Errorf("run service fail: %v", err)
	}
}

func RegisterServer(hs *http.Server, gs *grpc.Server, c config.Config) {
	conf := &srcConf.Config{}
	if err := c.ScanKey("business", conf); err != nil {
		panic("business config format error:" + err.Error())
	}

	srv := handler.New(conf)
	v1.RegisterServiceHTTPServer(hs, srv)
	v1.RegisterServiceServer(gs, srv)
}
