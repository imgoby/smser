package tools

import (
	"cn.sockstack/smser/entry"
	"cn.sockstack/smser/internal"
	"cn.sockstack/smser/internal/model"
	"github.com/sirupsen/logrus"
	"time"
)

const sourceFrom = "smser"

func RequestLogger(data interface{}) {
	log := entry.NewRequestLogEntry(time.Now(), sourceFrom, nil, data)
	model.GetMgoDB().C("request_log").Insert(log)
}

func WorkerLogger(data interface{}, response interface{}) {
	log := entry.NewWorkerLog(time.Now(), sourceFrom, response, data)
	model.GetMgoDB().C("worker_log").Insert(log)
}

func MessageLogger(data interface{}, response interface{}) {
	log := entry.NewMessageEntry(time.Now(), sourceFrom, response, data)
	model.GetMgoDB().C("message_send_log").Insert(log)
}

func Logger() *logrus.Logger {
	return internal.NewLog()
}
