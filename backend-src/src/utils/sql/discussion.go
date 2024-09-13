package sql

import (
	"src/global"
	"src/utils"
	"time"
)

// 获取讨论列表
func SelectDiscussList(page int, itemsPerPage int) ([]global.DiscussRequest, int) {
	var discussRequests []global.DiscussRequest
	var total int64

	db := utils.ConnectSql().Table("Discussions").
		Select("Discussions.Did, Discussions.Title, Users.Username, Discussions.Create_at").
		Joins("JOIN Users ON Discussions.Uid = Users.Uid").
		Order("Discussions.Create_at desc")

	db.Count(&total) // 获取总讨论数
	db.Offset((page - 1) * itemsPerPage).Limit(itemsPerPage).Find(&discussRequests)

	return discussRequests, int(total)
}

// 获取指定Did讨论及User表中发帖人的头像
func SelectDiscussionByDid(Did string) global.DiscsInfoRequest {
	var discussion global.DiscsInfoRequest
	utils.ConnectSql().Table("Discussions").
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
	err := utils.ConnectSql().Table("Discussions").Create(&global.Discussion{Uid: uid, Title: title, Content: content, Create_at: time.Now()}).Error
	return err == nil
}

// 删除讨论
func DelDiscussion(Did string) bool {
	err := utils.ConnectSql().Table("Discussions").Where("Did = ?", Did).Delete(&global.Discussion{}).Error
	return err == nil
}

// 添加评论
func AddComment(content string, did, uid int) bool {
	return utils.ConnectSql().Table("Comments").Create(&global.Comment{Did: did, Uid: uid, Content: content, Create_at: time.Now()}).Error == nil
}

// 获取指定讨论ID的所有评论信息
func SelectCommentsByDid(Did int) []global.CommentRequest {
	var comments []global.CommentRequest
	utils.ConnectSql().Table("Comments").
		Select("Comments.Cid, Comments.Did, Comments.Content, Comments.Create_at, Users.Uid,Users.Username, Users.Avatar").
		Joins("JOIN Users ON Comments.Uid = Users.Uid").
		Order("create_at desc").
		Where("Comments.Did = ?", Did).Find(&comments)
	return comments
}

// 删除指定评论
func DeleteCommentByCid(Cid int) bool {
	return utils.ConnectSql().Table("Comments").Where("Cid = ?", Cid).Delete(&global.Comment{}).Error == nil
}
