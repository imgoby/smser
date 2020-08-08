package worker

import (
	"cn.sockstack/smser/entry"
	"cn.sockstack/smser/services"
	"cn.sockstack/smser/tools"
	"errors"
)

func SendDingTalkTextMessage(queueEntry entry.QueueEntry) error {
	if queueEntry.Status == entry.AckStatus {
		return nil
	}
	service := services.NewDingTalkService()
	service.MessageID = queueEntry.ID

	messageEntry := entry.NewDingTalkTextMessageEntry()
	messageEntry.Decode([]byte(queueEntry.Payload))


	dingTalkEntry, err := service.GetAccessTokenAndSecret()
	if err != nil {
		tools.Logger().Error(err)
	}

	if dingTalkEntry.AccessToken == "" {
		tools.Logger().Error("access_token 为空")
		return errors.New("access_token 为空")
	}

	err = service.SetAccessTokenAndSecret(dingTalkEntry.AccessToken, dingTalkEntry.Secret).SetTextMessage(*messageEntry).Send()
	if err != nil {
		tools.RetryRecord(queueEntry)
		return err
	}
	return nil
}
