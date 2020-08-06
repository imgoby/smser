package services

import "testing"

func TestDingTalkService_Send(t *testing.T) {
	NewDingTalkService("").Send()
}
