package utils

import (
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"src/global"

	"github.com/go-redis/redis"
)

// 写入Redis连接配置到redisconfig.xml中
func InputRedisInfo() bool {
	var config global.RedisConfig
	fmt.Print("请输入Redis地址：")
	fmt.Scanln(&config.Address)
	fmt.Print("请输入Redis密码：")
	fmt.Scanln(&config.Password)
	configXml, err := xml.Marshal(config)
	if err != nil {
		return false
	}
	filePath := filepath.Join(global.ConfigsDir, "redisconfig.xml")
	os.WriteFile(filePath, configXml, 0644)
	return true
}

func LoadRedisConfig() global.RedisConfig {
	filePath := filepath.Join(global.ConfigsDir, "redisconfig.xml")
	var config global.RedisConfig
	configXml, err := os.ReadFile(filePath)
	if err != nil {
		return config
	}
	xml.Unmarshal(configXml, &config)
	return config
}

// 连接到Redis
func InitRedis() *redis.Client {
	filePath := filepath.Join(global.ConfigsDir, "redisconfig.xml")
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		InputRedisInfo()
	}
	config := LoadRedisConfig()
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Address,
		Password: config.Password,
		DB:       0,
	})
	return rdb
}
