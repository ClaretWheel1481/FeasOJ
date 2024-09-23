package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"src/codehandler"
	gincontext "src/gin"
	"src/global"
	"src/middleware"
	"src/utils"
	"src/utils/sql"
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
	// 初始化数据库连接配置
	if utils.InitSql() {
		fmt.Println("[FeasOJ]MySQL初始化完毕！")
	} else {
		fmt.Println("[FeasOJ]MySQL初始化失败，请确认数据库连接是否正确！")
		return
	}

	// 初始化管理员账户
	if sql.SelectAdminUser(1) {
		fmt.Println("[FeasOJ]管理员账号已存在！")
	} else {
		sql.Register(utils.InitAdminAccount())
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
	router1 := r.Group("/api/v1")
	router1.Use(middleware.HeaderVerify())
	{
		// 验证用户信息API
		router1.GET("/verify", gincontext.VerifyUserInfo)

		// 获取指定帖子的评论API
		router1.GET("/discussions/:Did/comments", gincontext.GetComment)

		// 获取竞赛列表
		router1.GET("/competitions", gincontext.GetCompetitionList)

		// 获取所有题目API
		router1.GET("/problems", gincontext.GetAllProblems)

		// 获取所有讨论帖子API
		router1.GET("/discussions", gincontext.GetAllDiscussions)

		// 根据题目ID获取题目信息
		router1.GET("/problems/:id", gincontext.GetProblemInfo)

		// 获取总提交记录API
		router1.GET("/submitrecords", gincontext.GetAllSubmitRecords)

		// 获取指定Did的帖子API
		router1.GET("/discussions/:Did", gincontext.GetDiscussionByDid)

		// 上传代码API
		router1.POST("/users/code", gincontext.UploadCode)

		// 创建讨论API
		router1.POST("/discussions", gincontext.CreateDiscussion)

		// 添加评论API
		router1.POST("/discussions/:Did/comments", gincontext.AddComment)

		// 用户上传头像API
		router1.PUT("/users/avatar", gincontext.UploadAvatar)

		// 简介更新API
		router1.PUT("/users/synopsis", gincontext.UpdateSynopsis)

		// 管理员晋升用户API
		router1.PUT("/admin/users/promote", gincontext.PromoteUser)

		// 管理员降级用户API
		router1.PUT("/admin/users/demote", gincontext.DemoteUser)

		// 管理员封禁用户API
		router1.PUT("/admin/users/ban", gincontext.BanUser)

		// 管理员解封用户API
		router1.PUT("/admin/users/unban", gincontext.UnbanUser)

		// 管理员获取竞赛列表
		router1.GET("/admin/competitions", gincontext.GetCompetitionListAdmin)

		// 管理员获取指定竞赛ID信息
		router1.GET("/admin/competitions/:cid", gincontext.GetCompetitionInfoAdmin)

		// 管理员获取指定题目的所有信息API
		router1.GET("/admin/problems/:Pid", gincontext.GetProblemAllInfo)

		// 管理员获取所有用户信息API
		router1.GET("/admin/users", gincontext.GetAllUsersInfo)

		// 管理员新增/更新题目信息API
		router1.POST("/admin/problems", gincontext.UpdateProblemInfo)

		// 管理员删除题目API
		router1.DELETE("/admin/problems/:Pid", gincontext.DeleteProblem)

		// 删除讨论API
		router1.DELETE("/discussions/:Did", gincontext.DeleteDiscussion)

		// 删除评论API
		router1.DELETE("/comments/:Cid", gincontext.DelComment)
	}

	router1 = r.Group("/api/v1")
	{
		// 注册API
		router1.POST("/register", gincontext.Register)

		// 登录API
		router1.GET("/login", gincontext.Login)

		// 获取验证码API
		router1.GET("/captcha", gincontext.GetCaptcha)

		// 获取用户信息API
		router1.GET("/users/:username", gincontext.GetUserInfo)

		// 获取指定用户的提交记录API
		router1.GET("/users/:username/submitrecords", gincontext.GetSubmitRecordsByUsername)

		// 密码修改API
		router1.PUT("/users/password", gincontext.UpdatePassword)
	}

	// 挂载头像文件夹
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
