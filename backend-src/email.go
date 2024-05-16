package main

//TODO:发件服务端配置
// 要求：生成随机4位验证码并发送、校验验证码、密码修改后发送邮件
import (
	"fmt"
	"math/rand"
	"time"
)

type mailConfig struct {
	Host string
	Port int
	User string
	Pass string
}

// 初始化邮箱配置并作为结构体返回给其他方法
func initMailConfig() *mailConfig {
	var hosts string
	var users string
	var password string
	fmt.Print("[FeasOJ]请输入Smtp服务器地址:")
	fmt.Scanln(&hosts)
	fmt.Print("[FeasOJ]请输入发件人邮箱地址（如114514@qq.com）:")
	fmt.Scanln(&users)
	fmt.Print("[FeasOJ]请输入邮箱密码（不一定是登录密码）:")
	fmt.Scanln(&password)
	return &mailConfig{
		Host: hosts,
		Port: 465,
		User: users,
		Pass: password,
	}
}

// 随机生成4位数字验证码
func generateVerifycode() string {
	return fmt.Sprintf("%04d", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))
}

// 发送验证码
func sendVerifycode(config *mailConfig, to string, verifycode string) error {
	// TODO:发送验证码待实现
	return nil
}
