package sql

import (
	"src/global"
	"src/utils"
	"time"

	"github.com/google/uuid"
)

// 用户注册
func Register(username, password, email, tokensecret string, role int) bool {
	time := time.Now()
	err := utils.ConnectSql().Create(&global.User{Username: username, Password: password, Email: email, CreateAt: time, Role: role, TokenSecret: tokensecret, IsBan: false}).Error
	return err == nil
}

// 管理员更新用户信息
func UpdateUser(Uid int, field string, value interface{}) bool {
	return utils.ConnectSql().Table("users").Where("uid = ?", Uid).Update(field, value).Error == nil
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
	utils.ConnectSql().Table("users").Find(&usersInfo)
	return usersInfo
}

// 更新用户的头像路径
func UpdateAvatar(username, avatarpath string) bool {
	err := utils.ConnectSql().Model(&global.User{}).
		Where("username = ?", username).Update("avatar", avatarpath).Error
	return err == nil
}

// 更新个人简介
func UpdateSynopsis(username, synopsis string) bool {
	err := utils.ConnectSql().Model(&global.User{}).
		Where("username = ?", username).Update("synopsis", synopsis).Error
	return err == nil
}

// 根据email与username判断是否该用户已存在
func IsUserExist(username, email string) bool {
	if utils.ConnectSql().Where("username = ?", username).
		First(&global.User{}).Error == nil || utils.ConnectSql().Where("email = ?", email).
		First(&global.User{}).Error == nil {
		return true
	}
	return false
}

// 根据邮箱获取用户信息
func SelectUserByEmail(email string) global.User {
	var user global.User
	utils.ConnectSql().Where("email = ?", email).First(&user)
	return user
}

// 根据email修改密码
func UpdatePassword(email, newpassword string) bool {
	tokensecret := uuid.New().String()
	err := utils.ConnectSql().Model(&global.User{}).
		Where("email = ?", email).Update("password", newpassword).
		Update("token_secret", tokensecret).Error
	return err == nil
}

// 根据username查询指定用户的除了password和tokensecret之外的所有信息
func SelectUserInfo(username string) global.UserInfoRequest {
	var user global.UserInfoRequest
	utils.ConnectSql().Table("users").Where("username = ?", username).
		First(&user)
	return user
}

// 查询管理员用户
func SelectAdminUser(role int) bool {
	// role = 1表示管理员
	var user global.User
	err := utils.ConnectSql().Where("role = ?", role).First(&user).Error
	return err == nil
}
