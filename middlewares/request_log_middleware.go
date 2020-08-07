package middlewares

import (
	"cn.sockstack/smser/tools"
	"github.com/gin-gonic/gin"
)

func RequestLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		tools.RequestLogger().Info(c.Request.URL)
	}
}
