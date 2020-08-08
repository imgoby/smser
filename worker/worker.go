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
	rwm sync.RWMutex
	w sync.WaitGroup
	num int
	rst bool
	rch chan bool
	ctx *context.Context
}

func newWorker() *worker {
	return &worker{
		num: internal.Cfg.WorkerNum,
		q: make(chan bool, 1),
		m: sync.Mutex{},
		w: sync.WaitGroup{},
		rwm: sync.RWMutex{},
		rch: make(chan bool, 1),
	}
}

func (this *worker) SetNum (num int) {
	this.rwm.Lock()
	defer this.rwm.Unlock()
	this.num = num

	this.rch <- true
}

func (this *worker) GetNum() int {
	this.rwm.RLock()
	defer this.rwm.RUnlock()
	return this.num
}

func (this *worker) Run ()  {
	this.start()
	go func() {
		for true {
			<-this.rch
			this.restart()
		}
	}()
}

func (this *worker) start ()  {
	ctx, cancelFunc := context.WithCancel(context.Background())
	go func() {
		for i := 1; i <= this.num; i++ {
			go this.consumer(ctx)
		}
		for true {
			select {
			case <-this.q:
				cancelFunc()
				break
			case <-ctx.Done():
				return
			default:
				time.Sleep(time.Millisecond * 50)
			}
		}
	}()
}

func (this *worker) restart ()  {
	tools.WorkerLogger("worker restarting", nil)
	this.m.Lock()
	defer this.m.Unlock()

	this.q <- true
	this.w.Wait()
	this.start()
	tools.WorkerLogger("worker restarted", nil)
}

func (this *worker) consumer (ctx context.Context) {
	this.w.Add(1)

	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer(internal.Cfg.NsqMessageTopic, "message_channel", config)
	// Gracefully stop the consumer.
	defer func() {
		fmt.Println("进程推出")
		this.w.Done()
		consumer.Stop()
	}()
	if err != nil {
		tools.WorkerLogger(err, nil)
		return
	}

	// Set the Handler for messages received by this Consumer. Can be called multiple times.
	// See also AddConcurrentHandlers.
	consumer.AddHandler(newHandle())

	// Use nsqlookupd to discover nsqd instances.
	// See also ConnectToNSQD, ConnectToNSQDs, ConnectToNSQLookupds.
	err = consumer.ConnectToNSQLookupd(fmt.Sprintf("%s:%s", internal.Cfg.NsqConsumerHost, internal.Cfg.NsqConsumerPort))
	if err != nil {
		tools.WorkerLogger(err, nil)
	}
	for true {
		select {
		case <-ctx.Done():
			return
		default:
			time.Sleep(time.Millisecond * 50)
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