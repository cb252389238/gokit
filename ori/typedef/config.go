package typedef

type Config struct {
	APP              string        //项目名称
	ENV              string        //环境值
	Debug            bool          //debug模式
	LogFileName      string        //日志文件名
	LogPath          string        //日志目录
	LogLevel         string        //日志输出等级
	WebHookToken     string        //钉钉token
	WebHookSecret    string        //钉钉secret
	Mysql            []Mysql       //数据库
	Redis            []Redis       //oriRedis
	Websocket        Websocket     //websocket 服务配置
	Http             Http          //http
	Monitor          MonitorConfig //监控
	StatusReportHour int           //程序状态报告时间
}

type MonitorConfig struct {
	MAX_CPU_PERCENT       float64
	CPU_FLUCTUATE         float64
	MAX_MEM_PERCENT       float64
	MEM_FLUCTUATE         float64
	MAX_DISK_PERCENT      float64
	DISK_FLUCTUATE        float64
	MAX_GOROUTINE_NUM     int
	GOROUTINE_FLUCTUATE   int
	MAX_CONCURRENCY_NUM   int
	CONCURRENCY_FLUCTUATE int
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
