package internal

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/weekface/mgorus"
)

var (
	log = logrus.New()
)

func NewLog(opts ...LogOptionHandle) *logrus.Entry {

	option := logOption{
		Db:         Cfg.MongodbName,
		Collection: Cfg.MongodbLogCollection,
		Username:   Cfg.MongodbUsername,
		Password:   Cfg.MongodbPassword,
		Host:       Cfg.MongodbHost,
		Port:       Cfg.MongodbPort,
	}

	for _, o := range opts{
		o(&option)
	}

	mgourl := fmt.Sprintf(
		"mongodb://%s:%s@%s:%s/?authSource=admin",
		option.Username,
		option.Password,
		option.Host,
		option.Port,
	)
	hooker, err := mgorus.NewHooker(mgourl,
		option.Db,
		option.Collection,
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Hooks.Add(hooker)

	fields := logrus.Fields{
		"source_from": Cfg.AppName,
	}
	for key, value := range option.Extra {
		fields[key] = value
	}
	entry := log.WithFields(fields)

	return entry
}
