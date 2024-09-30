package config

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"src/global"
)

// 写入配置到文件
func WriteConfigToFile(filePath string, config global.Config) error {
	configXml, err := xml.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, configXml, 0644)
}

// 读取配置文件
func ReadConfigFromFile(filePath string) (global.Config, error) {
	var config global.Config
	configFile, err := os.Open(filePath)
	if err != nil {
		return config, err
	}
	defer configFile.Close()
	err = xml.NewDecoder(configFile).Decode(&config)
	return config, err
}

// 初始化配置文件
func InitConfig() {
	filePath := filepath.Join(global.ConfigDir, "config.xml")
	// 判断是否有config.xml文件，没有则输入
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		InputSqlInfo()
		InputRedisInfo()
		InputMailConfig()
	}
}

// 初始化并保存Sql配置
func InputSqlInfo() bool {
	var sqlConfig global.SqlConfig
	log.Println("[FeasOJ]Please input the MySQL connection configuration：")
	fmt.Print("[FeasOJ]Database name: ")
	fmt.Scanln(&sqlConfig.DbName)
	fmt.Print("[FeasOJ]Database address with port: ")
	fmt.Scanln(&sqlConfig.DbAddress)
	fmt.Print("[FeasOJ]Database user: ")
	fmt.Scanln(&sqlConfig.DbUser)
	fmt.Print("[FeasOJ]Database password: ")
	fmt.Scanln(&sqlConfig.DbPassword)
	log.Println("[FeasOJ]Saving the connection configuration...")

	filePath := filepath.Join(global.ConfigDir, "config.xml")
	config := global.Config{}
	config.SqlConfig = sqlConfig
	if err := WriteConfigToFile(filePath, config); err != nil {
		log.Println("Error saving SQL config: ", err)
		return false
	}
	return true
}

// 加载Sql配置
func LoadSqlConfig() string {
	filePath := filepath.Join(global.ConfigDir, "config.xml")
	config, err := ReadConfigFromFile(filePath)
	if err != nil {
		log.Println("Error loading SQL config: ", err)
		return ""
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Asia%%2FShanghai",
		config.SqlConfig.DbUser, config.SqlConfig.DbPassword, config.SqlConfig.DbAddress, config.SqlConfig.DbName)
	return dsn
}

// 初始化并保存Redis配置
func InputRedisInfo() bool {
	var redisConfig global.RedisConfig
	log.Println("[FeasOJ]Please input the Redis connection configuration：")
	fmt.Print("[FeasOJ]Redis address with port: ")
	fmt.Scanln(&redisConfig.Address)
	fmt.Print("[FeasOJ]Redis password: ")
	fmt.Scanln(&redisConfig.Password)

	filePath := filepath.Join(global.ConfigDir, "config.xml")
	config := global.Config{}
	config.RedisConfig = redisConfig
	if err := WriteConfigToFile(filePath, config); err != nil {
		log.Println("Error saving Redis config: ", err)
		return false
	}
	return true
}

// 加载Redis配置
func LoadRedisConfig() global.RedisConfig {
	filePath := filepath.Join(global.ConfigDir, "config.xml")
	config, err := ReadConfigFromFile(filePath)
	if err != nil {
		log.Println("Error loading Redis config: ", err)
	}
	return config.RedisConfig
}

// 初始化并保存Email配置
func InputMailConfig() {
	var mailConfig global.MailConfig
	fmt.Print("[FeasOJ]Host of email server(e.g. smtp.qq.com): ")
	fmt.Scanln(&mailConfig.Host)
	fmt.Print("[FeasOJ]Email address: ")
	fmt.Scanln(&mailConfig.User)
	fmt.Print("[FeasOJ]Email Password: ")
	fmt.Scanln(&mailConfig.Password)
	mailConfig.Port = 465

	filePath := filepath.Join(global.ConfigDir, "config.xml")
	config := global.Config{}
	config.MailConfig = mailConfig
	if err := WriteConfigToFile(filePath, config); err != nil {
		log.Println("Error saving email config:", err)
	}
}

// 加载Email配置
func InitEmailConfig() global.MailConfig {
	filePath := filepath.Join(global.ConfigDir, "config.xml")
	config, err := ReadConfigFromFile(filePath)
	if err != nil {
		log.Println("Error loading Email config: ", err)
	}
	return config.MailConfig
}
