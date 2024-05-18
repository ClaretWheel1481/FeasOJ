package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// TODO:消息队列RabbitMQ待实现

// 用户Token生成后返回给前端
func GenerateToken(username string) string {
	token := jwt.New(jwt.SigningMethodHS256)
	// 设置Token的Claims
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	// 生成Token
	tokenString, err := token.SignedString([]byte(selectTokenSecret(username)))
	if err != nil {
		fmt.Println("生成Token失败：", err)
		return ""
	}
	return tokenString
}

// 校验Token
func VerifyToken(tokenString string) bool {
	// 解析Token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 验证签名方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// 返回签名密钥
		return []byte(selectTokenSecret(token.Claims.(jwt.MapClaims)["username"].(string))), nil

	})
	// 验证Token
	if err != nil {
		fmt.Println("Token验证失败：", err)
		return false
	}
	// 验证Token是否有效
	if !token.Valid {
		fmt.Println("Token无效")
		return false
	}
	return true
}

func main() {
	// TODO:多线程运行待实现
	if initSql() {
		fmt.Println("[FeasOJ]MySQL初始化完毕！")
	} else {
		fmt.Println("[FeasOJ]MySQL初始化失败，请确认数据库连接是否正确！")
		return
	}
	rdb := initRedis()
	fmt.Println("[FeasOJ]Redis连接信息为：", rdb)
	// 启动服务器
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	//TODO:响应前端部分功能待实现
	// 设置路由组
	router := r.Group("/api")
	{
		// 登录API
		router.GET("/login", func(c *gin.Context) {
			Username := c.Query("username")
			Password := c.Query("password")
			userExist := selectPassword(Username)
			if userExist != "" {
				userPHash := selectPassword(Username)
				// 校验密码是否正确
				if verifyPassword(Password, userPHash) {
					// 生成Token并返回至前端
					token := GenerateToken(Username)
					c.JSON(200, gin.H{"status:": 200, "token": token})
				} else {
					c.JSON(401, gin.H{"status:": 401, "error": "密码错误"})
				}
			} else {
				c.JSON(401, gin.H{"status:": 401, "error": "用户不存在"})
			}
		})

		// 注册API
		router.GET("/register", func(c *gin.Context) {
			// TODO:注册功能待实现
		})

		// 获取验证码API
		router.GET("/getCaptcha", func(c *gin.Context) {
			// TODO: 获取验证码后将验证码存入Redis，用邮箱做Key，验证码做Value等候校验
			// verifyCode := generateVerifycode()

		})
	}

	fmt.Println("[FeasOJ]服务器已启动，API地址：http://localhost:37881/api/")
	fmt.Println("[FeasOJ]若要修改数据库连接与邮箱配置信息，请修改目录下对应的.xml文件。")
	// 测试用户Token校验代码
	// fmt.Println(VerifyToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IkFkbWluIn0.x1q-3KM2EDlkn7XUmrQ42p83bOV2EFLWMBEF4IHubCY"))
	// 测试用户验证码Redis校验
	// fmt.Println(sendVerifycode(initEmailConfig(), "claretwheel1482@gmail.com", generateVerifycode()))
	// var fec string
	// fmt.Scanln(&fec)
	// fmt.Println(compareVerifyCode(fec, "claretwheel1482@gmail.com"))
	r.Run("0.0.0.0:37881")
}
