package api

import (
	"cn.sockstack/smser/pkg/service"
	"github.com/gin-gonic/gin"
)

func CreateRequest() service.EncodeRequestFunc {
	return func(ctx *gin.Context) (request interface{}, err error) {
		r := &Request{}
		return r, nil
	}
}

func CreateResponse() service.DecodeResponseFunc {
	return func(ctx *gin.Context, response interface{}, err error) {
		resp := response.(*Response)
		ctx.JSON(200, resp)
	}
}
