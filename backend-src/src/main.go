package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// TODO:消息队列待实现

// 全局变量 - 本地配置文件路径
var parentDir string
var configsDir string
var avatarsDir string

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
	router := r.Group("/api/v1")
	{
		// 登录API
		router.GET("/login", func(c *gin.Context) {
			Username := c.Query("username")
			Password := c.Query("password")
			userPHash := selectPassword(Username)
			if userPHash == "" {
				c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "用户不存在。"})
			} else {
				// 校验密码是否正确
				if verifyPassword(Password, userPHash) {
					// 生成Token并返回至前端
					token := GenerateToken(Username)
					c.JSON(http.StatusOK, gin.H{"status": 200, "message": "登录成功。", "token": token})
				} else {
					c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "密码错误。"})
				}
			}
		})

		// 获取验证码API
		router.GET("/getCaptcha", func(c *gin.Context) {
			// 获取邮箱地址
			email := c.Query("email")
			if sendVerifycode(initEmailConfig(), email, generateVerifycode()) == "Success" {
				c.JSON(http.StatusOK, gin.H{"status": 200, "message": "验证码发送成功。"})
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "验证码发送失败，可能是我们的问题。"})
			}
		})

		// 校验TokenAPI
		router.GET("/verifyToken", func(c *gin.Context) {
			token := c.GetHeader("token")
			username := c.Query("username")
			if VerifyToken(username, token) {
				c.JSON(http.StatusOK, gin.H{"status": 200, "message": "Token验证成功。"})
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "Token验证失败。"})
			}
		})

		// 获取用户信息API
		router.GET("/getUserInfo", func(c *gin.Context) {
			// 返回至前端以显示个人资料
			username := c.Query("username")
			// 根据uid来查询对应的用户信息
			userInfo := selectUserInfo(username)
			if userInfo.Username == "" {
				c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "用户不存在。"})
			} else {
				c.JSON(http.StatusOK, gin.H{"status": 200, "Info": userInfo})
			}

		})

		// 更新用户信息（非修改密码）API
		router.GET("/updateUserInfo", func(c *gin.Context) {
			// TODO:更新用户信息功能待实现、等待前端修改
		})

		// 从消息队列中获取代码运行状态API
		router.GET("/getCodeStatus", func(c *gin.Context) {
			// TODO:等待获取代码运行状态功能实现、消息队列实现
		})

		// 获取所有题目API
		router.GET("/getAllProblems", func(c *gin.Context) {
			problems := selectAllProblems()
			c.JSON(http.StatusOK, gin.H{"problems": problems})
		})

		// 根据题目ID获取题目信息
		router.GET("/getProblemInfo/:id", func(c *gin.Context) {
			problemID := c.Param("id")
			problemInfo := selectProblemInfo(problemID)
			c.JSON(http.StatusOK, gin.H{"problemInfo": problemInfo})
		})

		// 获取所有竞赛API
		router.GET("/getAllContests", func(c *gin.Context) {
			allContest := selectAllContests()
			c.JSON(http.StatusOK, gin.H{"contests": allContest})
		})

		// 根据竞赛ID获取题目信息
		router.GET("/getContestInfo/:cid", func(c *gin.Context) {
			contestID := c.Param("cid")
			contestInfo := selectProblemByCid(contestID)
			c.JSON(http.StatusOK, gin.H{"contestInfo": contestInfo})
		})

		// 获取总提交记录API
		router.GET("/getAllSubmitRecords", func(c *gin.Context) {
			submitrecords := selectAllSubmitRecords()
			c.JSON(http.StatusOK, gin.H{"submitrecords": submitrecords})
		})
	}

	router2 := r.Group("/api/v2")
	{
		// 用户上传头像API
		router2.POST("/uploadAvatar", func(c *gin.Context) {
			file, err := c.FormFile("avatar")
			username := c.Query("username")
			// 检查上传文件类型为png或jpg
			if !strings.HasSuffix(file.Filename, ".png") && !strings.HasSuffix(file.Filename, ".jpg") && !strings.HasSuffix(file.Filename, ".JPG") && !strings.HasSuffix(file.Filename, ".PNG") {
				c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "头像文件类型错误。"})
				return
			}
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "头像上传失败。"})
			}

			// 压缩文件并修改文件名为数据库中用户的uid，例如user_1.png/jpg
			user := fmt.Sprintf("%d", selectUserInfo(username).Uid)
			file.Filename = "user_" + user + file.Filename[strings.LastIndex(file.Filename, "."):]
			// 检查是否有重复名称，重复则删除
			filepath := "../avatars/" + file.Filename
			if _, err := os.Stat(filepath); err == nil {
				os.Remove(filepath)
			}
			// 存放到avatars文件夹中
			err = c.SaveUploadedFile(file, "../avatars/"+file.Filename)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "头像上传失败。"})
			}

			// 更新数据库中的头像路径地址
			if updateAvatar(username, "..\\..\\backend-src\\avatars\\"+file.Filename) {
				c.JSON(http.StatusOK, gin.H{"status": 200, "message": "头像上传成功。"})
			}
		})

		// 注册API
		router2.POST("/register", func(c *gin.Context) {
			var req RegisterRequest
			c.ShouldBind(&req)
			// 判断用户或邮箱是否存在
			if isUserExist(req.Username, req.Email) {
				c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "用户已存在或邮箱已使用。"})
				return
			}
			tokensecret := uuid.New().String()
			vcodeStatus := compareVerifyCode(req.Vcode, req.Email)
			if !vcodeStatus {
				c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "验证码错误。"})
				return
			}
			regstatus := register(req.Username, encryptPassword(req.Password), req.Email, tokensecret, 0)
			if regstatus {
				c.JSON(http.StatusOK, gin.H{"status": 200, "message": "注册成功，2秒后跳转至登录界面。"})
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "注册失败，请确认邮箱验证码是否正确。"})
			}
		})

		// 密码修改API
		router2.POST("/updatePassword", func(c *gin.Context) {
			var req updatePasswordRequest
			c.ShouldBind(&req)
			vcodeStatus := compareVerifyCode(req.Vcode, req.Email)
			if !vcodeStatus {
				c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "验证码错误。"})
				return
			}
			if updatePassword(req.Email, req.NewPassword) {
				c.JSON(http.StatusOK, gin.H{"status": 200, "message": "密码修改成功，2秒后跳转至登录页面。"})
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "密码修改失败。"})
			}
		})

		// 上传代码文件API
		router.POST("/uploadCode", func(c *gin.Context) {
			// TODO:上传代码文件功能待实现、等待前端实现
		})
	}

	fmt.Println("[FeasOJ]服务器已启动。")
	fmt.Println("[FeasOJ]GETAPI：http://localhost:37881/api/v1/")
	fmt.Println("[FeasOJ]POSTAPI：http://localhost:37881/api/v2/")
	fmt.Println("[FeasOJ]若要修改数据库连接与邮箱配置信息，请修改目录下对应的.xml文件。")
	r.Run("0.0.0.0:37881")
}
