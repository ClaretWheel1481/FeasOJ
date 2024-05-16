package main

//TODO:发件服务端配置
// 要求：生成随机4位验证码并发送、校验验证码、密码修改后发送邮件
import (
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"math/rand"
	"os"
	"time"

	"gopkg.in/gomail.v2"
)

type mailConfig struct {
	Host string `xml:"host"`
	Port int    `xml:"port"`
	User string `xml:"user"`
	Pass string `xml:"pass"`
}

// 从emailConfig.xml中读取邮箱配置并返回mailConfig
func initEmailConfig() mailConfig {
	// 判断是否有emailconfig.xml文件
	if _, err := os.Stat("emailconfig.xml"); os.IsNotExist(err) {
		inputMailConfig()
	}
	configFile, err := os.Open("emailconfig.xml")
	if err != nil {
		return mailConfig{}
	}
	defer configFile.Close()
	var mc mailConfig
	xml.NewDecoder(configFile).Decode(&mc)
	return mc
}

// 初始化邮箱配置并作为结构体返回给其他方法
func inputMailConfig() {
	var hosts string
	var users string
	var password string
	fmt.Print("[FeasOJ]请输入Smtp服务器地址:")
	fmt.Scanln(&hosts)
	fmt.Print("[FeasOJ]请输入发件人邮箱地址（如114514@qq.com）:")
	fmt.Scanln(&users)
	fmt.Print("[FeasOJ]请输入邮箱密码（不一定是登录密码）:")
	fmt.Scanln(&password)

	// 写入配置到emailconfig.xml文件中
	config := mailConfig{Host: hosts, Port: 465, User: users, Pass: password}
	configFile, _ := os.Create("emailconfig.xml")
	defer configFile.Close()
	xml.NewEncoder(configFile).Encode(config)
}

// 随机生成4位数字验证码
func generateVerifycode() string {
	return fmt.Sprintf("%04d", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))
}

// 发送验证码
func sendVerifycode(config mailConfig, to string, verifycode string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", config.User)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "FeasOJ验证码")
	m.SetBody("text/plain", "您的验证码是："+verifycode)
	d := gomail.NewDialer(config.Host, config.Port, config.User, config.Pass)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

// 检验验证码是否正确
func compareVerifyCode(frontendCode, backendCode string) bool {
	return frontendCode == backendCode
}
