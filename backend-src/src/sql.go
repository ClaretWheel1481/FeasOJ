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

// 数据库连接信息
type SqlConfig struct {
	DbName     string `xml:"dbname"`
	DbUser     string `xml:"dbuser"`
	DbPassword string `xml:"dbpassword"`
	DbAddress  string `xml:"dbaddress"`
	DbPort     string `xml:"dbport"`
}

// 用户表：uid, avatar, username, password, email, score, synopsis, submit_history, create_at
type User struct {
	Uid         int       `gorm:"comment:Uid;primaryKey;autoIncrement"`
	Avatar      string    `gorm:"comment:头像存放路径"`
	Username    string    `gorm:"comment:用户名;not null;unique"`
	Password    string    `gorm:"comment:密码;not null"`
	Email       string    `gorm:"comment:电子邮件;not null"`
	Synopsis    string    `gorm:"comment:简介"`
	Score       int       `gorm:"comment:分数"`
	CreateAt    time.Time `gorm:"comment:创建时间;not null"`
	Role        int       `gorm:"comment:角色;not null"` // 0: 普通用户, 1: 管理员
	TokenSecret string    `gorm:"comment:token密钥;not null"`
}

// 题目表: pid, title, content, time_limit, memory_limit, input, output, contest, submit_history
// contest为0则不属于任何比赛
type Problem struct {
	Pid              int    `gorm:"comment:Pid;primaryKey;autoIncrement"`
	Title            string `gorm:"comment:题目标题;not null"`
	Content          string `gorm:"comment:题目详细;not null"`
	Timelimit        int    `gorm:"comment:运行时间限制;not null"`
	Memorylimit      int    `gorm:"comment:内存大小限制;not null"`
	Input            string `gorm:"comment:输入样例;not null"`
	Output           string `gorm:"comment:输出样例;not null"`
	Contestid        int    `gorm:"comment:隶属竞赛;not null"` //0为不属于任何竞赛，1...32767表示隶属的竞赛号
	Input_full_path  string `gorm:"comment:输入样例完整路径"`
	Output_full_path string `gorm:"comment:输出样例完整路径"`
}

// 总提交记录表: Pid,Uid,Result,Time，Language
type SubmitRecord struct {
	Pid      int    `gorm:"comment:Pid;primaryKey"`
	Uid      int    `gorm:"comment:Uid;not null"`
	Result   string `gorm:"comment:结果"`
	Time     int    `gorm:"comment:时间"`
	Language string `gorm:"comment:语言"`
}

// 竞赛表: Contestid, Title, Start_time, End_time, Description, Pid
type Contest struct {
	Contestid   int    `gorm:"comment:Contestid;primaryKey;autoIncrement"`
	Title       string `gorm:"comment:竞赛标题"`
	Start_time  string `gorm:"comment:开始时间"`
	End_time    string `gorm:"comment:结束时间"`
	Description string `gorm:"comment:竞赛描述"`
	Pid         int    `gorm:"comment:题目"`
}

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
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Second * 10)
	return db
}

// 创建表
func initSql() bool {
	filePath := filepath.Join(configsDir, "sqlconfig.xml")
	//判断是否有config.xml文件，没有则输入
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		inputSqlInfo()
	}
	connectSql().AutoMigrate(&User{}, &Problem{}, &SubmitRecord{}, &Contest{})
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
func selectUserInfo(username string) User {
	var user User
	connectSql().Where("username = ?", username).First(&user)
	user.Password = ""
	user.TokenSecret = ""
	return user
}

// 根据email与username判断是否该用户已存在
func isUserExist(username, email string) bool {
	if connectSql().Where("username = ?", username).First(&User{}).Error == nil || connectSql().Where("email = ?", email).First(&User{}).Error == nil {
		return true
	}
	return false
}

// 根据email修改密码
func updatePassword(email, newpassword string) bool {
	newpassword = encryptPassword(newpassword)
	tokensecret := uuid.New().String()
	err := connectSql().Model(&User{}).Where("email = ?", email).Update("password", newpassword).Update("token_secret", tokensecret).Error
	return err == nil
}

// 返回SubmitRecord表中所有记录
func selectAllSubmitRecords() []SubmitRecord {
	var records []SubmitRecord
	connectSql().Find(&records)
	return records
}

// 更新数据库中用户的头像路径
func updateAvatar(username, avatarpath string) bool {
	err := connectSql().Model(&User{}).Where("username = ?", username).Update("avatar", avatarpath).Error
	return err == nil
}

// 获取Problem表中的所有数据
func selectAllProblems() []Problem {
	var problems []Problem
	connectSql().Find(&problems)
	return problems
}
