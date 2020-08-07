package model

import (
	"cn.sockstack/smser/internal"
	"fmt"
	"gopkg.in/mgo.v2"
	"time"
)

func GetMgoDB() *mgo.Database {
	url := fmt.Sprintf(
		"mongodb://%s:%s@%s:%s/?authSource=admin",
		internal.Cfg.MongodbUsername,
		internal.Cfg.MongodbPassword,
		internal.Cfg.MongodbHost,
		internal.Cfg.MongodbPort,
	)
	session, err := mgo.DialWithTimeout(url, time.Second*5)
	if err != nil {
		panic(err)
	}

	return session.DB(internal.Cfg.MongodbName)
}
