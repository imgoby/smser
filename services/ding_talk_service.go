package services

import (
	"cn.sockstack/smser/entry"
	"cn.sockstack/smser/internal"
	"cn.sockstack/smser/tools"
	"encoding/json"
	"fmt"
	"github.com/sockstack/dtrobot"
	"github.com/sockstack/dtrobot/message"
	"github.com/sockstack/dtrobot/pkg/encode"
	"gopkg.in/mgo.v2/bson"
	"time"
)

var (
	dingtalk = internal.GetMgoDB()
)

type DingTalkService struct {
	AccessToken string
	Secret      string
	message     message.Message
	MessageID 	bson.ObjectId
}

func NewDingTalkService() *DingTalkService {
	return &DingTalkService{}
}

func (this *DingTalkService) Send () error {
	robot, err := dtrobot.NewRobot(this.AccessToken, dtrobot.WithSecret(this.Secret))
	if err != nil {
		go tools.MessageLogger(err, nil)
		return nil
	}

	response, err := robot.Send(this.message)

	if err != nil {
		go tools.MessageLogger(err, nil)
		return nil
	}

	errcode, ok := response["errcode"]
	errcode = fmt.Sprintf("%v", errcode)
	if errcode != "0" && ok {
		robot.Options.Sign = encode.CalcSign(time.Now().Unix() * 1000, robot.Options.Secret)
		return this.Send()
	}

	message, _ := json.Marshal(this.message)
	go tools.MessageLogger(string(message), response)

	status := response["status"]
	if status != nil {
		//return errors.New("发送失败")
		return nil
	}

	tools.QueueAckRecordByMessageID(this.MessageID)

	return nil
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
	id := bson.NewObjectId()
	messageEntry.MessageId = id
	payload, err := messageEntry.Encode()
	if err != nil {
		return err
	}
	now := time.Now().Unix()
	queueEntry := entry.NewQueueEntry()
	queueEntry.ID = id
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
