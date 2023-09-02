package book

import (
	"github.com/fighthorse/sampleBookReader/domain/component/gocache"
	"github.com/fighthorse/sampleBookReader/domain/component/trace_redis"
	"github.com/fighthorse/sampleBookReader/interface/api/dao/redis"
	"time"
)

type Service struct {
	Cache      *trace_redis.RedisInstance
	LocalCache *gocache.Cache
}

func New() (*Service, error) {
	return &Service{
		Cache:      redis.GetClientByName("base"),
		LocalCache: gocache.New(5*time.Minute, 10*time.Minute),
	}, nil
}
