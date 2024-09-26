package config

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"src/global"
)

// 写入xml文件
func writeXMLToFile(filePath string, data interface{}) error {
	configXml, err := xml.Marshal(data)
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, configXml, 0644)
}

// 读取xml文件
func readXMLFromFile(filePath string, data interface{}) error {
	configFile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer configFile.Close()
	return xml.NewDecoder(configFile).Decode(data)
}

// 初始化数据库连接信息
func InputSqlInfo() bool {
	var sqlConfig global.SqlConfig
	log.Println("[FeasOJ]Please input the MySQL connection configuration：")
	log.Print("[FeasOJ]Database name: ")
	fmt.Scanln(&sqlConfig.DbName)
	log.Print("[FeasOJ]Database address with port: ")
	fmt.Scanln(&sqlConfig.DbAddress)
	log.Print("[FeasOJ]Database user: ")
	fmt.Scanln(&sqlConfig.DbUser)
	log.Print("[FeasOJ]Database password: ")
	fmt.Scanln(&sqlConfig.DbPassword)
	log.Println("[FeasOJ]Saving the connection configuration...")

	filePath := filepath.Join(global.ConfigsDir, "sqlconfig.xml")
	if err := writeXMLToFile(filePath, sqlConfig); err != nil {
		log.Println("Error saving SQL config: ", err)
		return false
	}
	return true
}

// 加载sqlconfig.xml并返回
func LoadSqlConfig() string {
	filePath := filepath.Join(global.ConfigsDir, "sqlconfig.xml")
	var sqlConfig global.SqlConfig
	if err := readXMLFromFile(filePath, &sqlConfig); err != nil {
		log.Println("Error loading SQL config: ", err)
		return ""
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Asia%%2FShanghai",
		sqlConfig.DbUser, sqlConfig.DbPassword, sqlConfig.DbAddress, sqlConfig.DbName)
	return dsn
}

// 初始化Redis连接配置
func InputRedisInfo() bool {
	var config global.RedisConfig
	log.Println("[FeasOJ]Please input the Redis connection configuration：")
	fmt.Print("[FeasOJ]Redis address with port: ")
	fmt.Scanln(&config.Address)
	fmt.Print("[FeasOJ]Redis password: ")
	fmt.Scanln(&config.Password)

	filePath := filepath.Join(global.ConfigsDir, "redisconfig.xml")
	if err := writeXMLToFile(filePath, config); err != nil {
		log.Println("Error saving Redis config: ", err)
		return false
	}
	return true
}

// 加载redisconfig.xml并返回
func LoadRedisConfig() global.RedisConfig {
	filePath := filepath.Join(global.ConfigsDir, "redisconfig.xml")
	var config global.RedisConfig
	if err := readXMLFromFile(filePath, &config); err != nil {
		log.Println("Error loading Redis config: ", err)
	}
	return config
}

// 加载emailconfig.xml并返回
func InitEmailConfig() global.MailConfig {
	filePath := filepath.Join(global.ConfigsDir, "emailconfig.xml")
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		InputMailConfig()
	}
	var mc global.MailConfig
	if err := readXMLFromFile(filePath, &mc); err != nil {
		log.Println("Error loading email config: ", err)
	}
	return mc
}

// 初始化邮箱配置
func InputMailConfig() {
	var config global.MailConfig
	log.Print("[FeasOJ]Host of your email server(like smtp.qq.com): ")
	fmt.Scanln(&config.Host)
	log.Print("[FeasOJ]Email address: ")
	fmt.Scanln(&config.User)
	log.Print("[FeasOJ]Email Password: ")
	fmt.Scanln(&config.Pass)
	config.Port = 465

	filePath := filepath.Join(global.ConfigsDir, "emailconfig.xml")
	if err := writeXMLToFile(filePath, config); err != nil {
		log.Println("Error saving email config:", err)
	}
}
