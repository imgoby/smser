package services

import (
	"cn.sockstack/smser/entry"
	"cn.sockstack/smser/internal/model"
	"cn.sockstack/smser/tools"
	"encoding/json"
	"github.com/sockstack/dtrobot"
	"github.com/sockstack/dtrobot/message"
	"gopkg.in/mgo.v2/bson"
	"time"
)

var (
	dingtalk = model.GetMgoDB()
)

type DingTalkService struct {
	AccessToken string
	Secret string
	message message.Message
}

func NewDingTalkService() *DingTalkService {
	return &DingTalkService{}
}

func (this *DingTalkService) Send ()  {
	robot, err := dtrobot.NewRobot(this.AccessToken, dtrobot.WithSecret(this.Secret))
	if err != nil {
		tools.MessageLogger(nil).Error(err)
	}

	response, err := robot.Send(this.message)

	if err != nil {
		tools.MessageLogger(nil).Error(err)
	}

	message, _ := json.Marshal(this.message)
	tools.MessageLogger(map[string]interface{}{"response": response}).Info(string(message))
}

//StoreAccessTokenAndSecret 持久化 DingTalk 配置
func (this *DingTalkService) StoreAccessTokenAndSecret (dingTalkEntry entry.DingTalkEntry) error {
	collection := dingtalk.C(dingTalkEntry.TableName())
	count, err := collection.Find(nil).Count()
	if err != nil {
		return err
	}

	if count > 0  {
		talkEntry := entry.DingTalkEntry{}
		err := collection.Find(nil).One(&talkEntry)
		tools.Logger().Info(talkEntry)
		if err != nil {
			return err
		}
		return collection.UpdateId(talkEntry.ID, dingTalkEntry)
	}


	return collection.Insert(dingTalkEntry)
}

func (this *DingTalkService) GetAccessTokenAndSecret() (entry.DingTalkEntry, error) {
	entry := entry.DingTalkEntry{}
	err := dingtalk.C(entry.TableName()).Find(nil).One(&entry)

	return entry, err
}

func (this *DingTalkService) StoreDingTalkTextMessage (messageEntry entry.DingTalkTextMessageEntry, callback func(entry entry.QueueEntry)) error {
	payload, err := messageEntry.Encode()
	if err != nil {
		return err
	}
	now := time.Now().Unix()
	queueEntry := entry.NewQueueEntry()
	queueEntry.ID = bson.NewObjectId()
	queueEntry.Type = entry.DingTalkTextMessage
	queueEntry.RetryNum = entry.RetryNum
	queueEntry.CreatedAt = now
	queueEntry.UpdatedAt = now
	queueEntry.Status = entry.PrepareStatus
	queueEntry.Payload = payload

	err = dingtalk.C(queueEntry.TableName()).Insert(queueEntry)
	if err != nil {
		return err
	}

	if callback != nil {
		callback(*queueEntry)
	}

	return nil
}

func (this *DingTalkService) SetAccessTokenAndSecret(accessToken, secret string) *DingTalkService {
	this.AccessToken = accessToken
	this.Secret = secret

	return this
}

func (this *DingTalkService) SetTextMessage(messageEntry entry.DingTalkTextMessageEntry) *DingTalkService {
	mobiles := message.Mobiles{
		AtMobiles: messageEntry.AtMobiles,
		IsAtAll: false,
	}

	this.message = message.NewTextMessage(message.WithText(message.Text{Content: messageEntry.Message}), message.WithMobiles(mobiles))

	return this
}