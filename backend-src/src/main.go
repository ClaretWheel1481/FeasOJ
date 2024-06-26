package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

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
	// 初始化数据库连接配置
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
		router.GET("/login", logins)

		// 获取验证码API
		router.GET("/getCaptcha", getCaptchas)

		// 校验TokenAPI
		router.GET("/verifyToken", verifyTokens)

		// 获取用户信息API
		router.GET("/getUserInfo", getUserInfos)

		// 更新用户信息（非修改密码）API
		// TODO:更新用户信息功能待实现、等待前端修改
		router.GET("/updateUserInfo")

		// 从消息队列中获取代码运行状态API
		// TODO:等待获取代码运行状态功能实现、消息队列实现
		router.GET("/getCodeStatus")

		// 获取所有题目API
		router.GET("/getAllProblems", getAllProblemss)

		// 根据题目ID获取题目信息
		router.GET("/getProblemInfo/:id", getProblemInfos)

		// 获取总提交记录API
		router.GET("/getAllSubmitRecords", getAllSubmitRecordss)

		// 获取指定用户的提交记录API
		router.GET("/getSubmitRecordsByUid/:uid", getSubmitRecordsByUids)

		// 获取所有讨论帖子API
		router.GET("/getAllDiscussions", getAllDiscussionss)

		// 获取指定Tid的帖子API
		router.GET("/getDiscussionByTid/:tid", getDiscussionByTids)
	}

	router2 := r.Group("/api/v2")
	{
		// 用户上传头像API
		router2.POST("/uploadAvatar", uploadAvatars)

		// 注册API
		router2.POST("/register", registers)

		// 密码修改API
		router2.POST("/updatePassword", updatePasswords)

		// 创建讨论API
		router2.POST("/addDiscussion", createDiscussion)

		// 删除讨论API
		router2.POST("/deleteDiscussion/:tid", deleteDiscussion)

		// 上传代码文件API
		// TODO:上传代码文件功能待实现、等待前端实现
		router.POST("/uploadCode")
	}

	fmt.Println("[FeasOJ]服务器已启动。")
	fmt.Println("[FeasOJ]GETAPI：http://localhost:37881/api/v1/")
	fmt.Println("[FeasOJ]POSTAPI：http://localhost:37881/api/v2/")
	fmt.Println("[FeasOJ]若要修改数据库连接与邮箱配置信息，请修改目录下对应的.xml文件。")
	r.Run("0.0.0.0:37881")
}
