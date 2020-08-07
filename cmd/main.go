package main

import (
	"cn.sockstack/smser/api"
	"cn.sockstack/smser/internal"
	"cn.sockstack/smser/middlewares"
)

func main() {
	internal.NewSmser().Route(api.T, api.OpenApiV1).Use(middlewares.RequestLog()).Run()
}