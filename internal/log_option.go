package internal

type logOption struct {
	Db string
	Collection string
	Username string
	Password string
	Host string
	Port string
	Extra map[string]string
}

type LogOptionHandle func(opt *logOption)

func WithDb(db string) LogOptionHandle {
	return func(opt *logOption) {
		opt.Db = db
	}
}

func WithCollection(c string) LogOptionHandle {
	return func(opt *logOption) {
		opt.Collection = c
	}
}

func WithExtra(extra map[string]string) LogOptionHandle {
	return func(opt *logOption) {
		opt.Extra = extra
	}
}
