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
	connectSql().AutoMigrate(&User{}, &Problem{}, &SubmitRecord{}, &Discussion{}, &Comment{})
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
		Select("Discussions.Tid,Discussions.Title, Users.Username, Discussions.Create_at").
		Joins("JOIN Users ON Discussions.Uid = Users.Uid").
		Find(&discussRequests)
	return discussRequests
}

// 获取指定Tid的讨论以及User表中发帖人的头像
func selectDiscussionByTid(tid string) discsInfoRequest {
	var discussion discsInfoRequest
	connectSql().Table("Discussions").
		Select("Discussions.Tid, Discussions.Title, Discussions.Content, Discussions.Create_at, Users.Username, Users.Avatar").
		Joins("JOIN Users ON Discussions.Uid = Users.Uid").
		Where("Discussions.Tid = ?", tid).First(&discussion)
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
