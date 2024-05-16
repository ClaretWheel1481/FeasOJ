package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 数据库连接信息
type Config struct {
	DbName     string `xml:"dbname"`
	DbUser     string `xml:"dbuser"`
	DbPassword string `xml:"dbpassword"`
	DbAddress  string `xml:"dbaddress"`
	DbPort     string `xml:"dbport"`
}

// 用户表：uid, avartar, username, password, email, score, synopsis, submit_history, create_at
type User struct {
	Uid           int       `gorm:"comment:Uid;primaryKey;autoIncrement"`
	Avartar       string    `gorm:"comment:头像"`
	Username      string    `gorm:"comment:用户名;not null;unique"`
	Password      string    `gorm:"comment:密码;not null"`
	Email         string    `gorm:"comment:电子邮件;not null"`
	Synopsis      string    `gorm:"comment:简介"`
	SubmitHistory string    `gorm:"comment:提交记录"`
	CreateAt      time.Time `gorm:"comment:创建时间;not null"`
	Role          int       `gorm:"comment:角色;not null"` // 0: 普通用户, 1: 管理员
}

// 题目表: pid, title, content, time_limit, memory_limit, input, output, contest, submit_history
// contest为0则不属于任何比赛
type Problem struct {
	Pid            int    `gorm:"comment:Pid;primaryKey;autoIncrement"`
	Title          string `gorm:"comment:题目标题;not null"`
	Content        string `gorm:"comment:题目详细;not null"`
	Time_limit     int    `gorm:"comment:运行时间限制;not null"`
	Memory_limit   int    `gorm:"comment:内存大小限制;not null"`
	Input          string `gorm:"comment:输入样例;not null"`
	Output         string `gorm:"comment:输出样例;not null"`
	Contestid      int    `gorm:"comment:隶属竞赛;not null"` //0为不属于任何竞赛，1...32767表示隶属的竞赛号
	Submit_history string `gorm:"comment:提交记录"`
}

func writeConfig(dbName, dbUser, dbPassword, dbAddress string) error {
	// 创建配置文件
	config := Config{
		DbName:     dbName,
		DbUser:     dbUser,
		DbPassword: dbPassword,
		DbAddress:  dbAddress,
	}
	configXml, err := xml.Marshal(config)
	if err != nil {
		return err
	}
	err = os.WriteFile("sqlconfig.xml", configXml, 0644)
	return err
}

func inputSqlInfo() bool {
	fmt.Println("[FeasOJ]请输入数据库连接信息：")
	fmt.Print("数据库名：")
	var dbName string
	fmt.Scanln(&dbName)
	fmt.Print("数据库用户名：")
	var dbUser string
	fmt.Scanln(&dbUser)
	fmt.Print("数据库密码：")
	var dbPassword string
	fmt.Scanln(&dbPassword)
	fmt.Print("数据库地址（加上端口）：")
	var dbAddress string
	fmt.Scanln(&dbAddress)
	fmt.Println("[FeasOJ]正在保存数据库连接信息...")
	// 将连接信息作为xml保存到目录
	err := writeConfig(dbName, dbUser, dbPassword, dbAddress)
	if err != nil {
		fmt.Println("[FeasOJ]保存失败，请检查权限。")
		return false
	}
	return true
}

func initAdminAccount() {
	if selectAdminUser(1) {
		fmt.Println("[FeasOJ]管理员已存在，退出创建。")
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
	register(adminUsername, encryptPassword(adminPassword), adminEmail, 1)
}

func loadConfig() string {
	// 读取config.xml文件
	configFile, err := os.Open("sqlconfig.xml")
	if err != nil {
		return ""
	}
	defer configFile.Close()
	var config Config
	err = xml.NewDecoder(configFile).Decode(&config)
	if err != nil {
		return ""
	}
	dbName := config.DbName
	dbUser := config.DbUser
	dbPassword := config.DbPassword
	dbAddress := config.DbAddress
	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbAddress + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	return dsn
}

// 连接数据库、创建表
func initSql() bool {
	//判断是否有config.xml文件，没有则输入
	if _, err := os.Stat("sqlconfig.xml"); os.IsNotExist(err) {
		inputSqlInfo()
	}
	dsn := loadConfig()
	fmt.Println("[FeasOJ]连接数据库中...")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("[FeasOJ]数据库连接失败，请手动前往config.xml进行配置。")
	}
	fmt.Println("[FeasOJ]数据库连接成功。")
	fmt.Println("[FeasOJ]创建数据表中...")
	db.AutoMigrate(&User{}, &Problem{})
	fmt.Println("[FeasOJ]创建数据表成功。")
	initAdminAccount()
	return true
}

// 注册用户添加至数据库
func register(username, password, email string, role int) bool {
	dsn := loadConfig()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("[FeasOJ]数据库连接失败，请手动前往config.xml进行配置。")
	}
	time := time.Now()
	err = db.Create(&User{Username: username, Password: password, Email: email, CreateAt: time, Role: role}).Error
	if err != nil {
		fmt.Println("[FeasOJ]添加用户失败，请检查用户名是否重复。")
		return false
	}
	fmt.Println("[FeasOJ]添加用户成功。")
	return true
}

// 查询管理员用户
func selectAdminUser(role int) bool {
	// role = 1表示管理员
	dsn := loadConfig()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("[FeasOJ]数据库连接失败，请手动前往config.xml进行配置。")
	}
	// 查询是否有role = 1，有则返回true，否则返回false
	var user User
	err = db.Where("role = ?", role).First(&user).Error
	return err == nil
}

// 查询指定用户的password值
func selectPassword(username string) string {
	// 查询用户
	dsn := loadConfig()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("[FeasOJ]数据库连接失败，请手动前往config.xml进行配置。")
	}
	var user User
	db.Where("username = ?", username).First(&user)
	return user.Password
}
