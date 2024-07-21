package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// TODO:为安全性考虑，未来将对任何API操作加入Token验证

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
	token := c.GetHeader("Authorization")
	encodedUsername := c.GetHeader("username")
	username, err := url.QueryUnescape(encodedUsername)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "无法获取用户名。"})
		return
	}
	if VerifyToken(username, token) {
		c.JSON(http.StatusOK, gin.H{"status": 200, "message": "Token验证成功。"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "Token验证失败。"})
	}
}

// 获取用户信息
func getUserInfos(c *gin.Context) {
	// 返回至前端以显示个人资料
	encodedUsername := c.GetHeader("username")
	username, err := url.QueryUnescape(encodedUsername)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "无法获取用户名。"})
		return
	}
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
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "无法获取头像文件。"})
		return
	}
	token := c.GetHeader("Authorization")
	encodedUsername := c.GetHeader("username")
	username, err := url.QueryUnescape(encodedUsername)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "无法获取用户名。"})
		return
	}
	// 校验Token
	if !VerifyToken(username, token) {
		c.JSON(http.StatusUnauthorized, gin.H{"status": 401, "message": "Token验证失败。"})
		return
	}
	// 获取用户信息
	userInfo := selectUserInfo(username)
	if userInfo.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "无法获取用户信息。"})
		return
	}
	newFilename := fmt.Sprintf("%d%s", userInfo.Uid, path.Ext(file.Filename))
	filePath := filepath.Join("../avatars", newFilename)
	if _, err := os.Stat(filePath); err == nil {
		os.Remove(filePath)
	}
	// 上传头像至指定路径
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "message": "头像上传失败。"})
		return
	}
	// 上传头像路径至数据库
	if !updateAvatar(username, filepath.Join(newFilename)) {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "message": "头像路径更新失败。"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": 200, "message": "头像上传成功。"})
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
func getDiscussionByDids(c *gin.Context) {
	discussion := selectDiscussionByDid(c.Param("Did"))
	c.JSON(http.StatusOK, gin.H{"discussionInfo": discussion})
}

// 创建讨论
func createDiscussion(c *gin.Context) {
	encodedUsername := c.GetHeader("username")
	username, err := url.QueryUnescape(encodedUsername)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "未找到该用户。"})
		return
	}
	title := c.PostForm("title")
	content := c.PostForm("content")
	uid := selectUserInfo(username).Uid
	if !addDiscussion(title, content, uid) {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "message": "创建讨论失败。"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": 200, "message": "创建讨论成功。"})
}

// 删除讨论
func deleteDiscussion(c *gin.Context) {
	did := c.Param("Did")
	if delDiscussion(did) {
		c.JSON(http.StatusOK, gin.H{"status": 200, "message": "删除讨论成功。"})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "message": "删除讨论失败。"})
	}
}

// 上传代码
func uploadCodes(c *gin.Context) {
	file, err := c.FormFile("code")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "无法获取代码文件。"})
		return
	}
	problem := c.Query("problem")
	encodedUsername := c.GetHeader("username")
	username, err := url.QueryUnescape(encodedUsername)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "无法获取用户名。"})
		return
	}
	token := c.GetHeader("Authorization")
	if !VerifyToken(username, token) {
		c.JSON(http.StatusUnauthorized, gin.H{"status": 401, "message": "Token验证失败。"})
		return
	}
	// 获取用户ID
	userInfo := selectUserInfo(username)
	// 将文件名改为用户ID_题目ID
	file.Filename = fmt.Sprintf("%d_%s%s", userInfo.Uid, problem, path.Ext(file.Filename))
	// 保存文件到指定路径
	if err := c.SaveUploadedFile(file, "../codefiles/"+file.Filename); err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": 200, "message": "代码上传成功。"})
}

// 管理员获取指定题目所有信息
func getProblemAllInfos(c *gin.Context) {
	username := c.GetHeader("username")
	token := c.GetHeader("Authorization")
	if !VerifyToken(username, token) {
		c.JSON(http.StatusUnauthorized, gin.H{"status": 401, "message": "Token验证失败。"})
		return
	}
	if selectUserInfo(username).Role != 1 {
		c.JSON(http.StatusUnauthorized, gin.H{"status": 401, "message": "权限不足。"})
		return
	}
	problemInfo := selectProblemTestCases(c.Param("Pid"))
	c.JSON(http.StatusOK, gin.H{"problemInfo": problemInfo})
}

// 更新题目信息
func updateProblemInfos(c *gin.Context) {
	var req adminProblemInfoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "请求参数错误"})
		return
	}

	// 验证用户权限
	username := c.GetHeader("username")
	token := c.GetHeader("Authorization")
	if !VerifyToken(username, token) {
		c.JSON(http.StatusUnauthorized, gin.H{"status": 401, "message": "Token验证失败"})
		return
	}
	if selectUserInfo(username).Role != 1 {
		c.JSON(http.StatusUnauthorized, gin.H{"status": 401, "message": "权限不足"})
		return
	}

	// 更新题目信息
	if err := updateProblem(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "message": "更新题目信息失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 200, "message": "题目信息更新成功"})
}

// 删除题目及其输入输出样例
func deleteProblems(c *gin.Context) {
	username := c.GetHeader("username")
	token := c.GetHeader("Authorization")
	pid := c.Param("Pid")
	pidInt, err := strconv.Atoi(pid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "题目ID错误"})
		return
	}
	if !VerifyToken(username, token) {
		c.JSON(http.StatusUnauthorized, gin.H{"status": 401, "message": "Token验证失败。"})
		return
	}
	if selectUserInfo(username).Role != 1 {
		c.JSON(http.StatusUnauthorized, gin.H{"status": 401, "message": "权限不足。"})
		return
	}
	if !deleteProblemAllInfo(pidInt) {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "删除失败。"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": 200, "message": "删除成功。"})
}

// 获取指定讨论的评论
func getComments(c *gin.Context) {
	username := c.GetHeader("username")
	token := c.GetHeader("Authorization")
	did := c.Param("Did")
	didInt, err := strconv.Atoi(did)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "讨论ID错误"})
	}
	if !VerifyToken(username, token) {
		c.JSON(http.StatusUnauthorized, gin.H{"status": 401, "message": "Token验证失败。"})
		return
	}
	comments := getCommentsByDid(didInt)
	c.JSON(http.StatusOK, gin.H{"comments": comments})
}

// 删除指定Cid的评论
func delComments(c *gin.Context) {
	username := c.GetHeader("username")
	token := c.GetHeader("Authorization")
	Cid := c.Param("Cid")
	CidInt, err := strconv.Atoi(Cid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "评论ID错误"})
	}
	if !VerifyToken(username, token) {
		c.JSON(http.StatusUnauthorized, gin.H{"status": 401, "message": "Token验证失败。"})
		return
	}
	if !deleteCommentByCid(CidInt) {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "删除失败。"})
	}
	c.JSON(http.StatusOK, gin.H{"status": 200, "message": "删除成功。"})
}
