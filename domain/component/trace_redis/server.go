package trace_redis

import (
	"context"
	"time"
)

type RedisInstance struct {
	Client *RedisClient
	Name   string
}

func (r *RedisInstance) Select(ctx context.Context, db int) (*RedisInstance, error) {
	cfg := &RedisInstance{
		Client: NewClientDb(r.Name, db),
		Name:   r.Name,
	}
	return cfg, nil
}

func (r *RedisInstance) GetKey(ctx context.Context, key string) (string, error) {
	return r.Client.Get(ctx, key)
}

func (r *RedisInstance) HMSetKey(ctx context.Context, key string, mdata map[string]interface{}) (string, error) {
	return r.Client.HMSet(ctx, key, mdata)
}

func (r *RedisInstance) HMGetKey(ctx context.Context, key string, vals ...string) ([]interface{}, error) {
	return r.Client.HMGet(ctx, key, vals...)
}

func (r *RedisInstance) SIsMember(ctx context.Context, key string, iface interface{}) (bool, error) {
	return r.Client.SIsMember(ctx, key, iface)
}

func (r *RedisInstance) SAdd(ctx context.Context, key string, vals ...interface{}) (int64, error) {
	return r.Client.SAdd(ctx, key, vals...)
}

func (r *RedisInstance) SDiff(ctx context.Context, vals ...string) ([]string, error) {
	return r.Client.SDiff(ctx, vals...)
}

func (r *RedisInstance) SetBit(ctx context.Context, key string, i int64, i1 int) (int64, error) {
	return r.Client.SetBit(ctx, key, i, i1)
}

func (r *RedisInstance) LPush(ctx context.Context, key string, vals ...interface{}) (int64, error) {
	return r.Client.LPush(ctx, key, vals...)
}

func (r *RedisInstance) Incr(ctx context.Context, key string) (int64, error) {
	return r.Client.Incr(ctx, key)
}

func (r *RedisInstance) HDel(ctx context.Context, key string, vals ...string) (int64, error) {
	return r.Client.HDel(ctx, key, vals...)
}

func (r *RedisInstance) HExists(ctx context.Context, key string, val string) (bool, error) {
	return r.Client.HExists(ctx, key, val)
}

func (r *RedisInstance) HGet(ctx context.Context, key string, val string) (string, error) {
	return r.Client.HGet(ctx, key, val)
}

func (r *RedisInstance) HSet(ctx context.Context, key string, val string, iface interface{}) (bool, error) {
	return r.Client.HSet(ctx, key, val, iface)
}
func (r *RedisInstance) TTL(ctx context.Context, key string) (time.Duration, error) {
	return r.Client.TTL(ctx, key)
}

func (r *RedisInstance) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	return r.Client.HGetAll(ctx, key)
}
