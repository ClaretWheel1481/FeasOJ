package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-redis/redis"
)

type redisConfig struct {
	Address  string `xml:"address"`
	Password string `xml:"password"`
}

// 写入Redis连接配置到redisconfig.xml中
func inputRedisInfo() bool {
	var config redisConfig
	fmt.Print("请输入Redis地址：")
	fmt.Scanln(&config.Address)
	fmt.Print("请输入Redis密码：")
	fmt.Scanln(&config.Password)
	configXml, err := xml.Marshal(config)
	if err != nil {
		return false
	}
	filePath := filepath.Join(configsDir, "redisconfig.xml")
	os.WriteFile(filePath, configXml, 0644)
	return true
}

func loadRedisConfig() redisConfig {
	filePath := filepath.Join(configsDir, "redisconfig.xml")
	var config redisConfig
	configXml, err := os.ReadFile(filePath)
	if err != nil {
		return config
	}
	xml.Unmarshal(configXml, &config)
	return config
}

// 连接到Redis
func initRedis() *redis.Client {
	filePath := filepath.Join(configsDir, "redisconfig.xml")
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		inputRedisInfo()
	}
	config := loadRedisConfig()
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Address,
		Password: config.Password,
		DB:       0,
	})
	return rdb
}
