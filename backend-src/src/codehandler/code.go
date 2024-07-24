package codehandler

import (
	"fmt"
	"os"
	"path/filepath"
	"src/global"
	"src/utils"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis"
)

// 初始化文件夹
func InitCodeFolder() {
	// TODO:每次编译前需要修改为currentDir，debug时用parentDir
	global.CodeDir = filepath.Join(global.ParentDir, "/codefiles")
	// 如果没找到codefiles，则创建codefiles文件夹
	if _, err := os.Stat(global.CodeDir); os.IsNotExist(err) {
		os.Mkdir(global.CodeDir, 0755)
	}
}

// 实时处理Redis任务队列中的任务
func ProcessJudgeTasks(rdb *redis.Client) {
	for {
		// 从Redis任务队列中取出一个任务
		task, err := rdb.LPop("judgeTask").Result()
		if err == redis.Nil {
			// 如果队列为空，等待一段时间后再检查
			time.Sleep(2 * time.Second)
			continue
		} else if err != nil {
			fmt.Println(err)
			continue
		}

		// 执行任务
		StartContainer()
		str := CompileAndRun(task)
		// 将结果从数据库中修改
		parts := strings.Split(task, "_")
		uid := parts[0]
		pid := strings.Split(parts[1], ".")[0]
		uidInt, err := strconv.Atoi(uid)
		if err != nil {
			fmt.Println(err)
		}
		pidInt, err := strconv.Atoi(pid)
		if err != nil {
			fmt.Println(err)
		}
		utils.ModifyJudgeStatus(uidInt, pidInt, str)
		TerminateContainer(global.ContainerID)
	}
}
