package main

import (
	"cn.sockstack/smser/api"
	"cn.sockstack/smser/src"
)

func main() {
	app := src.NewApp()
	api.V1(app)
	app.Run("0.0.0.0:8081")
}
