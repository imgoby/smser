package v1

import (
	"cn.sockstack/smser/entry"
	"cn.sockstack/smser/internal"
	"cn.sockstack/smser/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	service = services.NewDingTalkService()
)

//DingTalkSecretAndAccessTokenStore 添加 DingTalk 的 AccessToken 和 Secret
func StoreDingTalkSecretAndAccessToken(c *gin.Context)  {
	talkEntry := entry.DingTalkEntry{}
	err := c.ShouldBind(&talkEntry)
	if err != nil {
		translate := internal.Translate(err)
		c.JSON(http.StatusOK, entry.NewOpenApiFailResponse("参数校验失败", translate))
		return
	}

	err = service.StoreAccessTokenAndSecret(talkEntry)
	if err != nil {
		c.JSON(http.StatusOK, entry.NewOpenApiFailResponse("保存失败", err))
		return
	}

	c.JSON(http.StatusOK, entry.NewOpenApiSuccessResponse(entry.GetOpenApiResponseMessageByCode(entry.OpenApiSuccessCode), nil))
}

func GetTalkSecretAndAccessToken(c *gin.Context)  {
	entry, err := service.GetAccessTokenAndSecret()
	if err != nil {
		c.JSON(http.StatusOK, err)
		return
	}

	c.JSON(http.StatusOK, entry)
}
