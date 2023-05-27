package trace_redis

import "fmt"

type RedisKeys string

func (r RedisKeys) Fmt(params ...interface{}) string {
	return fmt.Sprintf(r.ToString(), params...)
}
func (r RedisKeys) ToString() string {
	return string(r)
}
