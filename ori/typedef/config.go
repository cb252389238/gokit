package typedef

type Config struct {
	APP         string    //项目名称
	ENV         string    //环境值
	Debug       bool      //debug模式
	LogFileName string    //日志文件名
	LogPath     string    //日志目录
	LogLevel    string    //日志输出等级
	Mysql       []Mysql   //数据库
	Redis       []Redis   //oriRedis
	Websocket   Websocket //websocket 服务配置
	Http        Http      //http
}

type Mysql struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	Name     string
}

type Redis struct {
	Host     string
	Port     string
	Password string
	Database string
	Name     string
}

type Websocket struct {
	Port uint16
}
type Http struct {
	Port uint16
}
