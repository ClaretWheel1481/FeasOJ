package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// 注册
func registers(c *gin.Context) {
	var req RegisterRequest
	c.ShouldBind(&req)
	// 判断用户或邮箱是否存在
	if isUserExist(req.Username, req.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "用户已存在或邮箱已使用。"})
		return
	}
	tokensecret := uuid.New().String()
	vcodeStatus := compareVerifyCode(req.Vcode, req.Email)
	if !vcodeStatus {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "验证码错误。"})
		return
	}
	regstatus := register(req.Username, encryptPassword(req.Password), req.Email, tokensecret, 0)
	if regstatus {
		c.JSON(http.StatusOK, gin.H{"status": 200, "message": "注册成功，2秒后跳转至登录界面。"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "注册失败，请确认邮箱验证码是否正确。"})
	}
}

// 登录
func logins(c *gin.Context) {
	Username := c.Query("username")
	Password := c.Query("password")
	userPHash := selectPassword(Username)
	if userPHash == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "用户不存在。"})
	} else {
		// 校验密码是否正确
		if verifyPassword(Password, userPHash) {
			// 生成Token并返回至前端
			token := GenerateToken(Username)
			c.JSON(http.StatusOK, gin.H{"status": 200, "message": "登录成功。", "token": token})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "密码错误。"})
		}
	}
}

// 获取验证码
func getCaptchas(c *gin.Context) {
	// 获取邮箱地址
	email := c.Query("email")
	if sendVerifycode(initEmailConfig(), email, generateVerifycode()) == "Success" {
		c.JSON(http.StatusOK, gin.H{"status": 200, "message": "验证码发送成功。"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "验证码发送失败，可能是我们的问题。"})
	}
}

// 校验TOKEN
func verifyTokens(c *gin.Context) {
	token := c.GetHeader("token")
	username := c.Query("username")
	if VerifyToken(username, token) {
		c.JSON(http.StatusOK, gin.H{"status": 200, "message": "Token验证成功。"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "Token验证失败。"})
	}
}

// 获取用户信息
func getUserInfos(c *gin.Context) {
	// 返回至前端以显示个人资料
	username := c.Query("username")
	// 根据uid来查询对应的用户信息
	userInfo := selectUserInfo(username)
	if userInfo.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "用户不存在。"})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": 200, "Info": userInfo})
	}

}

// 更新密码
func updatePasswords(c *gin.Context) {
	var req updatePasswordRequest
	c.ShouldBind(&req)
	vcodeStatus := compareVerifyCode(req.Vcode, req.Email)
	if !vcodeStatus {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "验证码错误。"})
		return
	}
	if updatePassword(req.Email, req.NewPassword) {
		c.JSON(http.StatusOK, gin.H{"status": 200, "message": "密码修改成功，2秒后跳转至登录页面。"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "密码修改失败。"})
	}
}

// 上传头像
func uploadAvatars(c *gin.Context) {
	file, err := c.FormFile("avatar")
	token := c.GetHeader("token")
	username := c.Query("username")
	// 校验Token
	if token == "" || !VerifyToken(username, token) {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "Token验证失败。"})
		return
	}
	// 检查上传文件类型为png或jpg
	if !strings.HasSuffix(file.Filename, ".png") && !strings.HasSuffix(file.Filename, ".jpg") && !strings.HasSuffix(file.Filename, ".JPG") && !strings.HasSuffix(file.Filename, ".PNG") {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "头像文件类型错误。"})
		return
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "头像上传失败。"})
	}

	// 压缩文件并修改文件名为数据库中用户的uid，例如user_1.png/jpg
	user := fmt.Sprintf("%d", selectUserInfo(username).Uid)
	file.Filename = "user_" + user + file.Filename[strings.LastIndex(file.Filename, "."):]
	// 检查是否有重复名称，重复则删除
	filepath := "../avatars/" + file.Filename
	if _, err := os.Stat(filepath); err == nil {
		os.Remove(filepath)
	}
	// 存放到avatars文件夹中
	err = c.SaveUploadedFile(file, "../avatars/"+file.Filename)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "头像上传失败。"})
	}

	// 更新数据库中的头像路径地址
	if updateAvatar(username, "..\\..\\backend-src\\avatars\\"+file.Filename) {
		c.JSON(http.StatusOK, gin.H{"status": 200, "message": "头像上传成功。"})
	}
}

// 获取所有题目
func getAllProblemss(c *gin.Context) {
	problems := selectAllProblems()
	c.JSON(http.StatusOK, gin.H{"problems": problems})
}

// 获取题目信息
func getProblemInfos(c *gin.Context) {
	problemInfo := selectProblemInfo(c.Param("id"))
	c.JSON(http.StatusOK, gin.H{"problemInfo": problemInfo})
}

// 获取所有竞赛API
func getAllContestss(c *gin.Context) {
	allContest := selectAllContests()
	c.JSON(http.StatusOK, gin.H{"contests": allContest})
}

// 获取竞赛信息
func getContestInfos(c *gin.Context) {
	contestInfo := selectProblemByCid(c.Param("cid"))
	c.JSON(http.StatusOK, gin.H{"contestInfo": contestInfo})
}

// 获取所有提交记录
func getAllSubmitRecordss(c *gin.Context) {
	submitrecords := selectAllSubmitRecords()
	c.JSON(http.StatusOK, gin.H{"submitrecords": submitrecords})
}

// 获取指定用户提交记录
func getSubmitRecordsByUids(c *gin.Context) {
	submitrecords := selectSubmitRecordsByUid(c.Param("uid"))
	c.JSON(http.StatusOK, gin.H{"submitrecords": submitrecords})
}

// 获取所有讨论列表
func getAllDiscussionss(c *gin.Context) {
	discussions := selectDiscussList()
	c.JSON(http.StatusOK, gin.H{"discussions": discussions})
}

// 获取指定id讨论信息
func getDiscussionByTids(c *gin.Context) {
	discussion := selectDiscussionByTid(c.Param("tid"))
	c.JSON(http.StatusOK, gin.H{"discussionInfo": discussion})
}
