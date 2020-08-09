package v1

import (
	"cn.sockstack/smser/entry"
	"cn.sockstack/smser/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MessageList(c *gin.Context)  {
	queueEntry := entry.NewQueueEntry()
	c.ShouldBind(queueEntry)
	queueService := services.NewQueueService()
	list, err := queueService.MessageList(*queueEntry)
	if err != nil {
		c.JSON(http.StatusOK, entry.NewOpenApiFailResponse("获取消息失败", err))
		return
	}

	c.JSON(http.StatusOK, entry.NewOpenApiSuccessResponse(entry.GetOpenApiResponseMessageByCode(entry.OpenApiSuccessCode), list))
}
