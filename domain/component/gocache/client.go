package gocache

import (
	"time"
)

var (
	gouache = New(30*time.Minute, 60*time.Minute)
)

func Set(key string, val interface{}, ttl time.Duration) {
	gouache.Set(key, val, ttl)
}

func Get(key string) (interface{}, bool) {
	return gouache.Get(key)
}

func Del(key string) {
	gouache.Delete(key)
}

func Flush() {
	gouache.Flush()
}
