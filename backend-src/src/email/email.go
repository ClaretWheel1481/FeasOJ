package email

import (
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"src/global"
	"src/structs"
	"src/utils"
	"time"

	"gopkg.in/gomail.v2"
)

// 从emailConfig.xml中读取邮箱配置并返回mailConfig
func InitEmailConfig() structs.MailConfig {
	// 判断是否有emailconfig.xml文件
	filePath := filepath.Join(global.ConfigsDir, "emailconfig.xml")
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		InputMailConfig()
	}
	configFile, err := os.Open(filePath)
	if err != nil {
		return structs.MailConfig{}
	}
	defer configFile.Close()
	var mc structs.MailConfig
	xml.NewDecoder(configFile).Decode(&mc)
	return mc
}

// 初始化邮箱配置并作为结构体返回给其他方法
func InputMailConfig() {
	var hosts string
	var users string
	var password string
	fmt.Print("[FeasOJ]请输入Smtp服务器地址:")
	fmt.Scanln(&hosts)
	fmt.Print("[FeasOJ]请输入发件人邮箱地址(如114514@qq.com):")
	fmt.Scanln(&users)
	fmt.Print("[FeasOJ]请输入邮箱密码(不一定是登录密码):")
	fmt.Scanln(&password)

	// 写入配置到emailconfig.xml文件中
	config := structs.MailConfig{Host: hosts, Port: 465, User: users, Pass: password}
	filePath := filepath.Join(global.ConfigsDir, "emailconfig.xml")
	configFile, _ := os.Create(filePath)
	defer configFile.Close()
	xml.NewEncoder(configFile).Encode(config)
}

// 随机生成4位数字验证码
func GenerateVerifycode() string {
	return fmt.Sprintf("%04d", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))
}

// 发送验证码
func SendVerifycode(config structs.MailConfig, to string, verifycode string) string {
	m := gomail.NewMessage()
	m.SetAddressHeader("From", config.User, "FeasOJ")
	m.SetHeader("To", to)
	m.SetHeader("Subject", "FeasOJ验证码")
	m.SetBody("text/plain", "注意，您正在进行敏感操作：\n您的验证码是："+verifycode+"，5分钟后失效。")
	d := gomail.NewDialer(config.Host, config.Port, config.User, config.Pass)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		return err.Error()
	}
	// 将验证码同时存进Redis中等待校验
	rdb := utils.InitRedis()
	err := rdb.Set(to, verifycode, 5*time.Minute).Err()
	if err != nil {
		return err.Error()
	}
	return "Success"
}

// 检验Redis中验证码与前端返回的是否相同
func CompareVerifyCode(frontendCode, to string) bool {
	// 通过邮箱来获取Redis中的验证码
	rdb := utils.InitRedis()
	verifyCode, err := rdb.Get(to).Result()
	if err != nil {
		return false
	}
	if verifyCode == frontendCode {
		// 移除Redis中的验证码
		rdb.Del(to)
		return true
	}
	return false
}
