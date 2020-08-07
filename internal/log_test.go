package internal

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"testing"
	"time"
)

func TestNewLog(t *testing.T) {
	//config := NewConfig("../config/config.ini")
	session, err := mgo.DialWithTimeout("mongodb://root:root@132.232.59.192:27017/?authSource=admin", time.Second*5)
	if err != nil {
		t.Fatal(err)
	}
	err = session.DB("smser").C("log").Insert(bson.M{"test": "ok"})
	if err != nil {
		t.Fatal(err)
	}
}
