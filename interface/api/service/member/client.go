package member

import (
	"github.com/fighthorse/sampleBookReader/domain/component/db"
	"github.com/fighthorse/sampleBookReader/domain/component/gocache"
	"github.com/fighthorse/sampleBookReader/domain/component/trace_redis"
	"github.com/fighthorse/sampleBookReader/interface/api/dao/redis"
	"time"
)

type Service struct {
	LocalCache *gocache.Cache
	Dao        *db.StoreDbClient
	RedisCache *trace_redis.RedisInstance
}

func New() (*Service, error) {

	return &Service{
		LocalCache: gocache.New(5*time.Minute, 10*time.Minute),
		Dao:        db.LoadDBByName("base"),
		RedisCache: redis.LoadOthersNew("base"),
	}, nil
}
