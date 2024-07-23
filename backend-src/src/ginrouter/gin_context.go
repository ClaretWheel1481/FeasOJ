package ginrouter

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"src/account"
	"src/global"
	"src/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// TODO:为安全性考虑，未来将对任何API操作加入Token验证

// 注册
func Registers(c *gin.Context) {
	var req global.RegisterRequest
	c.ShouldBind(&req)
	// 判断用户或邮箱是否存在
	if utils.IsUserExist(req.Username, req.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "用户已存在或邮箱已使用。"})
		return
	}
	tokensecret := uuid.New().String()
	vcodeStatus := utils.CompareVerifyCode(req.Vcode, req.Email)
	if !vcodeStatus {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "验证码错误。"})
		return
	}
	regstatus := utils.Register(req.Username, account.EncryptPassword(req.Password), req.Email, tokensecret, 0)
	if regstatus {
		c.JSON(http.StatusOK, gin.H{"status": 200, "message": "注册成功，2秒后跳转至登录界面。"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "注册失败，请确认邮箱验证码是否正确。"})
	}
}

// 登录
func Logins(c *gin.Context) {
	Username := c.Query("username")
	Password := c.Query("password")
	userPHash := utils.SelectPassword(Username)
	if userPHash == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "用户不存在。"})
	} else {
		// 校验密码是否正确
		if account.VerifyPassword(Password, userPHash) {
			// 生成Token并返回至前端
			token, err := utils.GenerateToken(Username)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "生成Token失败。"})
			}
			c.JSON(http.StatusOK, gin.H{"status": 200, "message": "登录成功。", "token": token})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "密码错误。"})
		}
	}
}

// 获取验证码
func GetCaptchas(c *gin.Context) {
	// 获取邮箱地址
	emails := c.Query("email")
	if utils.SendVerifycode(utils.InitEmailConfig(), emails, utils.GenerateVerifycode()) == "Success" {
		c.JSON(http.StatusOK, gin.H{"status": 200, "message": "验证码发送成功。"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "验证码发送失败，可能是我们的问题。"})
	}
}

// 校验TOKEN
func VerifyTokens(c *gin.Context) {
	token := c.GetHeader("Authorization")
	encodedUsername := c.GetHeader("username")
	username, err := url.QueryUnescape(encodedUsername)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "无法获取用户名。"})
		return
	}
	if utils.VerifyToken(username, token) {
		c.JSON(http.StatusOK, gin.H{"status": 200, "message": "Token验证成功。"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "Token验证失败。"})
	}
}

// 获取用户信息
func GetUserInfos(c *gin.Context) {
	// 返回至前端以显示个人资料
	encodedUsername := c.GetHeader("username")
	username, err := url.QueryUnescape(encodedUsername)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "无法获取用户名。"})
		return
	}
	// 根据uid来查询对应的用户信息
	userInfo := utils.SelectUserInfo(username)
	if userInfo.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "用户不存在。"})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": 200, "Info": userInfo})
	}

}

// 更新密码
func UpdatePasswords(c *gin.Context) {
	var req global.UpdatePasswordRequest
	c.ShouldBind(&req)
	vcodeStatus := utils.CompareVerifyCode(req.Vcode, req.Email)
	if !vcodeStatus {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "验证码错误。"})
		return
	}
	if utils.UpdatePassword(req.Email, req.NewPassword) {
		c.JSON(http.StatusOK, gin.H{"status": 200, "message": "密码修改成功，2秒后跳转至登录页面。"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "密码修改失败。"})
	}
}

// 上传头像
func UploadAvatars(c *gin.Context) {
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
	if !utils.VerifyToken(username, token) {
		c.JSON(http.StatusUnauthorized, gin.H{"status": 401, "message": "Token验证失败。"})
		return
	}
	// 获取用户信息
	userInfo := utils.SelectUserInfo(username)
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
	if !utils.UpdateAvatar(username, filepath.Join(newFilename)) {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "message": "头像路径更新失败。"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": 200, "message": "头像上传成功。"})
}

// 获取所有题目
func GetAllProblemss(c *gin.Context) {
	problems := utils.SelectAllProblems()
	c.JSON(http.StatusOK, gin.H{"problems": problems})
}

// 获取题目信息
func GetProblemInfos(c *gin.Context) {
	problemInfo := utils.SelectProblemInfo(c.Param("id"))
	c.JSON(http.StatusOK, gin.H{"problemInfo": problemInfo})
}

// 获取所有提交记录
func GetAllSubmitRecordss(c *gin.Context) {
	submitrecords := utils.SelectAllSubmitRecords()
	c.JSON(http.StatusOK, gin.H{"submitrecords": submitrecords})
}

// 获取指定用户提交记录
func GetSubmitRecordsByUids(c *gin.Context) {
	submitrecords := utils.SelectSubmitRecordsByUid(c.Param("uid"))
	c.JSON(http.StatusOK, gin.H{"submitrecords": submitrecords})
}

// 获取所有讨论列表
func GetAllDiscussionss(c *gin.Context) {
	discussions := utils.SelectDiscussList()
	c.JSON(http.StatusOK, gin.H{"discussions": discussions})
}

// 获取指定id讨论信息
func GetDiscussionByDids(c *gin.Context) {
	discussion := utils.SelectDiscussionByDid(c.Param("Did"))
	c.JSON(http.StatusOK, gin.H{"discussionInfo": discussion})
}

// 创建讨论
func CreateDiscussion(c *gin.Context) {
	encodedUsername := c.GetHeader("username")
	username, err := url.QueryUnescape(encodedUsername)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "未找到该用户。"})
		return
	}
	title := c.PostForm("title")
	content := c.PostForm("content")
	uid := utils.SelectUserInfo(username).Uid
	if !utils.AddDiscussion(title, content, uid) {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "message": "创建讨论失败。"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": 200, "message": "创建讨论成功。"})
}

// 删除讨论
func DeleteDiscussion(c *gin.Context) {
	did := c.Param("Did")
	if utils.DelDiscussion(did) {
		c.JSON(http.StatusOK, gin.H{"status": 200, "message": "删除讨论成功。"})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "message": "删除讨论失败。"})
	}
}

// 上传代码
func UploadCodes(c *gin.Context) {
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
	if !utils.VerifyToken(username, token) {
		c.JSON(http.StatusUnauthorized, gin.H{"status": 401, "message": "Token验证失败。"})
		return
	}
	// 获取用户ID
	userInfo := utils.SelectUserInfo(username)
	// 将文件名改为用户ID_题目ID
	file.Filename = fmt.Sprintf("%d_%s%s", userInfo.Uid, problem, path.Ext(file.Filename))
	// 保存文件到指定路径
	if err := c.SaveUploadedFile(file, "../codefiles/"+file.Filename); err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": 200, "message": "代码上传成功。"})
}

// 管理员获取指定题目所有信息
func GetProblemAllInfos(c *gin.Context) {
	encodedUsername := c.GetHeader("username")
	username, err := url.QueryUnescape(encodedUsername)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "无法获取用户名。"})
		return
	}
	token := c.GetHeader("Authorization")
	if !utils.VerifyToken(username, token) {
		c.JSON(http.StatusUnauthorized, gin.H{"status": 401, "message": "Token验证失败。"})
		return
	}
	if utils.SelectUserInfo(username).Role != 1 {
		c.JSON(http.StatusUnauthorized, gin.H{"status": 401, "message": "权限不足。"})
		return
	}
	problemInfo := utils.SelectProblemTestCases(c.Param("Pid"))
	c.JSON(http.StatusOK, gin.H{"problemInfo": problemInfo})
}

// 更新题目信息
func UpdateProblemInfos(c *gin.Context) {
	var req global.AdminProblemInfoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "请求参数错误"})
		return
	}

	// 验证用户权限
	encodedUsername := c.GetHeader("username")
	username, err := url.QueryUnescape(encodedUsername)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "无法获取用户名。"})
		return
	}
	token := c.GetHeader("Authorization")
	if !utils.VerifyToken(username, token) {
		c.JSON(http.StatusUnauthorized, gin.H{"status": 401, "message": "Token验证失败"})
		return
	}
	if utils.SelectUserInfo(username).Role != 1 {
		c.JSON(http.StatusUnauthorized, gin.H{"status": 401, "message": "权限不足"})
		return
	}

	// 更新题目信息
	if err := utils.UpdateProblem(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "message": "更新题目信息失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 200, "message": "题目信息更新成功"})
}

// 删除题目及其输入输出样例
func DeleteProblems(c *gin.Context) {
	encodedUsername := c.GetHeader("username")
	username, err := url.QueryUnescape(encodedUsername)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "无法获取用户名。"})
		return
	}
	token := c.GetHeader("Authorization")
	pid := c.Param("Pid")
	pidInt, err := strconv.Atoi(pid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "题目ID错误"})
		return
	}
	if !utils.VerifyToken(username, token) {
		c.JSON(http.StatusUnauthorized, gin.H{"status": 401, "message": "Token验证失败。"})
		return
	}
	if utils.SelectUserInfo(username).Role != 1 {
		c.JSON(http.StatusUnauthorized, gin.H{"status": 401, "message": "权限不足。"})
		return
	}
	if !utils.DeleteProblemAllInfo(pidInt) {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "删除失败。"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": 200, "message": "删除成功。"})
}

// 获取指定讨论的评论
func GetComments(c *gin.Context) {
	encodedUsername := c.GetHeader("username")
	username, err := url.QueryUnescape(encodedUsername)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "无法获取用户名。"})
		return
	}
	token := c.GetHeader("Authorization")
	did := c.Param("Did")
	didInt, err := strconv.Atoi(did)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "讨论ID错误"})
	}
	if !utils.VerifyToken(username, token) {
		c.JSON(http.StatusUnauthorized, gin.H{"status": 401, "message": "Token验证失败。"})
		return
	}
	comments := utils.GetCommentsByDid(didInt)
	c.JSON(http.StatusOK, gin.H{"comments": comments})
}

// 删除指定Cid的评论
func DelComments(c *gin.Context) {
	encodedUsername := c.GetHeader("username")
	username, err := url.QueryUnescape(encodedUsername)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "无法获取用户名。"})
		return
	}
	token := c.GetHeader("Authorization")
	Cid := c.Param("Cid")
	CidInt, err := strconv.Atoi(Cid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "评论ID错误"})
	}
	if !utils.VerifyToken(username, token) {
		c.JSON(http.StatusUnauthorized, gin.H{"status": 401, "message": "Token验证失败。"})
		return
	}
	if !utils.DeleteCommentByCid(CidInt) {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "删除失败。"})
	}
	c.JSON(http.StatusOK, gin.H{"status": 200, "message": "删除成功。"})
}

// 添加评论
func AddComments(c *gin.Context) {
	encodedUsername := c.GetHeader("username")
	username, err := url.QueryUnescape(encodedUsername)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "无法获取用户名。"})
		return
	}
	token := c.GetHeader("Authorization")
	content := c.PostForm("content")
	Did := c.Param("Did")
	DidInt, err := strconv.Atoi(Did)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "评论ID错误"})
	}
	if !utils.VerifyToken(username, token) {
		c.JSON(http.StatusUnauthorized, gin.H{"status": 401, "message": "Token验证失败。"})
		return
	}
	// 获取用户ID
	userInfo := utils.SelectUserInfo(username)
	if !utils.AddComment(content, DidInt, userInfo.Uid) {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "添加失败。"})
	}
	c.JSON(http.StatusOK, gin.H{"status": 200, "message": "添加成功。"})
}
