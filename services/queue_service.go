package services

import (
	"cn.sockstack/smser/entry"
	"cn.sockstack/smser/internal"
	"cn.sockstack/smser/internal/model"
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
			queueEntry.Status = entry.RetryStatus
			model.GetMgoDB().C(queueEntry.TableName()).UpdateId(queueEntry.ID, queueEntry)
			return
		}

		queueEntry.Status = entry.SendSuccessStatus
		model.GetMgoDB().C(queueEntry.TableName()).UpdateId(queueEntry.ID, queueEntry)
		// Gracefully stop the producer.
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

func NewQueueService() *QueueService {
	return &QueueService{}
}
