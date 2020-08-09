package tools

import (
	"cn.sockstack/smser/entry"
	"cn.sockstack/smser/internal"
	"github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"
	"time"
)

var sourceFrom = internal.Cfg.AppName

func RequestLogger(data interface{}) {
	log := entry.NewRequestLogEntry(time.Now(), sourceFrom, nil, data)
	internal.GetMgoDB().C("request_log").Insert(log)
}

func WorkerLogger(data interface{}, response interface{}) {
	log := entry.NewWorkerLog(time.Now(), sourceFrom, response, data)
	internal.GetMgoDB().C("worker_log").Insert(log)
}

func MessageLogger(data interface{}, response interface{}) {
	log := entry.NewMessageLogEntry(time.Now(), sourceFrom, response, data)
	internal.GetMgoDB().C("message_send_log").Insert(log)
}

func Logger() *logrus.Logger {
	return internal.NewLog()
}

func generateRetryAt(num int) int64 {
	return time.Now().Unix() + int64(num * 60)
}

func RetryRecord(queueEntry entry.QueueEntry)  {
	if queueEntry.RetryNum < 5 {
		queueEntry.Status = entry.RetryStatus
		queueEntry.RetryNum++
		queueEntry.RetryAt = generateRetryAt(queueEntry.RetryNum)
	} else {
		queueEntry.Status = entry.SendFailStatus
	}
	internal.GetMgoDB().C(queueEntry.TableName()).UpdateId(queueEntry.ID, queueEntry)
}

func QueueAckRecordByMessageID(id bson.ObjectId)  {
	queueEntry := entry.NewQueueEntry()
	err := internal.GetMgoDB().C(queueEntry.TableName()).FindId(id).One(&queueEntry)
	if err != nil {
		Logger().Info(err)
		return
	}

	queueEntry.Status = entry.AckStatus
	err = internal.GetMgoDB().C(queueEntry.TableName()).UpdateId(queueEntry.ID, queueEntry)
	if err != nil {
		Logger().Info(err)
		return
	}
}

func QueueSuccessRecord(queueEntry entry.QueueEntry)  {
	queueEntry.Status = entry.SendSuccessStatus
	internal.GetMgoDB().C(queueEntry.TableName()).UpdateId(queueEntry.ID, queueEntry)
}
