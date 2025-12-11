package conf

import (
	"context"
	"os"

	"github.com/RanFeng/ilog"
	"gopkg.in/yaml.v3"
)

var conf = make(map[string]interface{})

func MustGet[T any](key string) T {
	v, ok := conf[key]
	if !ok {
		panic("该配置key不存在，key:" + key)
	}
	t, ok := v.(T)
	if !ok {
		panic("该配置key类型错误，key:" + key)
	}
	return t
}

func Init() {
	confFile := "./conf/test.yml"
	if os.Getenv("RUN_ENV") == "prod" {
		confFile = "./conf/prod.yml"
	} else if os.Getenv("RUN_ENV") == "dev" {
		confFile = "./conf/dev.yml"
	}

	b, err := os.ReadFile(confFile)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(b, &conf)
	if err != nil {
		panic(err)
	}

	ilog.EventInfo(context.Background(), "init_config", "config", conf)
}
