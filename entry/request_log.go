package entry

import "time"

type RequestLogEntry struct {
	Time time.Time `bson:"time" json:"time"`
	SourceFrom string `bson:"source_from" json:"source_from"`
	Response interface{} `bson:"response,omitempty" json:"response"`
	Message interface{} `bson:"message,omitempty" json:"message"`
}

func NewRequestLogEntry(time time.Time, sourceFrom string, response interface{}, message interface{}) *RequestLogEntry {
	return &RequestLogEntry{SourceFrom: sourceFrom, Time: time, Response: response, Message: message}
}
