package utils

import (
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"src/account"
	"src/global"
	"time"

	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InputSqlInfo() bool {
	var Sqlconfig global.SqlConfig
	fmt.Println("[FeasOJ]请输入数据库连接信息：")
	fmt.Print("数据库名：")
	fmt.Scanln(&Sqlconfig.DbName)
	fmt.Print("数据库用户名：")
	fmt.Scanln(&Sqlconfig.DbUser)
	fmt.Print("数据库密码：")
	fmt.Scanln(&Sqlconfig.DbPassword)
	fmt.Print("数据库地址(加上端口)：")
	fmt.Scanln(&Sqlconfig.DbAddress)
	fmt.Println("[FeasOJ]正在保存数据库连接信息...")
	// 将连接信息作为xml保存到目录
	configXml, err := xml.Marshal(Sqlconfig)
	if err != nil {
		return false
	}
	filePath := filepath.Join(global.ConfigsDir, "sqlconfig.xml")
	os.WriteFile(filePath, configXml, 0644)
	return true
}

func InitAdminAccount() {
	if SelectAdminUser(1) {
		return
	}
	//创建管理员
	var adminUsername string
	var adminPassword string
	var adminEmail string
	fmt.Println("[FeasOJ]创建管理员账户")
	fmt.Print("[FeasOJ]请输入管理员账号:")
	fmt.Scanln(&adminUsername)
	fmt.Print("[FeasOJ]请输入管理员密码:")
	fmt.Scanln(&adminPassword)
	fmt.Print("[FeasOJ]请输入管理员邮箱:")
	fmt.Scanln(&adminEmail)

	tokensecret := uuid.New().String()
	Register(adminUsername, account.EncryptPassword(adminPassword), adminEmail, tokensecret, 1)
}

func LoadSqlConfig() string {
	// 读取config.xml文件
	filePath := filepath.Join(global.ConfigsDir, "sqlconfig.xml")
	configFile, err := os.Open(filePath)
	if err != nil {
		return ""
	}
	var SqlConfig global.SqlConfig
	err = xml.NewDecoder(configFile).Decode(&SqlConfig)
	if err != nil {
		return ""
	}
	dsn := SqlConfig.DbUser + ":" + SqlConfig.DbPassword + "@tcp(" + SqlConfig.DbAddress + ")/" + SqlConfig.DbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	return dsn
}

// 连接数据库
func connectSql() *gorm.DB {
	dsn := LoadSqlConfig()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("[FeasOJ]数据库连接失败，请手动前往config.xml进行配置。")
		return nil
	}
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("[FeasOJ]获取通用数据库对象失败。")
		return nil
	}
	// 设置连接池避免数据库连接过多导致的性能问题
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Second * 6)
	return db
}

// 创建表
func InitSql() bool {
	filePath := filepath.Join(global.ConfigsDir, "sqlconfig.xml")
	//判断是否有config.xml文件，没有则输入
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		InputSqlInfo()
	}
	connectSql().AutoMigrate(&global.User{}, &global.Problem{}, &global.SubmitRecord{}, &global.Discussion{}, &global.Comment{}, &global.TestCase{}, &global.Reply{})
	InitAdminAccount()
	return true
}

// 注册用户添加至数据库
func Register(username, password, email, tokensecret string, role int) bool {
	time := time.Now()
	err := connectSql().Create(&global.User{Username: username, Password: password, Email: email, CreateAt: time, Role: role, TokenSecret: tokensecret, IsBan: false}).Error
	return err == nil
}

// 查询管理员用户
func SelectAdminUser(role int) bool {
	// role = 1表示管理员
	// 查询是否有role = 1，有则返回true，否则返回false
	var user global.User
	err := connectSql().Where("role = ?", role).First(&user).Error
	return err == nil
}

// 查询指定用户的password加密值
func SelectPassword(username string) string {
	// 查询用户
	var user global.User
	// 查询用户名或邮箱
	connectSql().Where("username = ?", username).First(&user)
	return user.Password
}

// 查询指定用户的tokensecret
func SelectTokenSecret(username string) string {
	// 查询用户
	var user global.User
	connectSql().Where("username = ?", username).First(&user)
	return user.TokenSecret
}

// 根据username查询指定用户的除了password和tokensecret之外的所有信息
func SelectUserInfo(username string) global.UserInfoRequest {
	var user global.UserInfoRequest
	connectSql().Table("users").Where("username = ?", username).
		First(&user)
	return user
}

// 根据email与username判断是否该用户已存在
func IsUserExist(username, email string) bool {
	if connectSql().Where("username = ?", username).
		First(&global.User{}).Error == nil || connectSql().Where("email = ?", email).
		First(&global.User{}).Error == nil {
		return true
	}
	return false
}

// 根据email修改密码
func UpdatePassword(email, newpassword string) bool {
	newpassword = account.EncryptPassword(newpassword)
	tokensecret := uuid.New().String()
	err := connectSql().Model(&global.User{}).
		Where("email = ?", email).Update("password", newpassword).
		Update("token_secret", tokensecret).Error
	return err == nil
}

// 更新数据库中用户的头像路径
func UpdateAvatar(username, avatarpath string) bool {
	err := connectSql().Model(&global.User{}).
		Where("username = ?", username).Update("avatar", avatarpath).Error
	return err == nil
}

// 获取Problem表中的所有数据
func SelectAllProblems() []global.Problem {
	var problems []global.Problem
	connectSql().Find(&problems)
	return problems
}

// 获取指定PID的题目除了Input_full_path Output_full_path外的所有信息
func SelectProblemInfo(pid string) global.ProblemInfoRequest {
	var problem global.ProblemInfoRequest
	connectSql().Table("Problems").Where("pid = ?", pid).First(&problem)
	return problem
}

// 倒序查询指定用户ID的30天内的提交题目记录
func SelectSubmitRecordsByUid(uid string) []global.SubmitRecord {
	var records []global.SubmitRecord
	connectSql().Where("uid = ?", uid).
		Where("time > ?", time.Now().Add(-30*24*time.Hour)).Order("time desc").Find(&records)
	return records
}

// 返回SubmitRecord表中30天内的记录
func SelectAllSubmitRecords() []global.SubmitRecord {
	var records []global.SubmitRecord
	connectSql().
		Where("time > ?", time.Now().Add(-30*24*time.Hour)).Order("time desc").Find(&records)
	return records
}

// 获取讨论列表
func SelectDiscussList(page int, itemsPerPage int) ([]global.DiscussRequest, int) {
	var discussRequests []global.DiscussRequest
	var total int64

	db := connectSql().Table("Discussions").
		Select("Discussions.Did, Discussions.Title, Users.Username, Discussions.Create_at").
		Joins("JOIN Users ON Discussions.Uid = Users.Uid").
		Order("Discussions.Create_at desc")

	db.Count(&total) // 获取总讨论数
	db.Offset((page - 1) * itemsPerPage).Limit(itemsPerPage).Find(&discussRequests)

	return discussRequests, int(total)
}

// 获取指定Did的讨论以及User表中发帖人的头像
func SelectDiscussionByDid(Did string) global.DiscsInfoRequest {
	var discussion global.DiscsInfoRequest
	connectSql().Table("Discussions").
		Select("Discussions.Did, Discussions.Title, Discussions.Content, Discussions.Create_at, Users.Uid,Users.Username, Users.Avatar").
		Joins("JOIN Users ON Discussions.Uid = Users.Uid").
		Where("Discussions.Did = ?", Did).First(&discussion)
	return discussion
}

// 添加讨论
func AddDiscussion(title, content string, uid int) bool {
	if title == "" || content == "" {
		return false
	}
	err := connectSql().Table("Discussions").Create(&global.Discussion{Uid: uid, Title: title, Content: content, Create_at: time.Now()}).Error
	return err == nil
}

// 删除讨论
func DelDiscussion(Did string) bool {
	err := connectSql().Table("Discussions").Where("Did = ?", Did).Delete(&global.Discussion{}).Error
	return err == nil
}

// 添加评论
func AddComment(content string, did, uid int) bool {
	return connectSql().Table("Comments").Create(&global.Comment{Did: did, Uid: uid, Content: content, Create_at: time.Now()}).Error == nil
}

// 获取指定讨论ID的所有评论信息
func GetCommentsByDid(Did int) []global.CommentRequest {
	var comments []global.CommentRequest
	connectSql().Table("Comments").
		Select("Comments.Cid, Comments.Did, Comments.Content, Comments.Create_at, Users.Uid,Users.Username, Users.Avatar").
		Joins("JOIN Users ON Comments.Uid = Users.Uid").
		Order("create_at desc").
		Where("Comments.Did = ?", Did).Find(&comments)
	return comments
}

// 删除指定评论
func DeleteCommentByCid(Cid int) bool {
	return connectSql().Table("Comments").Where("Cid = ?", Cid).Delete(&global.Comment{}).Error == nil
}

// 获取指定题目的测试样例
func SelectTestCasesByPid(pid string) []global.TestCaseRequest {
	var testCases []global.TestCaseRequest
	connectSql().Table("test_cases").Where("pid = ?", pid).Select("input_data,output_data").Find(&testCases)
	return testCases
}

// 获取指定题目所有信息
func SelectProblemTestCases(pid string) global.AdminProblemInfoRequest {
	var problem global.Problem
	var testCases []global.TestCaseRequest
	var result global.AdminProblemInfoRequest

	if err := connectSql().First(&problem, pid).Error; err != nil {
		return result
	}

	if err := connectSql().Table("test_cases").Where("pid = ?", pid).Select("input_data,output_data").Find(&testCases).Error; err != nil {
		return result
	}

	result = global.AdminProblemInfoRequest{
		Pid:         problem.Pid,
		Difficulty:  problem.Difficulty,
		Title:       problem.Title,
		Content:     problem.Content,
		Timelimit:   problem.Timelimit,
		Memorylimit: problem.Memorylimit,
		Input:       problem.Input,
		Output:      problem.Output,
		TestCases:   testCases,
	}

	return result
}

// 更新题目信息
func UpdateProblem(req global.AdminProblemInfoRequest) error {
	// 更新题目表
	problem := global.Problem{
		Pid:         req.Pid,
		Difficulty:  req.Difficulty,
		Title:       req.Title,
		Content:     req.Content,
		Timelimit:   req.Timelimit,
		Memorylimit: req.Memorylimit,
		Input:       req.Input,
		Output:      req.Output,
	}
	if err := connectSql().Save(&problem).Error; err != nil {
		return err
	}

	// 获取该题目的测试样例
	var existingTestCases []global.TestCase
	if err := connectSql().Where("pid = ?", req.Pid).Find(&existingTestCases).Error; err != nil {
		return err
	}

	// 找出前端不存在的测试样例，并将其从数据库中删除
	existingTestCaseMap := make(map[string]global.TestCase)
	for _, testCase := range existingTestCases {
		existingTestCaseMap[testCase.InputData] = testCase
	}

	for _, testCase := range req.TestCases {
		delete(existingTestCaseMap, testCase.InputData)
	}

	for _, testCase := range existingTestCaseMap {
		if err := connectSql().Delete(&testCase).Error; err != nil {
			return err
		}
	}

	// 更新或添加新的测试样例
	for _, testCase := range req.TestCases {
		var existingTestCase global.TestCase
		if err := connectSql().Where("pid = ? AND input_data = ?", req.Pid, testCase.InputData).First(&existingTestCase).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				// 如果测试样例不存在，则创建新的样例
				newTestCase := global.TestCase{
					Pid:        req.Pid,
					InputData:  testCase.InputData,
					OutputData: testCase.OutputData,
				}
				if err := connectSql().Create(&newTestCase).Error; err != nil {
					return err
				}
			} else {
				return err
			}
		} else {
			// 如果测试样例存在，则更新该样例
			existingTestCase.OutputData = testCase.OutputData
			if err := connectSql().Save(&existingTestCase).Error; err != nil {
				return err
			}
		}
	}
	return nil
}

// 删除题目及其所有测试样例
func DeleteProblemAllInfo(pid int) bool {
	if connectSql().Table("problems").Where("pid = ?", pid).Delete(&global.Problem{}).Error != nil {
		return false
	}

	if connectSql().Table("test_cases").Where("pid = ?", pid).Delete(&global.TestCase{}).Error != nil {
		return false
	}

	return true
}

// 添加提交记录
func AddSubmitRecord(Uid, Pid int, Result, Language string) bool {
	err := connectSql().Table("submit_records").Create(&global.SubmitRecord{Uid: Uid, Pid: Pid, Result: Result, Time: time.Now(), Language: Language})
	return err == nil
}

// 修改提交记录状态
func ModifyJudgeStatus(Uid, Pid int, Result string) bool {
	// 将result为Running...的记录修改为返回状态
	err := connectSql().Table("submit_records").Where("uid = ? AND pid = ? AND result = ?", Uid, Pid, "Running...").Update("result", Result)
	return err == nil
}

// 获取所有用户信息
func GetAllUsersInfo() []global.UserInfoRequest {
	var usersInfo []global.UserInfoRequest
	connectSql().Table("users").Find(&usersInfo)
	return usersInfo
}
