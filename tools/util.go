package tools

import (
	"cn.sockstack/smser/internal"
	"github.com/sirupsen/logrus"
)

func RequestLogger() *logrus.Entry {
	return internal.NewLog(internal.WithCollection("request_log"))
}

func WorkerLogger() *logrus.Entry {
	return internal.NewLog(internal.WithCollection("worker_log"))
}

func MessageLogger(extra map[string]interface{}) *logrus.Entry {
	return internal.NewLog(internal.WithCollection("message_send_log"), internal.WithExtra(extra))
}

func Logger() *logrus.Entry {
	return internal.NewLog(internal.WithCollection("log"))
}
