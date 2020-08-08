package entry

import (
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
)

//DingTalkEntry DingTalk 实体
type DingTalkEntry struct {
	ID bson.ObjectId `form:"id" bson:"_id,omitempty" json:"id,omitempty"`
	AccessToken string `form:"access_token" bson:"access_token" json:"access_token" binding:"required"`
	Secret string `form:"secret" bson:"secret" json:"secret" binding:"required"`
}

func (d DingTalkEntry) TableName() string {
	return "ding_talk"
}

//NewDingTalkEntry DingTalk 结构体
func NewDingTalkEntry(accessToken string, secret string) *DingTalkEntry {
	return &DingTalkEntry{AccessToken: accessToken, Secret: secret}
}

type DingTalkTextMessageEntry struct {
	Message string `form:"message" binding:"required" json:"message"`
	AtMobiles []string `form:"at_mobiles" json:"at_mobiles"`
}

func (d *DingTalkTextMessageEntry) Decode(bytes []byte) error {
	return json.Unmarshal(bytes, d)
}

func (d *DingTalkTextMessageEntry) Encode() (string, error) {
	bytes, err := json.Marshal(d)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func NewDingTalkTextMessageEntry() *DingTalkTextMessageEntry {
	return &DingTalkTextMessageEntry{}
}
