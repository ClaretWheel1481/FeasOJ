package sql

import (
	"src/global"
	"src/utils"
)

// 管理员获取竞赛信息
func SelectCompetitionInfoAdmin() []global.AdminCompetitionInfoRequest {
	var competitions []global.AdminCompetitionInfoRequest
	utils.ConnectSql().Table("competitions").Find(&competitions)
	return competitions
}

// 管理员获取指定竞赛ID信息
func SelectCompetitionInfoAdminByCid(Cid int) global.AdminCompetitionInfoRequest {
	var competition global.AdminCompetitionInfoRequest
	utils.ConnectSql().Table("competitions").Where("contest_id = ?", Cid).Find(&competition)
	return competition
}

// 用户获取竞赛信息
func SelectCompetitionInfo() []global.CompetitionRequest {
	var competition []global.CompetitionRequest
	utils.ConnectSql().Table("competitions").Where("is_visible = ?", true).Order("start_at DESC").Find(&competition)
	return competition
}
