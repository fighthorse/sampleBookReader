package redis

import (
	"github.com/fighthorse/sampleBookReader/domain/component/trace_redis"
)

var (
	CategoryKey trace_redis.RedisKeys = "category:%s" // id
)
