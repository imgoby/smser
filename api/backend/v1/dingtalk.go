package v1

import (
	"cn.sockstack/smser/entry"
	"cn.sockstack/smser/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	service = services.NewDingTalkService()
)

//DingTalkSecretAndAccessTokenStore 添加 DingTalk 的 AccessToken 和 Secret
func StoreDingTalkSecretAndAccessToken(c *gin.Context)  {
	err := service.StoreAccessTokenAndSecret(*entry.NewDingTalkEntry("Test", "secret"))
	if err != nil {
		c.JSON(http.StatusOK, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func GetTalkSecretAndAccessToken(c *gin.Context)  {
	entry, err := service.GetAccessTokenAndSecret()
	if err != nil {
		c.JSON(http.StatusOK, err)
		return
	}

	c.JSON(http.StatusOK, entry)
}
