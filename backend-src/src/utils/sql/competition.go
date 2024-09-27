package sql

import (
	"src/global"
	"src/utils"
)

// 用户获取竞赛信息
func SelectCompetitionInfo() []global.CompetitionRequest {
	var competition []global.CompetitionRequest
	utils.ConnectSql().Table("competitions").Where("is_visible = ?", true).Order("start_at DESC").Find(&competition)
	return competition
}

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

// 管理员删除竞赛
func DeleteCompetition(Cid int) bool {
	return utils.ConnectSql().Table("competitions").Where("contest_id = ?", Cid).Delete(&global.CompetitionRequest{}) != nil
}

// 管理员更新/添加竞赛
func UpdateCompetition(req global.AdminCompetitionInfoRequest) error {
	if err := utils.ConnectSql().Table("competitions").Where("contest_id = ?", req.ContestID).Save(&req).Error; err != nil {
		return err
	}
	return nil
}
