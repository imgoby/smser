package worker

import (
	"cn.sockstack/smser/entry"
	"cn.sockstack/smser/internal"
	"cn.sockstack/smser/services"
)

func SendDingTalkTextMessage(queueEntry entry.QueueEntry) error {
	service := services.NewDingTalkService()

	messageEntry := entry.NewDingTalkTextMessageEntry()
	messageEntry.Decode([]byte(queueEntry.Payload))

	service.SetAccessTokenAndSecret(internal.Cfg.AccessToken, internal.Cfg.Secret).SetTextMessage(*messageEntry).Send()
	return nil
}
