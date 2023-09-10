package main

import (
	"flag"
	"fmt"
	"github.com/fighthorse/sampleBookReader/interface/api/service"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"

	"github.com/fighthorse/sampleBookReader/interface/api/conf"
	"github.com/fighthorse/sampleBookReader/interface/api/controller"
	"github.com/fighthorse/sampleBookReader/interface/api/dao"
	"github.com/fighthorse/sampleBookReader/interface/api/dao/httpserver"
	"github.com/fighthorse/sampleBookReader/interface/api/dao/middleware"
	"github.com/fighthorse/sampleBookReader/interface/api/dao/redis"
)

var (
	v   = viper.New()
	env = flag.String("env", "local", "config file name")
	//graceful = flag.Bool("graceful", false, "listen on fd open 3 (internal use only)")
	confidant = "./../config"
	// BuildID is compile-time variable
	BuildID = "0"
)

func readConfig(fileName, filePath string) {
	v.SetConfigName(fileName)
	v.AddConfigPath(filePath)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := conf.Init(v); err != nil {
		panic(err)
	}
}
func main() {
	flag.Parse()
	if filepath.IsAbs(*env) {
		confidant, *env = filepath.Split(*env)
	}
	//config init
	readConfig(*env, confidant)
	dao.InitComponent()
	// http
	httpserver.Init()
	//redis
	redis.Init("base")
	//init service
	_ = service.InitService()

	// start server
	//gin.SetMode(gin.ReleaseMode)
	gin.SetMode(gin.DebugMode)
	r := gin.New()

	// Global middleware
	r.Use(middleware.CustomRecovery)
	//instrument api count
	r.Use(middleware.Instrument) // defer
	// middle
	r.Use(middleware.Trace)
	//access log
	r.Use(middleware.AccessLogging) //defer
	if err := controller.Init(r); err != nil {
		panic(err)
	}
	//router.Run(":3000")
	srv := conf.GConfig.Transport

	if srv.InnerHTTP.Addr != "" {
		go InnerServer(srv.InnerHTTP)
	}
	err := r.Run(srv.HTTP.Addr)
	if err != nil {
		fmt.Printf("err:%s", err.Error())
	}
}

func InnerServer(cf conf.HTTPConfig) {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		// panic with a string -- the custom middleware could save this to a database or report it to the user
		c.String(http.StatusOK, "OK")
	})

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	r.GET("/metrics", func(c *gin.Context) {
		promhttp.Handler().ServeHTTP(c.Writer, c.Request)
	})

	// Listen and serve on 0.0.0.0:8080
	err := r.Run(cf.Addr)
	if err != nil {
		fmt.Printf("Inner err:%s", err.Error())
	}
}
