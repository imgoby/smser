package entry

type entry interface {
	TableName() string
}

type message interface {
	Encode() (string, error)
	Decode([]byte) (error)
}
