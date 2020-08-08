package api

import (
	v1 "cn.sockstack/smser/api/backend/v1"
	openapiv1 "cn.sockstack/smser/api/openapi/v1"
	"github.com/gin-gonic/gin"
	"net/http"
)

func R(r *gin.Engine)  {
	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{"message":"ok"})
	})
	OpenApiV1(r)
	BackendApiV1(r)
}

func OpenApiV1(r *gin.Engine)  {
	g := r.Group("/api/open/v1")
	g.POST("/send", openapiv1.Send)

	dingtalk := g.Group("/dingtalk")
	dingtalk.POST("/text", openapiv1.DingTalkTextMessageSend)
}

func BackendApiV1(r *gin.Engine)  {
	g := r.Group("/api/backend/v1")
	dingtalk := g.Group("/dingtalk")
	dingtalk.POST("/store", v1.StoreDingTalkSecretAndAccessToken)
	dingtalk.GET("/", v1.GetTalkSecretAndAccessToken)
}
