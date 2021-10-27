package redisclient

import (
	"context"
	"hermes/socket/config"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	client redis.Client
}

func NewRedisClient(env *config.Env) *RedisClient {
	return &RedisClient{
		client: *redis.NewClient(&redis.Options{
			Addr:     env.REDIS_URI,
			Password: env.REDIS_PASSWORD,
			DB:       env.REDIS_DB,
		}),
	}
}

func (r *RedisClient) Publish(channel string, message interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	res := r.client.Publish(ctx, channel, message)
	if err := res.Err(); err != nil {
		return err
	}
	return nil
}
