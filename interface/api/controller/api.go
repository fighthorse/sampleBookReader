package controller

import (
	"github.com/fighthorse/sampleBookReader/interface/api/controller/amap"
	"github.com/fighthorse/sampleBookReader/interface/api/controller/hc"
	"github.com/fighthorse/sampleBookReader/interface/api/controller/login"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) error {
	r.Static("/pages", "./../pages")
	r.StaticFile("/favicon.ico", "./../pages/favicon.ico")
	r.GET("/", func(c *gin.Context) {
		c.Request.URL.Path = "/pages"
		r.HandleContext(c)
	})
	hc.RegisterHttp(r)
	login.RegisterHttp(r)
	amap.RegisterHttp(r)

	return nil
}
