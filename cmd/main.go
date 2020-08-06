package main

import (
	"cn.sockstack/smser/api"
	"cn.sockstack/smser/internal"
	"github.com/gin-gonic/gin"
)

func main() {
	internal.NewSmser().Route(api.T).Use(gin.Logger()).Run()
}