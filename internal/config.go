package internal

import (
	"gopkg.in/ini.v1"
)

type Config struct {
	Mode string `ini:"app_mode"`
	Host string `ini:"http_host"`
	Port string `ini:"http_port"`
}

func NewConfig(path string) (config *Config) {
	c := new(Config)
	ini.MapTo(c, path)

	return c
}