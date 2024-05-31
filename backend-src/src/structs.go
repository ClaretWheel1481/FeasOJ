package main

import "time"

// MySQL数据库连接信息
type SqlConfig struct {
	DbName     string `xml:"dbname"`
	DbUser     string `xml:"dbuser"`
	DbPassword string `xml:"dbpassword"`
	DbAddress  string `xml:"dbaddress"`
	DbPort     string `xml:"dbport"`
}

// Redis数据连接信息
type redisConfig struct {
	Address  string `xml:"address"`
	Password string `xml:"password"`
}

// 邮件服务连接信息
type mailConfig struct {
	Host string `xml:"host"`
	Port int    `xml:"port"`
	User string `xml:"user"`
	Pass string `xml:"pass"`
}

// 注册请求体
type RegisterRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Vcode    string `json:"vcode"`
}

// 修改密码请求体
type updatePasswordRequest struct {
	Email       string `json:"email"`
	NewPassword string `json:"newpassword"`
	Vcode       string `json:"vcode"`
}

// 用户表：uid, avatar, username, password, email, score, synopsis, submit_history, create_at
type User struct {
	Uid         int       `gorm:"comment:Uid;primaryKey;autoIncrement"`
	Avatar      string    `gorm:"comment:头像存放路径"`
	Username    string    `gorm:"comment:用户名;not null;unique"`
	Password    string    `gorm:"comment:密码;not null"`
	Email       string    `gorm:"comment:电子邮件;not null"`
	Synopsis    string    `gorm:"comment:简介"`
	Score       int       `gorm:"comment:分数"`
	CreateAt    time.Time `gorm:"comment:创建时间;not null"`
	Role        int       `gorm:"comment:角色;not null"` // 0: 普通用户, 1: 管理员
	TokenSecret string    `gorm:"comment:token密钥;not null"`
}

// 题目表: pid, title, content, time_limit, memory_limit, input, output, contest, submit_history
// contest为0则不属于任何比赛
type Problem struct {
	Pid              int    `gorm:"comment:Pid;primaryKey;autoIncrement"`
	Title            string `gorm:"comment:题目标题;not null"`
	Content          string `gorm:"comment:题目详细;not null"`
	Timelimit        int    `gorm:"comment:运行时间限制;not null"`
	Memorylimit      int    `gorm:"comment:内存大小限制;not null"`
	Input            string `gorm:"comment:输入样例;not null"`
	Output           string `gorm:"comment:输出样例;not null"`
	Cid              int    `gorm:"comment:隶属竞赛;not null"` //0为不属于任何竞赛，1...32767表示隶属的竞赛号
	Input_full_path  string `gorm:"comment:输入样例完整路径"`
	Output_full_path string `gorm:"comment:输出样例完整路径"`
}

// 总提交记录表: Pid,Uid,Result,Time，Language
type SubmitRecord struct {
	Pid      int    `gorm:"comment:Pid;primaryKey"`
	Uid      int    `gorm:"comment:Uid;not null"`
	Result   string `gorm:"comment:结果"`
	Time     int    `gorm:"comment:时间"`
	Language string `gorm:"comment:语言"`
}

// 竞赛表: Contestid, Title, Start_time, End_time, Description, Pid
type Contest struct {
	Cid         int    `gorm:"comment:Cid;primaryKey;autoIncrement"`
	Title       string `gorm:"comment:竞赛标题"`
	Start_time  string `gorm:"comment:开始时间"`
	End_time    string `gorm:"comment:结束时间"`
	Description string `gorm:"comment:竞赛描述"`
}

// 讨论帖子表: Tid,Title,Content,Uid,Create_at
type Discuss struct {
	Tid       int    `gorm:"comment:Tid;primaryKey;autoIncrement"`
	Title     string `gorm:"comment:标题"`
	Content   string `gorm:"comment:内容"`
	Uid       int    `gorm:"comment:用户"`
	Create_at string `gorm:"comment:创建时间"`
}

// 讨论评论表: Cid,Tid,Content,Uid,Create_at
type Comment struct {
	Cid       int    `gorm:"comment:Cid;primaryKey;autoIncrement"`
	Tid       int    `gorm:"comment:Tid"`
	Content   string `gorm:"comment:内容"`
	Uid       int    `gorm:"comment:用户"`
	Create_at string `gorm:"comment:创建时间"`
}
