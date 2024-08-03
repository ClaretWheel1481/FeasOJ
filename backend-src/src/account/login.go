package account

import (
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

// 用户密码加密
func EncryptPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(bytes)
}

// 密码验证（将用户输入的密码与数据库中的密码进行比较）
func VerifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// 判断是否邮箱登录
func IsEmail(email string) bool {
	pattern := `^\w+([-+.]?\w+)*@\w+([-.]?\w+)*\.\w+([-.]?\w+)*$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}
