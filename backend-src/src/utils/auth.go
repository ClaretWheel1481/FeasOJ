package utils

import (
	"fmt"
	"regexp"
	"time"

	"github.com/golang-jwt/jwt/v5"
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

// 用户Token生成后返回给前端
func GenerateToken(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	// 设置Token的Claims
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(30 * 24 * time.Hour).Unix() // 设置30天过期时间
	// 生成Token
	tokenString, err := token.SignedString([]byte(SelectUser(username).TokenSecret))
	if err != nil {
		return "", fmt.Errorf("生成Token失败：%v", err)
	}
	return tokenString, nil
}

// 校验Token与username是不是配对
func VerifyToken(username, tokenString string) bool {
	// 解析Token Username
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 验证签名
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("%v", token.Header["alg"])
		}
		return []byte(SelectUser(username).TokenSecret), nil
	})
	if token.Valid {
		return true
	}
	return err == nil
}
