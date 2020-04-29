package pkg

import "github.com/gin-gonic/gin"

type R struct {
	Method string
	handler gin.HandlerFunc
}

type Router map[string]R

func (r Router)Route(method, path string, handlerFunc gin.HandlerFunc ) Router {
	r[path] = R{
		Method:  method,
		handler: handlerFunc,
	}
	return r
}

