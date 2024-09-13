package sql

import (
	"src/global"
	"src/utils"
	"time"
)

// 倒序查询指定用户ID的30天内的提交题目记录
func SelectSubmitRecordsByUid(uid int) []global.SubmitRecord {
	var records []global.SubmitRecord
	utils.ConnectSql().Where("uid = ?", uid).
		Where("time > ?", time.Now().Add(-30*24*time.Hour)).Order("time desc").Find(&records)
	return records
}

// 返回SubmitRecord表中30天内的记录
func SelectAllSubmitRecords() []global.SubmitRecord {
	var records []global.SubmitRecord
	utils.ConnectSql().
		Where("time > ?", time.Now().Add(-30*24*time.Hour)).Order("time desc").Find(&records)
	return records
}

// 添加提交记录
func AddSubmitRecord(Uid, Pid int, Result, Language, Username string) bool {
	err := utils.ConnectSql().Table("submit_records").Create(&global.SubmitRecord{Uid: Uid, Pid: Pid, Username: Username, Result: Result, Time: time.Now(), Language: Language})
	return err == nil
}

// 修改提交记录状态
func ModifyJudgeStatus(Uid, Pid int, Result string) bool {
	// 将result为Running...的记录修改为返回状态
	err := utils.ConnectSql().Table("submit_records").Where("uid = ? AND pid = ? AND result = ?", Uid, Pid, "Running...").Update("result", Result)
	return err == nil
}
