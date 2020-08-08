package internal

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"time"
)

func GetMgoDB() *mgo.Database {
	url := fmt.Sprintf(
		"mongodb://%s:%s@%s:%s/?authSource=admin",
		Cfg.MongodbUsername,
		Cfg.MongodbPassword,
		Cfg.MongodbHost,
		Cfg.MongodbPort,
	)
	session, err := mgo.DialWithTimeout(url, time.Second*5)
	if err != nil {
		panic(err)
	}

	return session.DB(Cfg.MongodbName)
}
