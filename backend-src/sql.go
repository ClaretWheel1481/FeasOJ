package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 连接数据库、创建表
func initSql() {
	fmt.Println("[FeasOJ]连接数据库中...")
	db, err := gorm.Open(mysql.Open("FeasOJ:YE8cimwfC5KxpN6N@tcp(cloud.claret.space:9936)/FeasOJ?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		fmt.Println("[FeasOJ]数据库连接失败。")
	}
	fmt.Println("[FeasOJ]数据库连接成功。")
	fmt.Println("[FeasOJ]创建数据表中...")
	db.AutoMigrate(&User{}, &Problem{})
	fmt.Println("[FeasOJ]创建数据表成功。")
}
