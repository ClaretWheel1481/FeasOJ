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
	global.CurrentDir, _ = os.Getwd()
	global.ParentDir = filepath.Dir(global.CurrentDir)
	// TODO: 每次编译前需要修改为CurrentDir，debug时用ParentDir
	global.ConfigsDir = filepath.Join(global.CurrentDir, "/configs")
	global.CertDir = filepath.Join(global.CurrentDir, "/certificate")
	// 如果没有找到configs，则创建configs文件夹
	if _, err := os.Stat(global.ConfigsDir); os.IsNotExist(err) {
		os.Mkdir(global.ConfigsDir, os.ModePerm)
	}

	// 如果没有找到certificate，则创建certificate文件夹
	if _, err := os.Stat(global.CertDir); os.IsNotExist(err) {
		os.Mkdir(global.CertDir, os.ModePerm)
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
	// 启动服务器
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// 设置路由组
	router := r.Group("/api/v1")
	{
		// 登录API
		router.GET("/login", ginrouter.Login)

		// 获取验证码API
		router.GET("/getCaptcha", ginrouter.GetCaptcha)

		// 获取用户信息API
		router.GET("/getUserInfo", ginrouter.GetUserInfo)

		// TODO: 更新用户信息（非修改密码）API
		router.GET("/updateUserInfo")

		// 获取所有题目API
		router.GET("/getAllProblems", ginrouter.GetAllProblems)

		// 根据题目ID获取题目信息
		router.GET("/getProblemInfo/:id", ginrouter.GetProblemInfo)

		// 获取总提交记录API
		router.GET("/getAllSubmitRecords", ginrouter.GetAllSubmitRecords)

		// 获取指定用户的提交记录API
		router.GET("/getSubmitRecordsByUid/:uid", ginrouter.GetSubmitRecordsByUid)

		// 获取所有讨论帖子API
		router.GET("/getAllDiscussions", ginrouter.GetAllDiscussions)

		// 获取指定Did的帖子API
		router.GET("/getDiscussionByDid/:Did", ginrouter.GetDiscussionByDid)

		// 获取指定帖子的讨论API
		router.GET("/getComment/:Did", ginrouter.GetComment)

		// 管理员获取指定题目的所有信息API
		router.GET("/getProblemAllInfo/:Pid", ginrouter.GetProblemAllInfo)

		// 管理员获取所有用户信息API
		router.GET("/getAllUserInfo", ginrouter.GetAllUsersInfo)
	}

	router2 := r.Group("/api/v2")
	{
		// 用户上传头像API
		router2.POST("/uploadAvatar", ginrouter.UploadAvatar)

		// 注册API
		router2.POST("/register", ginrouter.Register)

		// 密码修改API
		router2.POST("/updatePassword", ginrouter.UpdatePassword)

		// 创建讨论API
		router2.POST("/addDiscussion", ginrouter.CreateDiscussion)

		// 删除讨论API
		router2.POST("/deleteDiscussion/:Did", ginrouter.DeleteDiscussion)

		// 添加评论API
		router2.POST("/addComment/:Did", ginrouter.AddComment)

		// 删除评论API
		router2.POST("/delComment/:Cid", ginrouter.DelComment)

		// 上传代码文件API
		router2.POST("/uploadCode", ginrouter.UploadCode)

		// 管理员更新题目信息API
		router2.POST("/updateProblemInfo", ginrouter.UpdateProblemInfo)

		// 管理员删除题目API
		router2.POST("/delProblemAllInfo/:Pid", ginrouter.DeleteProblem)

		// 管理员晋升用户API
		router2.POST("/promoteUser", ginrouter.PromoteUser)

		// 管理员降级用户API
		router2.POST("/demoteUser", ginrouter.DemoteUser)

		// 管理员封禁用户API
		router2.POST("/banUser", ginrouter.BanUser)

		// 管理员解封用户API
		router2.POST("/unbanUser", ginrouter.UnbanUser)
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

	// TODO: 注意注意！！
	// HTTP服务器
	go func() {
		if err := r.Run("0.0.0.0:37881"); err != nil {
			fmt.Printf("[FeasOJ]服务器启动错误: %v\n", err)
			return
		}
	}()

	// 启动HTTPS服务器
	// go func() {
	// 	if err := r.RunTLS("0.0.0.0:37881", "./certificate/fullchain.pem", "./certificate/privkey.pem"); err != nil {
	// 		fmt.Printf("[FeasOJ]服务器启动错误: %v\n", err)
	// 		return
	// 	}
	// }()

	// 监听终端输入
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			if scanner.Text() == "quit" {
				fmt.Println("[FeasOJ]正在关闭服务器....")
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
