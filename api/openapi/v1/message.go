package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Send(c *gin.Context)  {
	c.JSON(http.StatusOK, http.StatusOK)
}
