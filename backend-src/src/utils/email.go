package utils

import (
	"crypto/tls"
	"fmt"
	"math/rand"
	"src/global"
	"time"

	"gopkg.in/gomail.v2"
)

// 随机生成4位数字验证码
func GenerateVerifycode() string {
	return fmt.Sprintf("%04d", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))
}

// 发送验证码
func SendVerifycode(config global.MailConfig, to string, verifycode string) bool {
	m := gomail.NewMessage()
	m.SetAddressHeader("From", config.User, "FeasOJ")
	m.SetHeader("To", to)
	m.SetHeader("Subject", "FeasOJ验证码")
	m.SetBody("text/plain", "注意，您正在进行敏感操作：\n您的验证码是："+verifycode+"，5分钟后失效。")
	d := gomail.NewDialer(config.Host, config.Port, config.User, config.Pass)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		return false
	}
	// 将验证码同时存进Redis中等待校验
	rdb := InitRedis()
	err := rdb.Set(to, verifycode, 5*time.Minute).Err()
	return err == nil
}

// 检验Redis中验证码与前端返回的是否相同
func CompareVerifyCode(frontendCode, to string) bool {
	// 通过邮箱来获取Redis中的验证码
	rdb := InitRedis()
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
