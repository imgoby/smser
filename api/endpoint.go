package api

import (
	"cn.sockstack/smser/pkg/service"
	"github.com/gin-gonic/gin"
)

type Request struct {

}

type Response struct {
	Result string
}

func Endpoint(service Service) service.Endpoint {
	return func(ctx *gin.Context, request interface{}) (response interface{}, err error) {
		service.test()
		return &Response{Result: "结果"}, nil
	}
}
