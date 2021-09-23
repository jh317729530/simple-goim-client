package conf

import (
	"flag"

	"github.com/BurntSushi/toml"
)

var (
	configPath string
	Conf       *Config
)

func init() {
	flag.StringVar(&configPath, "conf", "/Users/simontart/code/simple-goim-client/etc/config.toml", "default config path")
}

type Config struct {
	HttpServer *HttpServer
}

type HttpServer struct {
	Network string
	Addr    string
}

func Init() (err error) {
	Conf = Default()
	_, err = toml.DecodeFile(configPath, &Conf)
	if err != nil {
		panic(err)
	}

	return
}

func Default() *Config {
	return &Config{
		HttpServer: &HttpServer{
			Network: "tcp",
			Addr:    "8080",
		},
	}
}
