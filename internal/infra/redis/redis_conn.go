package redis

import (
	"context"

	"github.com/ahargunyllib/thera-be/internal/infra/env"
	"github.com/ahargunyllib/thera-be/pkg/log"
	"github.com/redis/go-redis/v9"
)

func NewRedisConn() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     env.AppEnv.RedisAddress,
		Password: env.AppEnv.RedisPassword,
		DB:       0,
	})

	ping, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Panic(log.CustomLogInfo{
			"error": err.Error(),
		}, "[REDIS][NewRedisConn] failed to connect to redis")
	}

	log.Info(log.CustomLogInfo{
		"ping": ping,
	}, "[REDIS][NewRedisConn] connected to redis")

	return client
}
