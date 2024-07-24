package global

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
type RedisConfig struct {
	Address  string `xml:"address"`
	Password string `xml:"password"`
}

// 邮件服务连接信息
type MailConfig struct {
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
type UpdatePasswordRequest struct {
	Email       string `json:"email"`
	NewPassword string `json:"newpassword"`
	Vcode       string `json:"vcode"`
}

// 讨论列表请求体
type DiscussRequest struct {
	Did       string `json:"did"`
	Title     string `json:"title"`
	Username  string `json:"username"`
	Create_at string `json:"create_at"`
}

// 讨论帖子页请求体
type DiscsInfoRequest struct {
	Did       int       `json:"did"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Uid       int       `json:"uid"`
	Create_at time.Time `json:"create_at"`
	Username  string    `json:"username"`
	Avatar    string    `json:"avatar"`
}

// 获取讨论评论请求体
type CommentRequest struct {
	Cid       int    `json:"cid"`
	Did       int    `json:"did"`
	Content   string `json:"content"`
	Uid       int    `json:"uid"`
	Username  string `json:"username"`
	Avatar    string `json:"avatar"`
	Create_at string `json:"create_at"`
}

// 用户信息请求体
type UserInfoRequest struct {
	Uid      int       `json:"uid"`
	Avatar   string    `json:"avatar"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Synopsis string    `json:"synopsis"`
	Score    int       `json:"score"`
	CreateAt time.Time `json:"create_at"`
	Role     int       `json:"role"`
}

// 题目信息请求体
type ProblemInfoRequest struct {
	Pid         int    `json:"pid"`
	Difficulty  string `json:"difficulty"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Timelimit   int    `json:"time_limit"`
	Memorylimit int    `json:"memory_limit"`
	Input       string `json:"input"`
	Output      string `json:"output"`
}

// 测试样例请求体
type TestCaseRequest struct {
	InputData  string `json:"input"`
	OutputData string `json:"output"`
}

// 管理员获取题目信息请求体
type AdminProblemInfoRequest struct {
	Pid         int               `json:"pid"`
	Difficulty  string            `json:"difficulty"`
	Title       string            `json:"title"`
	Content     string            `json:"content"`
	Timelimit   string            `json:"time_limit"`
	Memorylimit string            `json:"memory_limit"`
	Input       string            `json:"input"`
	Output      string            `json:"output"`
	TestCases   []TestCaseRequest `json:"test_cases"`
}

// 用户表：uid, avatar, username, password, email, score, synopsis, submit_history, create_at
type User struct {
	Uid         int       `gorm:"comment:用户ID;primaryKey;autoIncrement"`
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
type Problem struct {
	Pid         int    `gorm:"comment:题目ID;primaryKey;autoIncrement"`
	Difficulty  string `gorm:"comment:难度;not null"`
	Title       string `gorm:"comment:题目标题;not null"`
	Content     string `gorm:"comment:题目详细;not null"`
	Timelimit   string `gorm:"comment:运行时间限制;not null"`
	Memorylimit string `gorm:"comment:内存大小限制;not null"`
	Input       string `gorm:"comment:输入样例;not null"`
	Output      string `gorm:"comment:输出样例;not null"`
}

// 总提交记录表: Pid,Uid,Result,Time,Language
type SubmitRecord struct {
	Sid      int       `gorm:"comment:提交ID;primaryKey;autoIncrement"`
	Pid      int       `gorm:"comment:题目ID"`
	Uid      int       `gorm:"comment:用户ID;not null"`
	Result   string    `gorm:"comment:结果;"`
	Time     time.Time `gorm:"comment:时间;not null"`
	Language string    `gorm:"comment:语言;not null"`
}

// 讨论帖子表: Tid,Title,Content,Uid,Create_at
type Discussion struct {
	Did       int       `gorm:"comment:讨论ID;primaryKey;autoIncrement"`
	Title     string    `gorm:"comment:标题;not null"`
	Content   string    `gorm:"comment:内容;not null"`
	Uid       int       `gorm:"comment:用户;not null"`
	Create_at time.Time `gorm:"comment:创建时间;not null"`
}

// 讨论评论表: Cid,Tid,Content,Uid,Create_at
type Comment struct {
	Cid       int       `gorm:"comment:评论ID;primaryKey;autoIncrement"`
	Did       int       `gorm:"comment:帖子ID;not null"`
	Content   string    `gorm:"comment:内容;not null"`
	Uid       int       `gorm:"comment:用户;not null"`
	Create_at time.Time `gorm:"comment:创建时间;not null"`
}

// 回复评论表: Rid, Cid, Content, Uid, Create_at
type Reply struct {
	Rid       int       `gorm:"comment:回复ID;primaryKey;autoIncrement"`
	Cid       int       `gorm:"comment:评论ID;not null"`
	Content   string    `gorm:"comment:内容;not null"`
	Uid       int       `gorm:"comment:用户ID;not null"`
	Create_at time.Time `gorm:"comment:创建时间;not null"`
}

// 测试样例表: Tid, Pid, InputData, OutputData
type TestCase struct {
	Tid        int    `gorm:"comment:测试样例ID;primaryKey;autoIncrement"`
	Pid        int    `gorm:"comment:题目ID;not null"`
	InputData  string `gorm:"comment:输入数据;not null"`
	OutputData string `gorm:"comment:输出数据;not null"`
}
