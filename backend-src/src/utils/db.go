package utils

import (
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"

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
	dsn := SqlConfig.DbUser + ":" + SqlConfig.DbPassword + "@tcp(" + SqlConfig.DbAddress + ")/" + SqlConfig.DbName + "?charset=utf8mb4&parseTime=True&loc=Asia%2FShanghai"
	return dsn
}

func InitAdminAccount() {
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

	Register(adminUsername, EncryptPassword(adminPassword), adminEmail, uuid.New().String(), 1)
}

// 创建表
func InitSql() bool {
	filePath := filepath.Join(global.ConfigsDir, "sqlconfig.xml")
	//判断是否有config.xml文件，没有则输入
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		InputSqlInfo()
	}
	ConnectSql().AutoMigrate(&global.User{}, &global.Problem{}, &global.SubmitRecord{}, &global.Discussion{}, &global.Comment{}, &global.TestCase{}, &global.Competition{})
	return true
}

// 返回数据库连接对象
func ConnectSql() *gorm.DB {
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

// 用户注册
func Register(username, password, email, tokensecret string, role int) bool {
	time := time.Now()
	err := ConnectSql().Create(&global.User{Username: username, Password: password, Email: email, CreateAt: time, Role: role, TokenSecret: tokensecret, IsBan: false}).Error
	return err == nil
}

// 管理员获取竞赛信息
func SelectCompetitionInfoAdmin() []global.AdminCompetitionInfoRequest {
	var competitions []global.AdminCompetitionInfoRequest
	ConnectSql().Table("competitions").Find(&competitions)
	return competitions
}

// 管理员获取指定竞赛ID信息
func SelectCompetitionInfoAdminByCid(Cid int) global.AdminCompetitionInfoRequest {
	var competition global.AdminCompetitionInfoRequest
	ConnectSql().Table("competitions").Where("cid = ?", Cid).Find(&competition)
	return competition
}

// 用户获取竞赛信息
func SelectCompetitionInfo() []global.CompetitionRequest {
	var competition []global.CompetitionRequest
	ConnectSql().Table("competitions").Find(&competition)
	return competition
}

// 查询管理员用户
func SelectAdminUser(role int) bool {
	// role = 1表示管理员
	var user global.User
	err := ConnectSql().Where("role = ?", role).First(&user).Error
	return err == nil
}

// 获取用户信息
func SelectUser(username string) global.User {
	var user global.User
	ConnectSql().Where("username = ?", username).First(&user)
	return user
}

// 根据email修改密码
func UpdatePassword(email, newpassword string) bool {
	newpassword = EncryptPassword(newpassword)
	tokensecret := uuid.New().String()
	err := ConnectSql().Model(&global.User{}).
		Where("email = ?", email).Update("password", newpassword).
		Update("token_secret", tokensecret).Error
	return err == nil
}

// 管理员更新用户信息
func UpdateUser(Uid int, field string, value interface{}) bool {
	return ConnectSql().Table("users").Where("uid = ?", Uid).Update(field, value).Error == nil
}

// 封禁用户
func BanUser(Uid int) bool {
	return UpdateUser(Uid, "is_ban", true)
}

// 解除封禁
func UnbanUser(Uid int) bool {
	return UpdateUser(Uid, "is_ban", false)
}

// 晋升为管理员
func PromoteToAdmin(Uid int) bool {
	return UpdateUser(Uid, "role", 1)
}

// 降级为普通用户
func DemoteToUser(Uid int) bool {
	return UpdateUser(Uid, "role", 0)
}

// 管理员获取所有用户信息
func SelectAllUsersInfo() []global.UserInfoRequest {
	var usersInfo []global.UserInfoRequest
	ConnectSql().Table("users").Find(&usersInfo)
	return usersInfo
}

// 更新用户的头像路径
func UpdateAvatar(username, avatarpath string) bool {
	err := ConnectSql().Model(&global.User{}).
		Where("username = ?", username).Update("avatar", avatarpath).Error
	return err == nil
}

// 更新个人简介
func UpdateSynopsis(username, synopsis string) bool {
	err := ConnectSql().Model(&global.User{}).
		Where("username = ?", username).Update("synopsis", synopsis).Error
	return err == nil
}

// 根据username查询指定用户的除了password和tokensecret之外的所有信息
func SelectUserInfo(username string) global.UserInfoRequest {
	var user global.UserInfoRequest
	ConnectSql().Table("users").Where("username = ?", username).
		First(&user)
	return user
}

// 根据email与username判断是否该用户已存在
func IsUserExist(username, email string) bool {
	if ConnectSql().Where("username = ?", username).
		First(&global.User{}).Error == nil || ConnectSql().Where("email = ?", email).
		First(&global.User{}).Error == nil {
		return true
	}
	return false
}

// 根据邮箱获取信息
func SelectUserByEmail(email string) global.User {
	var user global.User
	ConnectSql().Where("email = ?", email).First(&user)
	return user
}
