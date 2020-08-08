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
	Type int `bson:"type,omitempty"`
	Payload string `bson:"payload,omitempty"`
	Status int `bson:"status,omitempty"`
	RetryNum int `bson:"retry_num,omitempty"`
	RetryAt int64 `bson:"retry_at,omitempty"`
	CreatedAt int64 `bson:"created_at,omitempty"`
	UpdatedAt int64 `bson:"updated_at,omitempty"`
	DeletedAt int64 `bson:"deleted_at,omitempty"`
}

func NewQueueEntry() *QueueEntry {
	return &QueueEntry{}
}

func (q QueueEntry) TableName() string {
	return "message"
}

