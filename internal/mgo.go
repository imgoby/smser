package internal

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"sync"
	"time"
)


var (
	session *mgo.Session
	m sync.Mutex
)

var err error

func GetMgoDB() *mgo.Database {
	m.Lock()
	defer m.Unlock()
	if session != nil {
		return session.DB(Cfg.MongodbName)
	}
	url := fmt.Sprintf(
		"mongodb://%s:%s@%s:%s/?authSource=admin",
		Cfg.MongodbUsername,
		Cfg.MongodbPassword,
		Cfg.MongodbHost,
		Cfg.MongodbPort,
	)
	session, err = mgo.DialWithTimeout(url, time.Second*5)
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Eventual, true)
	session.SetPoolLimit(10)
	return session.DB(Cfg.MongodbName)
}
