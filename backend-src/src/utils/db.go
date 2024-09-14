package utils

import (
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"src/global"
	"time"

	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

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

func LoadSqlConfig() string {
	// 读取config.xml文件
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

func InitAdminAccount() (string, string, string, string, int) {
	//创建管理员
	var adminUsername string
	var adminPassword string
	var adminEmail string
	fmt.Println("[FeasOJ]创建管理员账户")
	fmt.Print("[FeasOJ]请输入管理员账号:")
	fmt.Scanln(&adminUsername)
	fmt.Print("[FeasOJ]请输入管理员密码:")
	fmt.Scanln(&adminPassword)
	fmt.Print("[FeasOJ]请输入管理员邮箱:")
	fmt.Scanln(&adminEmail)

	return adminUsername, EncryptPassword(adminPassword), adminEmail, uuid.New().String(), 1
}

// 创建表
func InitSql() bool {
	filePath := filepath.Join(global.ConfigsDir, "sqlconfig.xml")
	//判断是否有config.xml文件，没有则输入
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		InputSqlInfo()
	}
	ConnectSql().AutoMigrate(&global.User{}, &global.Problem{}, &global.SubmitRecord{}, &global.Discussion{}, &global.Comment{}, &global.TestCase{}, &global.Competition{})
	return true
}

// 返回数据库连接对象
func ConnectSql() *gorm.DB {
	dsn := LoadSqlConfig()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("[FeasOJ]数据库连接失败，请手动前往config.xml进行配置。")
		return nil
	}
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("[FeasOJ]获取通用数据库对象失败。")
		return nil
	}
	// 设置连接池避免数据库连接过多导致的性能问题
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Second * 6)
	return db
}

// 根据用户名获取用户信息
func SelectUser(username string) global.User {
	var user global.User
	ConnectSql().Where("username = ?", username).First(&user)
	return user
}
