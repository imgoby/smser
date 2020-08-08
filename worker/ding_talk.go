package worker

import (
	"cn.sockstack/smser/entry"
	"cn.sockstack/smser/services"
	"cn.sockstack/smser/tools"
	"errors"
)

func SendDingTalkTextMessage(queueEntry entry.QueueEntry) error {
	service := services.NewDingTalkService()

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

	service.SetAccessTokenAndSecret(dingTalkEntry.AccessToken, dingTalkEntry.Secret).SetTextMessage(*messageEntry).Send()
	return nil
}
