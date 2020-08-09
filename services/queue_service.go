package services

import (
	"cn.sockstack/smser/entry"
	"cn.sockstack/smser/internal"
	"cn.sockstack/smser/tools"
	"encoding/json"
	"fmt"
	"github.com/nsqio/go-nsq"
)

type QueueService struct {
}

func (this *QueueService) Push(queueEntry entry.QueueEntry) {
	// Instantiate a producer.
	go func() {
		config := nsq.NewConfig()
		p, err := nsq.NewProducer(fmt.Sprintf("%s:%s", internal.Cfg.NsqHost, internal.Cfg.NsqPort), config)
		if err != nil {
			tools.Logger().Error(err)
			return
		}

		encode, err := this.Encode(queueEntry)
		if err != nil {
			tools.Logger().Error(err)
			return
		}

		messageBody := encode
		topicName := internal.Cfg.NsqMessageTopic

		// Synchronously publish a single message to the specified topic.
		// Messages can also be sent asynchronously and/or in batches.
		err = p.Publish(topicName, messageBody)
		if err != nil {
			tools.RetryRecord(queueEntry)
			return
		}

		// Gracefully stop the producer.
		tools.QueueSuccessRecord(queueEntry)
		p.Stop()
	}()
}

func (this *QueueService) Encode(queueEntry entry.QueueEntry) ([]byte, error) {
	return json.Marshal(queueEntry)
}

func (this *QueueService) Decode (data []byte) (entry.QueueEntry, error) {
	queueEntry := entry.NewQueueEntry()

	err := json.Unmarshal(data, queueEntry)
	return *queueEntry, err
}

func (this *QueueService) MessageList (queueEntry entry.QueueEntry) (messages []entry.QueueEntry, err error) {
	if queueEntry.Page > 0 {
		queueEntry.Page--
	}
	if queueEntry.Size <= 0 {
		queueEntry.Size = 10
	}

	err = internal.GetMgoDB().C(queueEntry.TableName()).Find(nil).Sort("-created_at").Skip(queueEntry.Size * queueEntry.Page).Limit(queueEntry.Size).All(&messages)
	return messages, err
}

func NewQueueService() *QueueService {
	return &QueueService{}
}
