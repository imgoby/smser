package v1

import (
	"cn.sockstack/smser/entry"
	"cn.sockstack/smser/internal"
	"cn.sockstack/smser/worker"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetWorkerNum(c *gin.Context)  {
	workerEntry := entry.NewWorkerEntry()
	err := c.ShouldBind(workerEntry)
	if err != nil {
		translate := internal.Translate(err)
		c.JSON(http.StatusOK, entry.NewOpenApiFailResponse("参数校验失败", translate))
		return
	}

	worker.Worker.SetNum(workerEntry.Number)
	worker.Worker.SetRst(true)

	c.JSON(http.StatusOK, entry.NewOpenApiSuccessResponse(entry.GetOpenApiResponseMessageByCode(entry.OpenApiSuccessCode), nil))
}

func GetWorkerNum(c *gin.Context)  {
	workerEntry := entry.NewWorkerEntry()
	workerEntry.Number = worker.Worker.GetNum()
	c.JSON(http.StatusOK, entry.NewOpenApiSuccessResponse(entry.GetOpenApiResponseMessageByCode(entry.OpenApiSuccessCode), workerEntry))
}
