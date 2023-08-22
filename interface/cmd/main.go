package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"
	"net/http"
	"path/filepath"

	"github.com/fighthorse/sampleBookReader/interface/api/conf"
	"github.com/fighthorse/sampleBookReader/interface/api/controller"
	"github.com/fighthorse/sampleBookReader/interface/api/dao"
	"github.com/fighthorse/sampleBookReader/interface/api/dao/httpserver"
	"github.com/fighthorse/sampleBookReader/interface/api/dao/middleware"
	"github.com/fighthorse/sampleBookReader/interface/api/dao/mysql"
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

	// 找到并读取配置文件并且 处理错误读取配置文件
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := conf.Init(v); err != nil {
		panic(err)
	}
}

func main() {

	flag.Parse()
	// 如果env使用的是绝对路径，则configpath为路径，env为文件名
	if filepath.IsAbs(*env) {
		confidant, *env = filepath.Split(*env)
	}
	//config init
	readConfig(*env, confidant)

	// 初始化基础组件
	dao.InitComponent()

	// 加载依赖组件
	// http
	httpserver.Init()
	//redis
	redis.Init()
	//mysql
	mysql.Init()
	// start server
	//gin.SetMode(gin.ReleaseMode)
	// Creates a router without any middleware by default
	r := gin.New()
	// Global middleware
	// Recovery middleware
	r.Use(middleware.CustomRecovery)
	//instrument api count
	r.Use(middleware.Instrument) // defer
	// middle
	r.Use(middleware.Trace)
	//access log
	r.Use(middleware.AccessLogging) //defer

	// 初始化api依赖的各个模块
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
