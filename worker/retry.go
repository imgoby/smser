package worker

import (
	"cn.sockstack/smser/entry"
	"cn.sockstack/smser/internal"
	"fmt"
	"github.com/robfig/cron/v3"
	"gopkg.in/mgo.v2/bson"
	"time"
)

func Retry()  {
	c := cron.New(cron.WithSeconds())
	_, err := c.AddFunc("*/5 * * * * *", check)
	if err != nil {
		fmt.Println(err)
	}
	c.Start()

	for true {
		time.Sleep(time.Second)
	}
}

func check()  {
	var message []entry.QueueEntry
	queueEntry := entry.QueueEntry{}
	err := internal.GetMgoDB().C(queueEntry.TableName()).
		Find(bson.M{"status": bson.M{"$in": []int{entry.PrepareStatus, entry.RetryStatus, entry.SendSuccessStatus}}, "retry_at": bson.M{"$lt": time.Now().Unix()}}).
		Sort("updated_at").Limit(100).All(&message)
	if err != nil {
		return
	}

	for _, q := range message{
		Send(q)
	}
}
