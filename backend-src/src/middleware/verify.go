package middleware

import (
	"net/http"
	"net/url"
	"src/utils"

	"github.com/gin-gonic/gin"
)

func HeaderVerify() gin.HandlerFunc {
	return func(c *gin.Context) {
		encodedUsername := c.GetHeader("username")
		username, err := url.QueryUnescape(encodedUsername)
		token := c.GetHeader("Authorization")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "未找到该用户。"})
			return
		}
		if !utils.VerifyToken(username, token) {
			c.JSON(http.StatusUnauthorized, gin.H{"status": 401, "message": "Token验证失败。"})
			return
		}
		c.Next()
	}
}
