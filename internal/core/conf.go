package core

import (
	"os"

	kconfig "github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/limes-cloud/kratosx/config"
)

var conf = &Conf{}

// Conf 业务配置，对应 config.yaml 中 business 键
type Conf struct {
}

func configSource() kconfig.Source {
	host := os.Getenv("CONF_HOST")
	token := os.Getenv("CONF_TOKEN")
	name := os.Getenv("APP_NAME")
	if host != "" && token != "" && name != "" {
		// 若需要接入远程配置中心，在此引入对应 client 包
		// return client.New(host, name, token)
		_ = host
		_ = token
		_ = name
	}
	return file.NewSource("conf/")
}

// configScanWatch 监听 business 配置热更新
func configScanWatch(watch config.Watcher) {
	watch("business", func(value config.Value) {
		if err := value.Scan(&conf); err != nil {
			panic(err)
		}
	})
}
