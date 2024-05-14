package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

//消息队列RabbitMQ

// 用户表：uid, avartar, username, password, email, score, synopsis, submit_history, create_at
type User struct {
	Uid           int       `gorm:"primaryKey;autoIncrement"`
	Avartar       string    `gorm:"not null"`
	Username      string    `gorm:"not null"`
	Password      string    `gorm:"not null"`
	Email         string    `gorm:"not null"`
	Score         int       `gorm:"not null"`
	Synopsis      string    `gorm:"not null"`
	SubmitHistory string    `gorm:"not null"`
	CreateAt      time.Time `gorm:"not null"`
}

// 题目表: pid, title, content, time_limit, memory_limit, input, output, contest, submit_history
// contest为0则不属于任何比赛
type Problem struct {
	Pid            int    `gorm:"primaryKey;autoIncrement"`
	Title          string `gorm:"not null"`
	Content        string `gorm:"not null"`
	Time_limit     int    `gorm:"not null"`
	Memory_limit   int    `gorm:"not null"`
	Input          string `gorm:"not null"`
	Output         string `gorm:"not null"`
	Contest        int    `gorm:"not null"`
	Submit_history string `gorm:"not null"`
}

func main() {
	initSql()
	// 启动服务器
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	fmt.Println("[FeasOJ]启动服务器中...")
	//TODO:响应前端部分功能待实现
	// 设置路由组
	router := r.Group("/api")
	{
		// 登录API
		router.POST("/login", func(c *gin.Context) {
			var loginInfo struct {
				Username string `json:"username"`
				Password string `json:"password"`
			}
			if err := c.ShouldBindJSON(&loginInfo); err != nil {
				c.JSON(400, gin.H{"error": "请求格式错误。"})
				return
			}
			// TODO:登录成功则生成Token返回至前端
		})

		// 注册API
		router.POST("/register", func(c *gin.Context) {
			// TODO：注册功能待实现
		})
	}

	fmt.Println("服务器启动成功，API地址：http://localhost:37881/api/")

	r.Run("0.0.0.0:37881")
}
