package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func inputSqlInfo() bool {
	var Sqlconfig SqlConfig
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
	filePath := filepath.Join(configsDir, "sqlconfig.xml")
	os.WriteFile(filePath, configXml, 0644)
	return true
}

func initAdminAccount() {
	if selectAdminUser(1) {
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
	register(adminUsername, encryptPassword(adminPassword), adminEmail, tokensecret, 1)
}

func loadSqlConfig() string {
	// 读取config.xml文件
	filePath := filepath.Join(configsDir, "sqlconfig.xml")
	configFile, err := os.Open(filePath)
	if err != nil {
		return ""
	}
	var SqlConfig SqlConfig
	err = xml.NewDecoder(configFile).Decode(&SqlConfig)
	if err != nil {
		return ""
	}
	dsn := SqlConfig.DbUser + ":" + SqlConfig.DbPassword + "@tcp(" + SqlConfig.DbAddress + ")/" + SqlConfig.DbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	return dsn
}

// 连接数据库
func connectSql() *gorm.DB {
	dsn := loadSqlConfig()
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
func initSql() bool {
	filePath := filepath.Join(configsDir, "sqlconfig.xml")
	//判断是否有config.xml文件，没有则输入
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		inputSqlInfo()
	}
	connectSql().AutoMigrate(&User{}, &Problem{}, &SubmitRecord{}, &Discussion{}, &Comment{}, &TestCase{}, &Reply{})
	initAdminAccount()
	return true
}

// 注册用户添加至数据库
func register(username, password, email, tokensecret string, role int) bool {
	time := time.Now()
	err := connectSql().Create(&User{Username: username, Password: password, Email: email, CreateAt: time, Role: role, TokenSecret: tokensecret}).Error
	return err == nil
}

// 查询管理员用户
func selectAdminUser(role int) bool {
	// role = 1表示管理员
	// 查询是否有role = 1，有则返回true，否则返回false
	var user User
	err := connectSql().Where("role = ?", role).First(&user).Error
	return err == nil
}

// 查询指定用户的password加密值
func selectPassword(username string) string {
	// 查询用户
	var user User
	// 查询用户名或邮箱
	connectSql().Where("username = ?", username).First(&user)
	return user.Password
}

// 查询指定用户的tokensecret
func selectTokenSecret(username string) string {
	// 查询用户
	var user User
	connectSql().Where("username = ?", username).First(&user)
	return user.TokenSecret
}

// 根据username查询指定用户的除了password和tokensecret之外的所有信息
func selectUserInfo(username string) userInfoRequest {
	var user userInfoRequest
	connectSql().Table("users").Where("username = ?", username).
		First(&user)
	return user
}

// 根据email与username判断是否该用户已存在
func isUserExist(username, email string) bool {
	if connectSql().Where("username = ?", username).
		First(&User{}).Error == nil || connectSql().Where("email = ?", email).
		First(&User{}).Error == nil {
		return true
	}
	return false
}

// 根据email修改密码
func updatePassword(email, newpassword string) bool {
	newpassword = encryptPassword(newpassword)
	tokensecret := uuid.New().String()
	err := connectSql().Model(&User{}).
		Where("email = ?", email).Update("password", newpassword).
		Update("token_secret", tokensecret).Error
	return err == nil
}

// 更新数据库中用户的头像路径
func updateAvatar(username, avatarpath string) bool {
	err := connectSql().Model(&User{}).
		Where("username = ?", username).Update("avatar", avatarpath).Error
	return err == nil
}

// 获取Problem表中的所有数据
func selectAllProblems() []Problem {
	var problems []Problem
	connectSql().Find(&problems)
	return problems
}

// 获取指定PID的题目除了Input_full_path Output_full_path外的所有信息
func selectProblemInfo(pid string) problemInfoRequest {
	var problem problemInfoRequest
	connectSql().Table("Problems").Where("pid = ?", pid).First(&problem)
	return problem
}

// 倒序查询指定用户ID的30天内的提交题目记录
func selectSubmitRecordsByUid(uid string) []SubmitRecord {
	var records []SubmitRecord
	connectSql().Where("uid = ?", uid).
		Where("time > ?", time.Now().Add(-30*24*time.Hour)).Order("time desc").Find(&records)
	return records
}

// 返回SubmitRecord表中30天内的记录
func selectAllSubmitRecords() []SubmitRecord {
	var records []SubmitRecord
	connectSql().
		Where("time > ?", time.Now().Add(-30*24*time.Hour)).Order("time desc").Find(&records)
	return records
}

// 获取讨论列表
func selectDiscussList() []discussRequest {
	var discussRequests []discussRequest
	connectSql().Table("Discussions").
		Select("Discussions.Did,Discussions.Title, Users.Username, Discussions.Create_at").
		Joins("JOIN Users ON Discussions.Uid = Users.Uid").
		Order("Discussions.Create_at desc").
		Find(&discussRequests)
	return discussRequests
}

// 获取指定Did的讨论以及User表中发帖人的头像
func selectDiscussionByDid(Did string) discsInfoRequest {
	var discussion discsInfoRequest
	connectSql().Table("Discussions").
		Select("Discussions.Did, Discussions.Title, Discussions.Content, Discussions.Create_at, Users.Uid,Users.Username, Users.Avatar").
		Joins("JOIN Users ON Discussions.Uid = Users.Uid").
		Where("Discussions.Did = ?", Did).First(&discussion)
	return discussion
}

// 添加讨论
func addDiscussion(title, content string, uid int) bool {
	if title == "" || content == "" {
		return false
	}
	err := connectSql().Table("Discussions").Create(&Discussion{Uid: uid, Title: title, Content: content, Create_at: time.Now()}).Error
	return err == nil
}

// 删除讨论
func delDiscussion(Did string) bool {
	err := connectSql().Table("Discussions").Where("Did = ?", Did).Delete(&Discussion{}).Error
	return err == nil
}

// 获取指定题目所有输入输出样例
func selectProblemTestCases(pid string) adminProblemInfoRequest {
	var problem Problem
	var testCases []TestCaseRequest
	var result adminProblemInfoRequest

	if err := connectSql().First(&problem, pid).Error; err != nil {
		return result
	}

	if err := connectSql().Table("test_cases").Where("pid = ?", pid).Select("input_data,output_data").Find(&testCases).Error; err != nil {
		return result
	}

	result = adminProblemInfoRequest{
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
func updateProblem(req adminProblemInfoRequest) error {
	// 更新题目表
	problem := Problem{
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
	var existingTestCases []TestCase
	if err := connectSql().Where("pid = ?", req.Pid).Find(&existingTestCases).Error; err != nil {
		return err
	}

	// 找出前端不存在的测试样例，并将其从数据库中删除
	existingTestCaseMap := make(map[string]TestCase)
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
		var existingTestCase TestCase
		if err := connectSql().Where("pid = ? AND input_data = ?", req.Pid, testCase.InputData).First(&existingTestCase).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				// 如果测试样例不存在，则创建新的样例
				newTestCase := TestCase{
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
func deleteProblemAllInfo(pid int) bool {
	if connectSql().Table("problems").Where("pid = ?", pid).Delete(&Problem{}).Error != nil {
		return false
	}

	if connectSql().Table("test_cases").Where("pid = ?", pid).Delete(&TestCase{}).Error != nil {
		return false
	}

	return true
}

// 获取指定讨论ID的所有评论信息
func getCommentsByDid(Did int) []CommentRequest {
	var comments []CommentRequest
	connectSql().Table("Comments").
		Select("Comments.Cid, Comments.Did, Comments.Content, Comments.Create_at, Users.Uid,Users.Username, Users.Avatar").
		Joins("JOIN Users ON Comments.Uid = Users.Uid").
		Order("create_at desc").
		Where("Comments.Did = ?", Did).Find(&comments)
	return comments
}
