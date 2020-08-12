package internal

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var (
	Cfg = NewConfig("config/config.ini")
)

type Smser struct {
	r *gin.Engine
	middlewares []gin.HandlerFunc
	routes []RouteHandleFunc
}

func NewSmser() *Smser {
	return &Smser{
		r: gin.Default(),
	}
}

func (this *Smser) Run(f func(app *Smser))  {
	if Cfg.Mode != "development" {
		gin.SetMode(gin.ReleaseMode)
	}

	this.r.Use(this.middlewares...)
	for _, o := range this.routes {
		o(this.r)
	}
	f(this)
	s := &http.Server{
		Addr:           fmt.Sprintf("%s:%s", Cfg.Host, Cfg.Port),
		Handler:        this.r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

type RouteHandleFunc func(r *gin.Engine)

func (this *Smser) Route(routes ...RouteHandleFunc) *Smser {
	this.routes = routes
	return this
}

func (this *Smser) Use(middlewares ...gin.HandlerFunc) *Smser {
	this.middlewares = middlewares
	return this
}

func init() {
	registerRequestTranslator()
}
