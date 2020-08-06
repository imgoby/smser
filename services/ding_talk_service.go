package services

import "github.com/royeo/dingrobot"

type DingTalkService struct {
	webhook string
}

func NewDingTalkService(webhook string) *DingTalkService {
	return &DingTalkService{webhook: webhook}
}

func (this *DingTalkService) Send ()  {
	robot := dingrobot.NewRobot(this.webhook)
	err := robot.SendText("ok", []string{}, false)
	if err != nil {
		panic(err)
	}
}
