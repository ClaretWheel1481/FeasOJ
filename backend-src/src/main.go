package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// TODO:消息队列待实现

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

var parentDir string
var configsDir string

func main() {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("[FeasOJ]获取目录失败，即将退出。")
		return
	}
	parentDir = filepath.Dir(currentDir)
	configsDir = filepath.Join(parentDir, "configs")
	// 如果没找到configs，则创建configs文件夹
	if _, err := os.Stat(configsDir); os.IsNotExist(err) {
		os.Mkdir(configsDir, 0755)
	}

	if initSql() {
		fmt.Println("[FeasOJ]MySQL初始化完毕！")
	} else {
		fmt.Println("[FeasOJ]MySQL初始化失败，请确认数据库连接是否正确！")
		return
	}
	rdb := initRedis()
	fmt.Println("[FeasOJ]Redis连接信息为:", rdb)
	mcfg := initEmailConfig()
	fmt.Println("[FeasOJ]邮箱配置信息为:", mcfg)
	// 创建存放头像文件夹
	initAvatarFolder()
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
			userPHash := selectPassword(Username)
			if userPHash == "" {
				c.JSON(401, gin.H{"status:": 401, "message": "用户不存在"})
			} else {
				// 校验密码是否正确
				if verifyPassword(Password, userPHash) {
					// 生成Token并返回至前端
					token := GenerateToken(Username)
					c.JSON(200, gin.H{"status:": 200, "message": "登录成功", "token": token})
				} else {
					c.JSON(401, gin.H{"status:": 401, "message": "密码错误"})
				}
			}
		})

		// 获取验证码API
		router.GET("/getCaptcha", func(c *gin.Context) {
			// 获取邮箱地址
			email := c.Query("email")
			if sendVerifycode(initEmailConfig(), email, generateVerifycode()) == "Success" {
				c.JSON(200, gin.H{"status:": 200, "message": "验证码发送成功"})
			} else {
				c.JSON(400, gin.H{"status:": 400, "error": "验证码发送失败"})
			}
		})

		// 校验验证码API
		router.GET("/verifyCaptcha", func(c *gin.Context) {
			// 获取邮箱地址和验证码
			email := c.Query("email")
			vcode := c.Query("vcode")
			// 校验验证码是否正确
			if compareVerifyCode(vcode, email) {
				c.JSON(200, gin.H{"status:": 200, "message": "验证码校验成功"})
			} else {
				c.JSON(400, gin.H{"status:": 400, "error": "验证码校验失败"})
			}
		})

		// 注册API
		router.GET("/register", func(c *gin.Context) {
			email := c.Query("email")
			username := c.Query("username")
			password := c.Query("password")
			// 判断用户或邮箱是否存在
			if isUserExist(username, email) {
				c.JSON(400, gin.H{"status:": 400, "error": "用户已存在或邮箱已使用"})
				return
			}
			tokensecret := uuid.New().String()
			regstatus := register(username, encryptPassword(password), email, tokensecret, 0)
			if regstatus {
				c.JSON(200, gin.H{"status:": 200, "message": "注册成功"})
			} else {
				c.JSON(400, gin.H{"status:": 400, "error": "注册失败"})
			}
		})

		// 获取用户信息API
		router.GET("/getUserInfo", func(c *gin.Context) {
			// 返回至前端以显示个人资料
			username := c.Query("username")
			// 根据uid来查询对应的用户信息
			userInfo := selectUserInfo(username)
			c.JSON(200, gin.H{"status:": 200, "Info": userInfo})
		})

		// 密码修改API
		router.GET("/updatePassword", func(c *gin.Context) {
			// 获取账号以及新密码
			newpassword := c.Query("newpassword")
			// 获取邮箱地址和验证码
			email := c.Query("email")
			vcode := c.Query("vcode")
			if compareVerifyCode(vcode, email) && updatePassword(email, newpassword) {
				c.JSON(200, gin.H{"status:": 200, "message": "密码修改成功"})
			} else {
				c.JSON(400, gin.H{"status:": 400, "error": "密码修改失败"})
			}
		})

		// 更新用户信息（非修改密码）API
		router.GET("/updateUserInfo", func(c *gin.Context) {
			// TODO:更新用户信息功能待实现、等待前端修改
		})

		// 用户上传头像API
		router.POST("/uploadAvatar", func(c *gin.Context) {
			file, err := c.FormFile("avatar")
			username := c.Query("username")
			// 检查上传文件类型为png或jpg
			if !strings.HasSuffix(file.Filename, ".png") && !strings.HasSuffix(file.Filename, ".jpg") {
				c.JSON(400, gin.H{"status:": 400, "error": "头像文件类型错误"})
				return
			}
			if err != nil {
				c.JSON(400, gin.H{"status:": 400, "error": "头像上传失败"})
			}

			// 压缩文件并修改文件名为数据库中用户的uid，例如user_1.png/jpg
			user := fmt.Sprintf("%d", selectUserInfo(username).Uid)
			file.Filename = "user_" + user + file.Filename[strings.LastIndex(file.Filename, "."):]

			// 存放到avatars文件夹中
			err = c.SaveUploadedFile(file, "../avatars/"+file.Filename)
			if err != nil {
				c.JSON(400, gin.H{"status:": 400, "error": "头像上传失败"})
			}

			// 更新数据库中的头像地址
			if updateAvatar(username, avatarsDir+"\\"+file.Filename) {
				c.JSON(200, gin.H{"status:": 200, "message": "头像上传成功"})
			}
		})

		// 获取总提交记录API
		router.GET("/getAllSubmitRecords", func(c *gin.Context) {
			submitrecords := selectAllSubmitRecords()
			c.JSON(200, gin.H{"submitrecords": submitrecords})
		})
	}

	fmt.Println("[FeasOJ]服务器已启动，API地址：http://localhost:37881/api/")
	fmt.Println("[FeasOJ]若要修改数据库连接与邮箱配置信息，请修改目录下对应的.xml文件。")
	r.Run("0.0.0.0:37881")
}
