package main

import (
	"cn.sockstack/smser/api"
	"cn.sockstack/smser/pkg"
)

func main() {
	app := pkg.NewApp()
	api.V1(app)
	app.Run("0.0.0.0:8081")
}
