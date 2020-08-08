package worker

import (
	"cn.sockstack/smser/entry"
	"cn.sockstack/smser/internal"
	"cn.sockstack/smser/internal/model"
	"cn.sockstack/smser/services"
	"cn.sockstack/smser/tools"
	"context"
	"fmt"
	"github.com/nsqio/go-nsq"
	"sync"
	"time"
)

var Worker = newWorker()

type worker struct {
	q chan bool
	m sync.Mutex
	w sync.WaitGroup
	Num int
}

func newWorker() *worker {
	return &worker{
		Num: internal.Cfg.WorkerNum,
		q: make(chan bool, 1),
		m: sync.Mutex{},
		w: sync.WaitGroup{},
	}
}

func (this *worker) Start ()  {
	tools.WorkerLogger().Info("worker starting")
	ctx, cancelFunc := context.WithCancel(context.Background())
	go func() {
		defer func() {
			tools.WorkerLogger().Info("worker exited")
			cancelFunc()
		}()
		for i := 1; i <= this.Num; i++ {
			fmt.Println(i)
			go this.consumer(ctx)
		}
		<-this.q
	}()
	tools.WorkerLogger().Info("worker started")
}

func (this *worker) Restart ()  {
	tools.WorkerLogger().Info("worker restarting")
	this.q <- true
	this.w.Wait()
	tools.WorkerLogger().Info("worker restarted")
	this.Start()
}

func (this *worker) consumer (ctx context.Context) {
	this.w.Add(1)

	var flag bool
	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer(internal.Cfg.NsqMessageTopic, "message_channel", config)
	// Gracefully stop the consumer.
	defer func() {
		fmt.Println("进程推出")
		this.w.Done()
		consumer.Stop()
	}()
	if err != nil {
		tools.WorkerLogger().Error(err)
		return
	}

	// Set the Handler for messages received by this Consumer. Can be called multiple times.
	// See also AddConcurrentHandlers.
	consumer.AddHandler(newHandle())

	// Use nsqlookupd to discover nsqd instances.
	// See also ConnectToNSQD, ConnectToNSQDs, ConnectToNSQLookupds.
	err = consumer.ConnectToNSQLookupd(fmt.Sprintf("%s:%s", internal.Cfg.NsqConsumerHost, internal.Cfg.NsqConsumerPort))
	if err != nil {
		tools.WorkerLogger().Error(err)
	}
	for true {
		select {
		case <-ctx.Done():
			flag = true
			break
		default:
			time.Sleep(time.Millisecond * 50)
		}

		if flag {
			break
		}
	}

}
type handle struct {}

func newHandle() *handle {
	return &handle{}
}

// HandleMessage implements the Handler interface.
func (h *handle) HandleMessage(m *nsq.Message) error {
	if len(m.Body) == 0 {
		// Returning nil will automatically send a FIN command to NSQ to mark the message as processed.
		return nil
	}

	err := processMessage(m.Body)

	// Returning a non-nil error will automatically send a REQ command to NSQ to re-queue the message.
	return err
}

func processMessage(body []byte) error {
	queueService := services.NewQueueService()
	queueEntry, err := queueService.Decode(body)
	if err != nil {
		queueEntry.Status = entry.RetryStatus
		model.GetMgoDB().C(queueEntry.TableName()).UpdateId(queueEntry.ID, queueEntry)
		return err
	}

	switch queueEntry.Type {
	case entry.DingTalkTextMessage:
		err = SendDingTalkTextMessage(queueEntry)
		break
	default:
	}

	return err
}