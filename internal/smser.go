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
}

func NewSmser() *Smser {
	return &Smser{
		r: gin.Default(),
	}
}

func (this *Smser) Run()  {
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
	for _, handle := range routes{
		handle(this.r)
	}

	return this
}

func (this *Smser) Use(middlewares ...gin.HandlerFunc) *Smser {
	this.r.Use(middlewares...)
	return this
}
