package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"src/account"
	"src/codehandler"
	"src/ginrouter"
	"src/global"
	"src/utils"
	"syscall"

	"github.com/gin-gonic/gin"
)

func main() {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("[FeasOJ]获取目录失败，即将退出。")
		return
	}
	global.ParentDir = filepath.Dir(currentDir)
	// TODO:每次编译前需要修改为currentDir，debug时用parentDir
	global.ConfigsDir = filepath.Join(global.ParentDir, "/configs")
	// 如果没有找到configs，则创建configs文件夹
	if _, err := os.Stat(global.ConfigsDir); os.IsNotExist(err) {
		os.Mkdir(global.ConfigsDir, os.ModePerm)
	}
	// 创建存放头像文件夹
	account.InitAvatarFolder()
	// 创建存放代码文件夹
	codehandler.InitCodeFolder()
	// 初始化数据库连接配置
	if utils.InitSql() {
		fmt.Println("[FeasOJ]MySQL初始化完毕！")
	} else {
		fmt.Println("[FeasOJ]MySQL初始化失败，请确认数据库连接是否正确！")
		return
	}
	utils.InitRedis()
	utils.InitEmailConfig()
	// if codehandler.StartContainer() {
	// 	fmt.Println("[FeasOJ]SandBox启动成功！")
	// } else {
	// 	fmt.Println("[FeasOJ]SandBox启动失败！")
	// 	return
	// }
	// 启动服务器
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	//TODO:响应前端部分功能待实现
	// 设置路由组
	router := r.Group("/api/v1")
	{
		// 登录API
		router.GET("/login", ginrouter.Logins)

		// 获取验证码API
		router.GET("/getCaptcha", ginrouter.GetCaptchas)

		// 校验TokenAPI
		router.GET("/verifyToken", ginrouter.VerifyTokens)

		// 获取用户信息API
		router.GET("/getUserInfo", ginrouter.GetUserInfos)

		// 更新用户信息（非修改密码）API
		// TODO:更新用户信息功能待实现、等待前端修改
		router.GET("/updateUserInfo")

		// 从消息队列中获取代码运行状态API
		// TODO:等待获取代码运行状态功能实现、消息队列实现
		router.GET("/getCodeStatus")

		// 获取所有题目API
		router.GET("/getAllProblems", ginrouter.GetAllProblemss)

		// 根据题目ID获取题目信息
		router.GET("/getProblemInfo/:id", ginrouter.GetProblemInfos)

		// 获取总提交记录API
		router.GET("/getAllSubmitRecords", ginrouter.GetAllSubmitRecordss)

		// 获取指定用户的提交记录API
		router.GET("/getSubmitRecordsByUid/:uid", ginrouter.GetSubmitRecordsByUids)

		// 获取所有讨论帖子API
		router.GET("/getAllDiscussions", ginrouter.GetAllDiscussionss)

		// 获取指定Did的帖子API
		router.GET("/getDiscussionByDid/:Did", ginrouter.GetDiscussionByDids)

		// 获取指定帖子的讨论API
		router.GET("/getComment/:Did", ginrouter.GetComments)

		// 管理员获取指定题目的所有信息API
		router.GET("/getProblemAllInfo/:Pid", ginrouter.GetProblemAllInfos)

		// 管理员获取所有用户信息API
		router.GET("/getAllUserInfo", ginrouter.GetAllUsersInfos)
	}

	router2 := r.Group("/api/v2")
	{
		// 用户上传头像API
		router2.POST("/uploadAvatar", ginrouter.UploadAvatars)

		// 注册API
		router2.POST("/register", ginrouter.Registers)

		// 密码修改API
		router2.POST("/updatePassword", ginrouter.UpdatePasswords)

		// 创建讨论API
		router2.POST("/addDiscussion", ginrouter.CreateDiscussion)

		// 删除讨论API
		router2.POST("/deleteDiscussion/:Did", ginrouter.DeleteDiscussion)

		// 添加评论API
		router2.POST("/addComment/:Did", ginrouter.AddComments)

		// 删除评论API
		router2.POST("/delComment/:Cid", ginrouter.DelComments)

		// 上传代码文件API
		router2.POST("/uploadCode", ginrouter.UploadCodes)

		// 管理员更新题目信息API
		router2.POST("/updateProblemInfo", ginrouter.UpdateProblemInfos)

		// 管理员删除题目API
		router2.POST("/delProblemAllInfo/:Pid", ginrouter.DeleteProblems)
	}

	// 头像http服务器
	r.StaticFS("/avatar", http.Dir(global.AvatarsDir))
	if codehandler.BuildImage() {
		fmt.Println("[FeasOJ]SandBox构建成功！")
	} else {
		fmt.Println("[FeasOJ]SandBox构建失败！")
		return
	}
	fmt.Println("[FeasOJ]服务已启动。")
	fmt.Println("[FeasOJ]若要修改数据库连接与邮箱配置信息，请修改目录下对应的.xml文件。")

	// 实时检测Redis JudgeTask中是否有任务
	rdb := utils.InitRedis()
	go codehandler.ProcessJudgeTasks(rdb)

	// 启动服务器
	go func() {
		if err := r.Run("0.0.0.0:37881"); err != nil {
			fmt.Printf("Server exited with error: %v\n", err)
		}
	}()

	// 监听终端输入
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			if scanner.Text() == "quit" {
				fmt.Println("[FeasOJ]正在关闭服务器....")
				// codehandler.TerminateContainer(global.ContainerID)
				os.Exit(0)
			}
		}
	}()

	// 等待中断信号以优雅地关闭服务器
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("[FeasOJ]输入“quit”可停止服务器并删除容器。")
	<-quit
}
