package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// TODO:消息队列RabbitMQ待实现

// TODO:用户Token生成后返回给前端
func GenerateToken(username string) string {
	token := jwt.New(jwt.SigningMethodHS256)
	// 设置Token的Claims
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	// 设置Token过期时间为30天
	claims["exp"] = time.Now().Add(time.Hour * 24 * 30).Unix()
	// 生成Token
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		fmt.Println("生成Token失败：", err)
		return ""
	}
	return tokenString
}

// TODO:校验Token，返回结果给前端，若不正确，前端退出登录
func VerifyToken(tokenString string) string {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 验证签名
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// 返回用于验证签名的密钥
		return []byte("secret"), nil
	})
	if err != nil {
		return ""
	}
	// 验证Token的有效性
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// 获取用户名
		username := claims["username"].(string)
		return username
	}
	return ""
}

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
	fmt.Println("若要修改数据库连接信息，请修改config.xml文件。")
	r.Run("0.0.0.0:37881")
}
