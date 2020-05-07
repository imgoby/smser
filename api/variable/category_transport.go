package variable

import (
	"cn.sockstack/smser/src"
	"github.com/gin-gonic/gin"
)

func CreateAddParams() src.EncodeRequestFunc {
	return func(ctx *gin.Context) (request interface{}, err error) {
		params := AddParams{}
		err = ctx.ShouldBind(&params)
		if err != nil {
			return nil, err
		}

		return params, nil
	}
}

func CreateAddResult() src.DecodeResponseFunc {
	return func(ctx *gin.Context, response interface{}, err error) {
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		ctx.JSON(200, response)
	}
}
