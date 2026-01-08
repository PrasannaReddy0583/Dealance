package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	client *redis.Client
}

func NewRedis(addr string) *Redis {
	return &Redis{
		client: redis.NewClient(&redis.Options{
			Addr: addr,
		}),
	}
}

func (r *Redis) Set(
	ctx context.Context,
	key string,
	value string,
	ttl time.Duration,
) error {
	return r.client.Set(ctx, key, value, ttl).Err()
}

func (r *Redis) Get(
	ctx context.Context,
	key string,
) (string, error) {
	return r.client.Get(ctx, key).Result()
}

func (r *Redis) Del(
	ctx context.Context,
	key string,
) error {
	return r.client.Del(ctx, key).Err()
}
