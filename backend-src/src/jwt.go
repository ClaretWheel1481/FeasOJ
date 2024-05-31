package main

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

// 用户Token生成后返回给前端
func GenerateToken(username string) string {
	token := jwt.New(jwt.SigningMethodHS256)
	// 设置Token的Claims
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	// 生成Token
	tokenString, err := token.SignedString([]byte(selectTokenSecret(username)))
	if err != nil {
		fmt.Println("生成Token失败：", err)
		return ""
	}
	return tokenString
}

// 校验Token与username是不是配对
func VerifyToken(username, tokenString string) bool {
	// 解析Token Username
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 验证签名
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(selectTokenSecret(username)), nil
	})
	if token.Valid {
		return true
	}
	return err == nil
}
