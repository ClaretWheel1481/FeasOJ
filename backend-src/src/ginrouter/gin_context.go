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

// 注册
func Register(c *gin.Context) {
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
func Login(c *gin.Context) {
	Username := c.Query("username")
	Password := c.Query("password")
	userPassword := utils.SelectPassword(Username)
	if userPassword == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "用户不存在。"})
	} else {
		// 用户是否被封禁
		if utils.SelectUserInfo(Username).IsBan {
			c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "该用户已被封禁。"})
			return
		}
		// 校验密码是否正确
		if account.VerifyPassword(Password, userPassword) {
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
func GetCaptcha(c *gin.Context) {
	// 获取邮箱地址
	emails := c.Query("email")
	if utils.SendVerifycode(utils.InitEmailConfig(), emails, utils.GenerateVerifycode()) == "Success" {
		c.JSON(http.StatusOK, gin.H{"status": 200, "message": "验证码发送成功。"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "验证码发送失败，可能是我们的问题。"})
	}
}

// 获取用户信息
func GetUserInfo(c *gin.Context) {
	token := c.GetHeader("Authorization")
	// 返回至前端以显示个人资料
	encodedUsername := c.GetHeader("username")
	username, err := url.QueryUnescape(encodedUsername)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "无法获取用户名。"})
		return
	}
	if !utils.VerifyToken(username, token) {
		c.JSON(http.StatusUnauthorized, gin.H{"status": 401, "message": "Token验证失败。"})
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
func UpdatePassword(c *gin.Context) {
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
func UploadAvatar(c *gin.Context) {
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
func GetAllProblems(c *gin.Context) {
	problems := utils.SelectAllProblems()
	c.JSON(http.StatusOK, gin.H{"problems": problems})
}

// 获取题目信息
func GetProblemInfo(c *gin.Context) {
	problemInfo := utils.SelectProblemInfo(c.Param("id"))
	c.JSON(http.StatusOK, gin.H{"problemInfo": problemInfo})
}

// 获取所有提交记录
func GetAllSubmitRecords(c *gin.Context) {
	submitrecords := utils.SelectAllSubmitRecords()
	c.JSON(http.StatusOK, gin.H{"submitrecords": submitrecords})
}

// 获取指定用户提交记录
func GetSubmitRecordsByUid(c *gin.Context) {
	submitrecords := utils.SelectSubmitRecordsByUid(c.Param("uid"))
	c.JSON(http.StatusOK, gin.H{"submitrecords": submitrecords})
}

// 获取所有讨论列表
func GetAllDiscussions(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	itemsPerPage, _ := strconv.Atoi(c.DefaultQuery("itemsPerPage", "12"))

	discussions, total := utils.SelectDiscussList(page, itemsPerPage)
	c.JSON(http.StatusOK, gin.H{
		"discussions": discussions,
		"total":       total,
	})
}

// 获取指定id讨论信息
func GetDiscussionByDid(c *gin.Context) {
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
	Did := c.Param("Did")
	if utils.DelDiscussion(Did) {
		c.JSON(http.StatusOK, gin.H{"status": 200, "message": "删除讨论成功。"})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "message": "删除讨论失败。"})
	}
}

// 上传代码
func UploadCode(c *gin.Context) {
	file, err := c.FormFile("code")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "无法获取代码文件。"})
		return
	}
	problem := c.Query("problem")
	pidInt, _ := strconv.Atoi(problem)
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
	var language string
	if path.Ext(file.Filename) == ".cpp" {
		language = "C++"
	} else if path.Ext(file.Filename) == ".java" {
		language = "Java"
	} else if path.Ext(file.Filename) == ".py" {
		language = "Python"
	} else if path.Ext(file.Filename) == ".go" {
		language = "Go"
	} else {
		language = "Unknown"
	}

	// 上传任务至Redis任务队列
	rdb := utils.InitRedis()
	err = rdb.RPush("judgeTask", file.Filename).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "message": "代码任务提交失败，请尝试重新提交。"})
		return
	}
	utils.AddSubmitRecord(userInfo.Uid, pidInt, "Running...", language)
	c.JSON(http.StatusOK, gin.H{"status": 200, "message": "代码提交成功，请等候判题结果。"})
}

// 管理员获取指定题目所有信息
func GetProblemAllInfo(c *gin.Context) {
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
func UpdateProblemInfo(c *gin.Context) {
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
func DeleteProblem(c *gin.Context) {
	encodedUsername := c.GetHeader("username")
	username, err := url.QueryUnescape(encodedUsername)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "无法获取用户名。"})
		return
	}
	token := c.GetHeader("Authorization")
	pid := c.Param("Pid")
	pidInt, _ := strconv.Atoi(pid)
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

// 管理员获取所有用户信息
func GetAllUsersInfo(c *gin.Context) {
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
	usersInfo := utils.GetAllUsersInfo()
	c.JSON(http.StatusOK, gin.H{"usersInfo": usersInfo})
}

// 获取指定讨论的评论
func GetComment(c *gin.Context) {
	encodedUsername := c.GetHeader("username")
	username, err := url.QueryUnescape(encodedUsername)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "无法获取用户名。"})
		return
	}
	token := c.GetHeader("Authorization")
	Did := c.Param("Did")
	DidInt, _ := strconv.Atoi(Did)
	if !utils.VerifyToken(username, token) {
		c.JSON(http.StatusUnauthorized, gin.H{"status": 401, "message": "Token验证失败。"})
		return
	}
	comments := utils.GetCommentsByDid(DidInt)
	c.JSON(http.StatusOK, gin.H{"comments": comments})
}

// 删除指定Cid的评论
func DelComment(c *gin.Context) {
	Cid := c.Param("Cid")
	CidInt, _ := strconv.Atoi(Cid)
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
	if !utils.DeleteCommentByCid(CidInt) {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "删除失败。"})
	}
	c.JSON(http.StatusOK, gin.H{"status": 200, "message": "删除成功。"})
}

// 添加评论
func AddComment(c *gin.Context) {
	encodedUsername := c.GetHeader("username")
	username, err := url.QueryUnescape(encodedUsername)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "无法获取用户名。"})
		return
	}
	token := c.GetHeader("Authorization")
	content := c.PostForm("content")
	Did := c.Param("Did")
	DidInt, _ := strconv.Atoi(Did)
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

// 晋升用户
func PromoteUser(c *gin.Context) {
	uid := c.Query("uid")
	uidInt, _ := strconv.Atoi(uid)
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
	if !utils.PromoteToAdmin(uidInt) {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "晋升失败。"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": 200, "message": "晋升成功。"})
}

// 降级用户
func DemoteUser(c *gin.Context) {
	uid := c.Query("uid")
	uidInt, _ := strconv.Atoi(uid)
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
	if !utils.DemoteToUser(uidInt) {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "降级失败。"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": 200, "message": "降级成功。"})
}

// 封禁用户
func BanUser(c *gin.Context) {
	uid := c.Query("uid")
	uidInt, _ := strconv.Atoi(uid)
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
	if !utils.BanUser(uidInt) {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "封禁失败。"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": 200, "message": "封禁成功。"})
}

// 解封用户
func UnbanUser(c *gin.Context) {
	uid := c.Query("uid")
	uidInt, _ := strconv.Atoi(uid)
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
	if !utils.UnbanUser(uidInt) {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "解封失败。"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": 200, "message": "解封成功。"})
}
