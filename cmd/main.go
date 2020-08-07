package main

import (
	"cn.sockstack/smser/api"
	"cn.sockstack/smser/internal"
)

func main() {
	internal.NewSmser().Route(api.T, api.OpenApiV1).Run()
}