package redisdb

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisDB struct {
	context context.Context
	client  *redis.Client
}
