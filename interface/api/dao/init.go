package dao

import (
	"github.com/fighthorse/sampleBookReader/domain/component/db"
	"github.com/fighthorse/sampleBookReader/domain/component/httpclient"
	"github.com/fighthorse/sampleBookReader/domain/component/log"
	"github.com/fighthorse/sampleBookReader/domain/component/trace"
	"github.com/fighthorse/sampleBookReader/domain/component/trace_redis"
	"github.com/fighthorse/sampleBookReader/interface/api/conf"
)

func InitComponent() {
	// redis cfg
	trace_redis.InitCfg(conf.GConfig.Redis)
	// mysql cfg
	_ = db.NewDB(conf.GConfig.Mysql)
	// http cfg
	httpclient.InitSelfService(conf.GConfig.HttpServer.SelfServiceName, conf.GConfig.HttpServer.CloseBreaker)
	httpclient.InitCircuitBreaker(conf.GConfig.HttpBreaker)
	httpclient.InitChildService(conf.GConfig.ChildServer)
	//trace
	trace.Init()
	// log
	log.Init(conf.GConfig.Log)
}
