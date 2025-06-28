package redisdb

import (
	"api/internal/config"
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	// DB is the exported database connection instance
	DB = New()
)

func New() *RedisDB {
	var db RedisDB

	// Create a new context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Configuration.Redis.Address,
		Password: config.Configuration.Redis.Password,
		DB:       0,
	})

	db.client = rdb
	db.context = ctx

	// Check the connection
	pingResult := rdb.Ping(ctx)
	if err := pingResult.Err(); err != nil {
		fmt.Println("Error pinging Redis:", err)
		return nil
	}

	return &db
}

/* func (r *RedisDB) Set(key string, value interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.client.Set(ctx, key, value, 0).Err()
	if err != nil {
		fmt.Println("Error setting key in Redis:", err)
		return err
	}
	return nil
} */

func (r *RedisDB) HSet(key string, fields map[string]interface{}) error {
	err := r.client.HSet(context.Background(), key, fields).Err()
	if err != nil {
		fmt.Println("Error setting hash in Redis:", err)
		return err
	}

	return nil
}
