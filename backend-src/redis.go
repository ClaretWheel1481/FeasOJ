package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

// TODO:连接Redis数据库，用于存储临时的邮箱验证码

type redisConfig struct {
	Address  string `xml:"address"`
	User     string `xml:"user"`
	Password string `xml:"password"`
}

// 写入Redis连接配置到redisconfig.xml中
func inputRedisInfo() bool {
	var config redisConfig
	fmt.Print("请输入Redis地址：")
	fmt.Scanln(&config.Address)
	fmt.Print("请输入Redis用户名(没有的话请直接回车)：")
	fmt.Scanln(&config.User)
	fmt.Print("请输入Redis密码：")
	fmt.Scanln(&config.Password)
	configXml, err := xml.Marshal(config)
	if err != nil {
		return false
	}
	os.WriteFile("redisconfig.xml", configXml, 0644)
	return true
}

// TODO：读取redisconfig.xml
// func loadRedisConfig() string {

// }

// TODO：连接到Redis
func initRedis() bool {
	// rdb := redis.NewClient(&redis.Options{
	// 	Addr:     "localhost:6379",
	// 	Password: "", // no password set
	// 	DB:       0,  // use default DB
	// })
	return true
}
