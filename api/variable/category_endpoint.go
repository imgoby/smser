package variable

import (
	"cn.sockstack/smser/pkg"
	"github.com/gin-gonic/gin"
)

type AddCategoryParams struct {
	Name string `form:"name"`
}

type AddCategoryResult struct {
	Status bool
}

func AddCategoryEndpoint(service CategoryService) pkg.Endpoint {
	return func(ctx *gin.Context, request interface{}) (response interface{}, err error) {
		return &AddCategoryResult{}, nil
	}
}
