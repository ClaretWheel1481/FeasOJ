package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"src/config"
	"src/global"
	"time"

	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 创建管理员
func InitAdminAccount() (string, string, string, string, int) {
	var adminUsername string
	var adminPassword string
	var adminEmail string
	log.Println("[FeasOJ]Please input the administrator account configuration: ")
	log.Print("[FeasOJ]Admin account: ")
	fmt.Scanln(&adminUsername)
	log.Print("[FeasOJ]Admin password: ")
	fmt.Scanln(&adminPassword)
	log.Print("[FeasOJ]Admin email: ")
	fmt.Scanln(&adminEmail)

	return adminUsername, EncryptPassword(adminPassword), adminEmail, uuid.New().String(), 1
}

// 创建表
func InitSql() bool {
	filePath := filepath.Join(global.ConfigsDir, "sqlconfig.xml")
	//判断是否有config.xml文件，没有则输入
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		config.InputSqlInfo()
	}
	ConnectSql().AutoMigrate(&global.User{}, &global.Problem{}, &global.SubmitRecord{}, &global.Discussion{}, &global.Comment{}, &global.TestCase{}, &global.Competition{})
	return true
}

// 返回数据库连接对象
func ConnectSql() *gorm.DB {
	dsn := config.LoadSqlConfig()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("[FeasOJ]Database connection failed, please go to sqlconfig.xml manually to configure.")
		return nil
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Println("[FeasOJ]Failed to get generic database object.")
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
