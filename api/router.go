package api

import (
	openapiv1 "cn.sockstack/smser/api/openapi/v1"
	"github.com/gin-gonic/gin"
	"net/http"
)

func T(r *gin.Engine)  {
	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{"message":"ok"})
	})
}

func OpenApiV1(r *gin.Engine)  {
	g := r.Group("/api/v1")
	g.POST("/send", openapiv1.Send)
}
