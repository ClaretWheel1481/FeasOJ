package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
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

// 题目表: pid, title, content, time_limit, memory_limit, input, output, contest, submit_history
// contest为0则不属于任何比赛
type Problem struct {
	Pid            int    `gorm:"primaryKey"`
	Title          string `gorm:"not null"`
	Content        string `gorm:"not null"`
	Time_limit     int    `gorm:"not null"`
	Memory_limit   int    `gorm:"not null"`
	Input          string `gorm:"not null"`
	Output         string `gorm:"not null"`
	Contest        int    `gorm:"not null"`
	Submit_history string `gorm:"not null"`
}

// 排行榜表：rank, uid, username, score
// 按照User表的score进行排序
// 每10分钟更新一次，将User表的score更新到Rank表中
type Rank struct {
	Rank     int    `gorm:"primaryKey"`
	Uid      int    `gorm:"not null"`
	Username string `gorm:"not null"`
	Score    int    `gorm:"not null"`
}

// 连接数据库、创建表
func initSql(db *gorm.DB) {
	fmt.Println("[FeasOJ]初始化数据库中...")
	db.AutoMigrate(&User{}, &Problem{}, &Rank{})
	fmt.Println("[FeasOJ]初始化数据库成功。")
}

func main() {
	fmt.Println("[FeasOJ]连接数据库中...")
	db, err := gorm.Open(mysql.Open("FeasOJ:YE8cimwfC5KxpN6N@tcp(cloud.claret.space:9936)/FeasOJ?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		fmt.Println("[FeasOJ]数据库连接失败。")
	}
	fmt.Println("[FeasOJ]数据库连接成功。")
	initSql(db)

	// 启动服务器
	r := gin.Default()
	fmt.Println("[FeasOJ]启动服务器中...")
	// 设置路由
	r.GET("/", func(c *gin.Context) {

	})
	fmt.Println("服务器启动成功，端口号为37881。")
	r.Run("0.0.0.0:37881")
}
