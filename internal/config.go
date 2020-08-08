package internal

import (
	"gopkg.in/ini.v1"
)

type Config struct {
	Mode string `ini:"app_mode"`
	AppName string `ini:"app_name"`
	WorkerNum int `ini:"worker_num"`
	Http
	DingTalk
	MongoDB
	Nsq
}

type DingTalk struct {
	AccessToken string `ini:"access_token"`
	Secret string `ini:"secret"`
}

type Http struct {
	Host string `ini:"http_host"`
	Port string `ini:"http_port"`
}

type MongoDB struct {
	MongodbHost string `ini:"mongodb_host"`
	MongodbPort string `ini:"mongodb_port"`
	MongodbName string `ini:"mongodb_db"`
	MongodbLogCollection string `ini:"mongodb_log_collection"`
	MongodbUsername string `ini:"mongodb_username"`
	MongodbPassword string `ini:"mongodb_password"`
}

type Nsq struct {
	NsqHost string `ini:"nsq_host"`
	NsqPort string `ini:"nsq_port"`
	NsqConsumerHost string `ini:"nsq_consumer_host"`
	NsqConsumerPort string `ini:"nsq_consumer_port"`
	NsqMessageTopic string `ini:"nsq_message_topic"`
}

func NewConfig(path string) (config *Config) {
	c := new(Config)
	ini.MapTo(c, path)

	return c
}