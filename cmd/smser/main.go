package main

import (
	"cn.sockstack/smser/api"
	"cn.sockstack/smser/internal"
	"cn.sockstack/smser/middlewares"
	"cn.sockstack/smser/worker"
)

func main() {
	go worker.Worker.Run()
	internal.NewSmser().Route(api.R).Use(middlewares.RequestLog()).Run()
}
