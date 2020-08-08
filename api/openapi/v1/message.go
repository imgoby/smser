package v1

import (
	"cn.sockstack/smser/entry"
	"cn.sockstack/smser/internal"
	"cn.sockstack/smser/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Send(c *gin.Context)  {
	c.JSON(http.StatusOK, http.StatusOK)
}

func DingTalkTextMessageSend(c *gin.Context)  {
	dingTalkTextMessageEntry := entry.NewDingTalkTextMessageEntry()

	err := c.ShouldBind(dingTalkTextMessageEntry)
	if err != nil {
		translate := internal.Translate(err)
		c.JSON(http.StatusOK, entry.NewOpenApiFailResponse("参数校验失败", translate))
		return
	}

	service := services.NewDingTalkService()
	err = service.StoreDingTalkTextMessage(*dingTalkTextMessageEntry, func(entry entry.QueueEntry) {
		services.NewQueueService().Push(entry)
	})
	if err != nil {
		c.JSON(http.StatusOK, entry.NewOpenApiFailResponse("消息持久化失败", nil))
		return
	}

	c.JSON(http.StatusOK, entry.NewOpenApiSuccessResponse(entry.GetOpenApiResponseMessageByCode(entry.OpenApiSuccessCode), nil))
}
