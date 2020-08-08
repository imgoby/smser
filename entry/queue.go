package entry

import "gopkg.in/mgo.v2/bson"

const (
	DingTalkTextMessage = 1

	PrepareStatus = 0
	SendSuccessStatus = 1
	SendFailStatus = 2
	RetryStatus = 3
	AckStatus = 4

	RetryNum = 0
)

type QueueEntry struct {
	ID bson.ObjectId `bson:"_id,omitempty"`
	Type int `bson:"type"`
	Payload string `bson:"payload"`
	Status int `bson:"status"`
	RetryNum int `bson:"retry_num"`
	RetryAt int64 `bson:"retry_at"`
	CreatedAt int64 `bson:"created_at"`
	UpdatedAt int64 `bson:"updated_at"`
	DeletedAt int64 `bson:"deleted_at,omitempty"`
}

func NewQueueEntry() *QueueEntry {
	return &QueueEntry{}
}

func (q QueueEntry) TableName() string {
	return "message"
}

