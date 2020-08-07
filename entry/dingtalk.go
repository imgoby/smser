package entry

import "gopkg.in/mgo.v2/bson"

//DingTalkEntry DingTalk 实体
type DingTalkEntry struct {
	ID bson.ObjectId `bson:"_id,omitempty" json:"id"`
	AccessToken string `bson:"access_token" json:"access_token"`
	Secret string `bson:"secret" json:"secret"`
}

func (d DingTalkEntry) TableName() string {
	return "ding_talk"
}

//NewDingTalkEntry DingTalk 结构体
func NewDingTalkEntry(accessToken string, secret string) *DingTalkEntry {
	return &DingTalkEntry{AccessToken: accessToken, Secret: secret}
}
