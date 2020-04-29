package variable

import (
	"cn.sockstack/smser/pkg"
	"github.com/gin-gonic/gin"
)

func CreateAddParams() pkg.EncodeRequestFunc {
	return func(ctx *gin.Context) (request interface{}, err error) {
		params := AddParams{}
		err = ctx.ShouldBind(&params)
		if err != nil {
			return nil, err
		}

		return params, nil
	}
}

func CreateAddResult() pkg.DecodeResponseFunc {
	return func(ctx *gin.Context, response interface{}, err error) {
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		ctx.JSON(200, response)
	}
}
