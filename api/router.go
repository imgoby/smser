package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func T(r *gin.Engine)  {
	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{"message":"ok"})
	})
}
