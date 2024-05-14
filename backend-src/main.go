package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

//消息队列RabbitMQ

func main() {
	if initSql() {
		fmt.Println("[FeasOJ]数据库初始化成功！")
	} else {
		fmt.Println("[FeasOJ]数据库初始化失败，请确认数据库连接是否正确！")
		return
	}
	// 启动服务器
	gin.SetMode(gin.ReleaseMode)
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
	fmt.Println("若要修改数据库连接信息，请修改config.xml文件.")
	r.Run("0.0.0.0:37881")
}
