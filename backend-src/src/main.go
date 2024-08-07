package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
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

	// 定义目录映射
	dirs := map[string]*string{
		"configs":     &global.ConfigsDir,
		"certificate": &global.CertDir,
		"avatars":     &global.AvatarsDir,
		"codefiles":   &global.CodeDir,
	}

	// 遍历map，设置路径并创建不存在的目录
	for name, dir := range dirs {
		// TODO: 每次编译前需要修改为CurrentDir，debug时用ParentDir
		*dir = filepath.Join(global.ParentDir, name)
		if _, err := os.Stat(*dir); os.IsNotExist(err) {
			os.Mkdir(*dir, os.ModePerm)
		}
	}
	// 初始化连接配置
	if utils.InitSql() {
		fmt.Println("[FeasOJ]MySQL初始化完毕！")
	} else {
		fmt.Println("[FeasOJ]MySQL初始化失败，请确认数据库连接是否正确！")
		return
	}

	// 构建沙盒镜像
	if codehandler.BuildImage() {
		fmt.Println("[FeasOJ]SandBox构建成功！")
	} else {
		fmt.Println("[FeasOJ]SandBox构建失败！")
		return
	}

	// 启动服务器
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// 设置路由组
	router := r.Group("/api/v1")
	{
		// 登录API
		router.GET("/login", ginrouter.Login)

		// 获取验证码API
		router.GET("/captcha", ginrouter.GetCaptcha)

		// 验证用户信息API
		router.GET("/verify", ginrouter.VerifyUserInfo)

		// 获取用户信息API
		router.GET("/users/:username", ginrouter.GetUserInfo)

		// 获取所有题目API
		router.GET("/problems", ginrouter.GetAllProblems)

		// 根据题目ID获取题目信息
		router.GET("/problems/:id", ginrouter.GetProblemInfo)

		// 获取总提交记录API
		router.GET("/submitrecords", ginrouter.GetAllSubmitRecords)

		// 获取指定用户的提交记录API
		router.GET("/users/:username/submitrecords", ginrouter.GetSubmitRecordsByUsername)

		// 获取所有讨论帖子API
		router.GET("/discussions", ginrouter.GetAllDiscussions)

		// 获取指定Did的帖子API
		router.GET("/discussions/:Did", ginrouter.GetDiscussionByDid)

		// 获取指定帖子的评论API
		router.GET("/discussions/:Did/comments", ginrouter.GetComment)

		// 管理员获取指定题目的所有信息API
		router.GET("/admin/problems/:Pid", ginrouter.GetProblemAllInfo)

		// 管理员获取所有用户信息API
		router.GET("/admin/users", ginrouter.GetAllUsersInfo)

		// 上传代码API
		router.POST("/users/:username/code", ginrouter.UploadCode)

		// 管理员新增/更新题目信息API
		router.POST("/admin/problems", ginrouter.UpdateProblemInfo)

		// 注册API
		router.POST("/register", ginrouter.Register)

		// 创建讨论API
		router.POST("/discussions", ginrouter.CreateDiscussion)

		// 添加评论API
		router.POST("/discussions/:Did/comments", ginrouter.AddComment)

		// 用户修改头像API
		router.PUT("/users/:username/avatar", ginrouter.UploadAvatar)

		// 密码修改API
		router.PUT("/users/password", ginrouter.UpdatePassword)

		// 简介更新API
		router.PUT("/users/:username/synopsis", ginrouter.UpdateSynopsis)

		// 管理员晋升用户API
		router.PUT("/admin/users/promote", ginrouter.PromoteUser)

		// 管理员降级用户API
		router.PUT("/admin/users/demote", ginrouter.DemoteUser)

		// 管理员封禁用户API
		router.PUT("/admin/users/ban", ginrouter.BanUser)

		// 管理员解封用户API
		router.PUT("/admin/users/unban", ginrouter.UnbanUser)

		// 删除讨论API
		router.DELETE("/discussions/:Did", ginrouter.DeleteDiscussion)

		// 删除评论API
		router.DELETE("/comments/:Cid", ginrouter.DelComment)

		// 管理员删除题目API
		router.DELETE("/admin/problems/:Pid", ginrouter.DeleteProblem)
	}

	// 头像http服务器
	r.StaticFS("/avatar", http.Dir(global.AvatarsDir))

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
