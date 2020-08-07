package services

import (
	"fmt"
	"github.com/sockstack/dtrobot"
	"github.com/sockstack/dtrobot/message"
	"time"
)

type DingTalkService struct {
	AccessToken string
	Secret string
}

func NewDingTalkService() *DingTalkService {
	return &DingTalkService{
		AccessToken: "79a01c796146d462b59bd6befc8d43e2c87dc446218d8757acca10d752c4fa03",
		Secret: "SEC5a7d0e259fd08f1f19573101617713dcb19e5733b69f158969beaa723648410d",
	}
}

func (this *DingTalkService) Send ()  {
	robot, err := dtrobot.NewRobot(this.AccessToken, dtrobot.WithSecret(this.Secret))
	if err != nil {
		panic(err)
	}

	textMessage := message.NewTextMessage(message.WithText(message.Text{Content: time.Now().String()}))
	send, err := robot.Send(textMessage)

	if err != nil {
		panic(err)
	}

	fmt.Println(send)
}
