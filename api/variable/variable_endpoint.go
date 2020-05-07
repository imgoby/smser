package variable

import (
	"cn.sockstack/smser/src"
	"github.com/gin-gonic/gin"
)

type AddParams struct {
	Name string `form:"name"`
	Sign string `form:"sign"`
	Category string `form:"category"`
}

type AddResult struct {
	Status bool
}

func AddVariableEndpoint(service VariableService) src.Endpoint {
	return func(ctx *gin.Context, request interface{}) (response interface{}, err error) {

		return &AddResult{}, nil
	}
}