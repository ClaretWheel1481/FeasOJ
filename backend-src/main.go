package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 用户表：uid, avartar, username, password, email, score, synopsis, submit_history
type User struct {
	Uid           int    `gorm:"primaryKey"`
	Avartar       string `gorm:"not null"`
	Username      string `gorm:"not null"`
	Password      string `gorm:"not null"`
	Email         string `gorm:"not null"`
	Score         int    `gorm:"not null"`
	Synopsis      string `gorm:"not null"`
	SubmitHistory string `gorm:"not null"`
}

// 连接数据库、创建表
func initSql(db *gorm.DB) {
	fmt.Println("初始化数据库中...")
	db.AutoMigrate(&User{})
	fmt.Println("初始化数据库成功。")
}

func main() {
	// r := gin.Default()
	fmt.Println("连接数据库中...")
	db, err := gorm.Open(mysql.Open("FeasOJ:YE8cimwfC5KxpN6N@tcp(cloud.claret.space:9936)/FeasOJ?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		fmt.Println("数据库连接失败。")
	}
	fmt.Println("数据库连接成功。")
	initSql(db)
	// r.Run() // listen and serve on 0.0.0.0:8080
}
