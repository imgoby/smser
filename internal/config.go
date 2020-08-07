package internal

import (
	"gopkg.in/ini.v1"
)

type DingTalk struct {
	AccessToken string `ini:"access_token"`
	Secret string `ini:"secret"`
}

type Http struct {
	Host string `ini:"http_host"`
	Port string `ini:"http_port"`
}

type Config struct {
	Mode string `ini:"app_mode"`
	AppName string `ini:"app_name"`
	Http
	DingTalk
	MongoDB
}

type MongoDB struct {
	MongodbHost string `ini:"mongodb_host"`
	MongodbPort string `ini:"mongodb_port"`
	MongodbName string `ini:"mongodb_db"`
	MongodbLogCollection string `ini:"mongodb_log_collection"`
	MongodbUsername string `ini:"mongodb_username"`
	MongodbPassword string `ini:"mongodb_password"`
}

func NewConfig(path string) (config *Config) {
	c := new(Config)
	ini.MapTo(c, path)

	return c
}