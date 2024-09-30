package utils

import (
	"src/config"

	"github.com/go-redis/redis"
)

// 连接到Redis并返回redis.Client对象
func ConnectRedis() *redis.Client {
	config := config.LoadRedisConfig()
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Address,
		Password: config.Password,
		DB:       0,
	})
	return rdb
}
