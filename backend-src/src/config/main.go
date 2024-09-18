package config

import (
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"src/global"
)

// 初始化数据库连接信息
func InputSqlInfo() bool {
	var Sqlconfig global.SqlConfig
	fmt.Println("[FeasOJ]请输入数据库连接信息：")
	fmt.Print("数据库名：")
	fmt.Scanln(&Sqlconfig.DbName)
	fmt.Print("数据库用户名：")
	fmt.Scanln(&Sqlconfig.DbUser)
	fmt.Print("数据库密码：")
	fmt.Scanln(&Sqlconfig.DbPassword)
	fmt.Print("数据库地址(加上端口)：")
	fmt.Scanln(&Sqlconfig.DbAddress)
	fmt.Println("[FeasOJ]正在保存数据库连接信息...")

	// 将连接信息作为xml保存到目录
	configXml, err := xml.Marshal(Sqlconfig)
	if err != nil {
		return false
	}
	filePath := filepath.Join(global.ConfigsDir, "sqlconfig.xml")
	os.WriteFile(filePath, configXml, 0644)
	return true
}

// 加载sqlconfig.xml并返回
func LoadSqlConfig() string {
	filePath := filepath.Join(global.ConfigsDir, "sqlconfig.xml")
	configFile, err := os.Open(filePath)
	if err != nil {
		return ""
	}
	var SqlConfig global.SqlConfig
	err = xml.NewDecoder(configFile).Decode(&SqlConfig)
	if err != nil {
		return ""
	}
	dsn := SqlConfig.DbUser + ":" + SqlConfig.DbPassword + "@tcp(" + SqlConfig.DbAddress + ")/" + SqlConfig.DbName + "?charset=utf8mb4&parseTime=True&loc=Asia%2FShanghai"
	return dsn
}

// 初始化Redis连接配置
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

// 加载redisconfig.xml并返回
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

// 加载emailconfig.xml并返回
func InitEmailConfig() global.MailConfig {
	// 判断是否有emailconfig.xml文件
	filePath := filepath.Join(global.ConfigsDir, "emailconfig.xml")
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		InputMailConfig()
	}
	configFile, err := os.Open(filePath)
	if err != nil {
		return global.MailConfig{}
	}
	defer configFile.Close()
	var mc global.MailConfig
	xml.NewDecoder(configFile).Decode(&mc)
	return mc
}

// 初始化邮箱配置
func InputMailConfig() {
	var hosts string
	var users string
	var password string
	fmt.Print("[FeasOJ]请输入SMTP服务器地址:")
	fmt.Scanln(&hosts)
	fmt.Print("[FeasOJ]请输入发件人邮箱地址(如114514@qq.com):")
	fmt.Scanln(&users)
	fmt.Print("[FeasOJ]请输入邮箱密码(不一定是登录密码):")
	fmt.Scanln(&password)

	// 写入配置到emailconfig.xml文件中
	config := global.MailConfig{Host: hosts, Port: 465, User: users, Pass: password}
	filePath := filepath.Join(global.ConfigsDir, "emailconfig.xml")
	configFile, _ := os.Create(filePath)
	defer configFile.Close()
	xml.NewEncoder(configFile).Encode(config)
}
