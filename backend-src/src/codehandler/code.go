package codehandler

import (
	"log"
	"src/utils/sql"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis"
)

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
			log.Panic(err)
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
			log.Panic(err)
		}
		pidInt, err := strconv.Atoi(pid)
		if err != nil {
			log.Panic(err)
		}
		sql.ModifyJudgeStatus(uidInt, pidInt, str)
	}
}
