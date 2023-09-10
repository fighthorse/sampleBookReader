package redis

import (
	"github.com/fighthorse/sampleBookReader/domain/component/trace_redis"
)

var (
	CategoryKey trace_redis.RedisKeys = "category:%s" // id
	// LockRegisterNameKey login register key
	LockRegisterNameKey trace_redis.RedisKeys = "login:register:%s" // id

	LoginUidKey trace_redis.RedisKeys = "login:%d"
	ShelfUidKey trace_redis.RedisKeys = "shelf:%d"
	FeedUidKey  trace_redis.RedisKeys = "feed:%d"

	CommentUidKey trace_redis.RedisKeys = "comment:%d"
)
