package entry

import "time"

type MessageLogEntry struct {
	Time time.Time `bson:"time" json:"time"`
	SourceFrom string `bson:"source_from" json:"source_from"`
	Response interface{} `bson:"response,omitempty" json:"response"`
	Message interface{} `bson:"message,omitempty" json:"message"`
}

func NewMessageLogEntry(time time.Time, sourceFrom string, response interface{}, message interface{}) *MessageLogEntry {
	return &MessageLogEntry{Time: time, SourceFrom: sourceFrom, Response: response, Message: message}
}
