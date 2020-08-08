package worker

import (
	"cn.sockstack/smser/entry"
	"cn.sockstack/smser/internal"
	"cn.sockstack/smser/internal/model"
	"cn.sockstack/smser/services"
	"cn.sockstack/smser/tools"
	"fmt"
	"github.com/nsqio/go-nsq"
	"time"
)

func Consumer() {
	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer(internal.Cfg.NsqMessageTopic, "message_channel", config)
	// Gracefully stop the consumer.
	defer consumer.Stop()
	if err != nil {
		tools.WorkerLogger().Error(err)
		return
	}

	// Set the Handler for messages received by this Consumer. Can be called multiple times.
	// See also AddConcurrentHandlers.
	consumer.AddHandler(newWorker())

	// Use nsqlookupd to discover nsqd instances.
	// See also ConnectToNSQD, ConnectToNSQDs, ConnectToNSQLookupds.
	err = consumer.ConnectToNSQLookupd(fmt.Sprintf("%s:%s", internal.Cfg.NsqConsumerHost, internal.Cfg.NsqConsumerPort))
	if err != nil {
		tools.WorkerLogger().Error(err)
	}
	for true {
		time.Sleep(time.Millisecond * 50)
	}

}
type Worker struct {}

func newWorker() *Worker {
	return &Worker{}
}

// HandleMessage implements the Handler interface.
func (h *Worker) HandleMessage(m *nsq.Message) error {
	if len(m.Body) == 0 {
		// Returning nil will automatically send a FIN command to NSQ to mark the message as processed.
		return nil
	}

	err := processMessage(m.Body)

	// Returning a non-nil error will automatically send a REQ command to NSQ to re-queue the message.
	return err
}

func processMessage(body []byte) error {
	service := services.NewDingTalkService()
	queueService := services.NewQueueService()
	queueEntry, err := queueService.Decode(body)
	if err != nil {
		queueEntry.Status = entry.RetryStatus
		model.GetMgoDB().C(queueEntry.TableName()).UpdateId(queueEntry.ID, queueEntry)
		return err
	}

	messageEntry := entry.NewDingTalkTextMessageEntry()
	messageEntry.Decode([]byte(queueEntry.Payload))
	fmt.Println(string(body))
	fmt.Println(queueEntry)

	service.SetAccessTokenAndSecret(internal.Cfg.AccessToken, internal.Cfg.Secret).SetTextMessage(*messageEntry).Send()
	return nil
}