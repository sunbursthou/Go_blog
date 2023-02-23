package configs

var Cfg Config

// 配置文件的结构体
type Config struct {
	Server Server
	Mysql  Mysql
}

type Server struct {
	AppMode   string
	BackPort  string // 后端接口
	FrontPort string // 前端接口
}

type Mysql struct {
	Host     string
	Port     string
	Config   string // 其他配置
	Dbname   string
	Username string
	Password string
	LogMode  string // 日志级别
}
type Redis struct {
	DB       int    // 指定 Redis 数据库
	Addr     string // 服务器地址:端口
	Password string // 密码
}
