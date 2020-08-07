package tools

import (
	"cn.sockstack/smser/internal"
	"github.com/sirupsen/logrus"
)

func RequestLogger() *logrus.Entry {
	return internal.NewLog(internal.WithCollection("request_log"))
}

