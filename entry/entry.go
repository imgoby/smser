package entry

import "time"

type entry interface {
	TableName() string
}

type message interface {
	Encode() (string, error)
	Decode([]byte) (error)
}

type LogCommon struct {
	Time time.Time `bson:"time" json:"time"`
	SourceFrom string `bson:"source_from" json:"source_from"`
}

func NewLogCommon(time time.Time, sourceFrom string) *LogCommon {
	return &LogCommon{Time: time, SourceFrom: sourceFrom}
}
