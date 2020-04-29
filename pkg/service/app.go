package service

import (
	"github.com/gin-gonic/gin"
)

type Endpoint func(ctx *gin.Context, request interface{}) (response interface{}, err error)

type EncodeRequestFunc func(ctx *gin.Context) (request interface{}, err error)

type DecodeResponseFunc func(ctx *gin.Context, response interface{}, err error)

func RegisterService(endpoint Endpoint, requestFunc EncodeRequestFunc, responseFunc DecodeResponseFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req, err := requestFunc(ctx)
		if err != nil {
			responseFunc(ctx, nil, err)
			return
		}

		response, err := endpoint(ctx, req)

		responseFunc(ctx, response, err)
	}
}

var service *app

func NewApp(router Router) *app {
	if service != nil {
		return service
	}

	return &app{r:router, gin: gin.Default()}
}

type app struct {
	r Router
	gin *gin.Engine
}

func (app *app)registerHandler()  {
	for path, r := range app.r {
		app.gin.Handle(r.Method, path, r.handler)
	}
}

func (app *app) Run(addr... string) {
	app.registerHandler()
	app.gin.Run(addr...)
}
