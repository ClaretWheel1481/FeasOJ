package utils

import (
	"os"
	"path/filepath"
	"src/config"
	"src/global"

	"github.com/go-redis/redis"
)

// 连接到Redis并返回redis.Client对象
func InitRedis() *redis.Client {
	filePath := filepath.Join(global.ConfigsDir, "redisconfig.xml")
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		config.InputRedisInfo()
	}
	config := config.LoadRedisConfig()
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Address,
		Password: config.Password,
		DB:       0,
	})
	return rdb
}
