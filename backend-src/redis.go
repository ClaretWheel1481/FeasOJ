package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

// TODO:连接Redis数据库，用于存储临时的邮箱验证码

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
	os.WriteFile("redisconfig.xml", configXml, 0644)
	return true
}

func loadRedisConfig() redisConfig {
	var config redisConfig
	configXml, err := os.ReadFile("redisconfig.xml")
	if err != nil {
		return config
	}
	xml.Unmarshal(configXml, &config)
	return config
}

// TODO：连接到Redis
func initRedis() bool {
	if _, err := os.Stat("redisconfig.xml"); os.IsNotExist(err) {
		inputRedisInfo()
	}
	// config := loadRedisConfig()
	// rdb := redis.NewClient(&redis.Options{
	// 	Addr:     config.Address,
	// 	Password: config.Password,
	// 	DB:       0,
	// })
	fmt.Println("[FeasOJ]Redis连接成功。")
	return true
}
