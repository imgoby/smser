package main

import (
	"cn.sockstack/smser/api"
	"cn.sockstack/smser/internal"
	"cn.sockstack/smser/middlewares"
	"cn.sockstack/smser/worker"
	"time"
)

func main() {
	worker.Worker.Start()
	time.Sleep(time.Second)
	go func() {
		time.Sleep(time.Second * 3)
		worker.Worker.Restart()
	}()
	internal.NewSmser().Route(api.R).Use(middlewares.RequestLog()).Run()
}