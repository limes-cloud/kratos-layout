package core

import (
	kconfig "github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/limes-cloud/configure/api/configure/client"
	"github.com/limes-cloud/kratosx/config"
	"os"
)

var conf = &Conf{}

type Conf struct {
	AiBot struct {
		Secret  string
		Appid   string
		Setting struct {
			Name    string
			Logo    string
			Desc    string
			Guiding []struct {
				Text string
				Type string
			}
		}
	}
}

func configSource() kconfig.Source {
	host := os.Getenv("CONF_HOST")
	token := os.Getenv("CONF_TOKEN")
	name := os.Getenv("APP_NAME")
	if host != "" && token != "" && name != "" {
		return client.New(host, name, token)
	}
	return file.NewSource("conf/")
}

// configScanWatch 初始化
func configScanWatch(watch config.Watcher) {
	watch("business", func(value config.Value) {
		if err := value.Scan(&conf); err != nil {
			panic(err)
		}
	})
}
