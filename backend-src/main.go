package backendsrc

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

// 连接数据库、创建表
// 用户表：uid, avartar, username, password, email, score, synopsis, submit_history
func initSql() {
	//连接MySQL数据库
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//创建表
	sqlCreateTable := `
	    CREATE TABLE IF NOT EXISTS user (
	        uid INT AUTO_INCREMENT PRIMARY KEY,
			avartar VARCHAR(255),
			username VARCHAR(24),
			password VARCHAR(32),
			email VARCHAR(255),
			score INT,
			synopsis VARCHAR(255),
			submit_history VARCHAR(255)
	    )
	`
	_, err = db.Exec(sqlCreateTable)
}

func main() {
	r := gin.Default()
	r.Run() // listen and serve on 0.0.0.0:8080
}
