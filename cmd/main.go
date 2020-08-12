package main

import (
	"cn.sockstack/smser/api"
	"cn.sockstack/smser/internal"
	"cn.sockstack/smser/middlewares"
	"cn.sockstack/smser/worker"
)

func main() {
	internal.NewSmser().Route(api.R).Use(middlewares.RequestLog()).Run(func(app *internal.Smser) {
		go worker.Worker.Run()
		go worker.Retry()
	})
}
