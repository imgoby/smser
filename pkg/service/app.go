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

var service *App

func NewApp() *App {
	if service != nil {
		return service
	}

	return &App{R:Router{}, gin: gin.Default()}
}

type App struct {
	R Router
	gin *gin.Engine
}

func (app *App)registerHandler()  {
	for path, r := range app.R {
		app.gin.Handle(r.Method, path, r.handler)
	}
}

func (app *App) Run(addr... string) {
	app.registerHandler()
	_ = app.gin.Run(addr...)
}
