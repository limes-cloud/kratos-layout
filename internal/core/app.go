package core

import (
	"github.com/limes-cloud/kratosx"
)

// InitApp 初始化应用，默认注入配置源和热更新监听
func InitApp(opts ...kratosx.Option) *kratosx.App {
	defOpts := []kratosx.Option{
		kratosx.WithConfigSource(configSource()),
		kratosx.WithConfigWatch(configScanWatch),
	}
	return kratosx.New(append(defOpts, opts...)...)
}
