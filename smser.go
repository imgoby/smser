package main

import (
	"cn.sockstack/smser/api"
	"cn.sockstack/smser/pkg/service"
	"github.com/gin-gonic/gin"
)

func main() {
	router := service.Router{}
	router.Route("GET", "/test", service.RegisterService(api.Endpoint(api.Service{}), api.CreateRequest(), api.CreateResponse()))
	router.Route("GET", "/", func(context *gin.Context) {
		context.JSON(200, "ok")
	})
	app := service.NewApp(router)
	app.Run("0.0.0.0:8081")
}
