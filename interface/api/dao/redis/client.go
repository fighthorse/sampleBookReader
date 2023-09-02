package redis

import (
	"context"
	"fmt"
	"github.com/fighthorse/sampleBookReader/domain/component/trace_redis"
)

var (
	Others = map[string]*trace_redis.RedisInstance{}
)

func Init(names ...string) {
	for _, v := range names {
		Others[v] = LoadOthersNew(v)
	}
}

func GetClientByName(name string) *trace_redis.RedisInstance {
	if l, ok := Others[name]; ok {
		return l
	}
	return nil
}

func LoadOthersNew(name string) *trace_redis.RedisInstance {
	if name == "" {
		return nil
	}
	cfg := &trace_redis.RedisInstance{}
	cfg.Name = name
	cfg.Client = trace_redis.NewClient(cfg.Name)
	Others[name] = cfg
	return cfg
}

func LoadOthersDB(name string, db int) *trace_redis.RedisInstance {
	dbZero, ok := Others[name]
	if !ok {
		return nil
	}
	if db == 0 {
		return dbZero
	}
	nameNew := fmt.Sprintf("%s_%d", name, db)
	if ll, ok := Others[nameNew]; ok {
		return ll
	}
	ll, err := dbZero.Select(context.Background(), db)
	if err != nil {
		return nil
	}
	Others[nameNew] = ll
	return ll
}
