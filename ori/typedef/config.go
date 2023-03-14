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
	Redis            []Redis       //redis
	Websocket        Websocket     //websocket 服务配置
	Http             Http          //http
	Monitor          MonitorConfig //监控
	Services         Services      //开启得服务列表
	StatusReportHour int           //程序状态报告时间
}

type Services struct {
	CONFIG_HOT_UPDATE_SERVER bool //配置热更新服务
	MONITOR_SERVER           bool //监控服务
	SQL_SERVER               bool //数据库服务
	REDIS_SERVER             bool //redis服务
	STATUS_REPORT_SERVER     bool //状态报告服务
	HTTP_SERVER              bool //http服务
	WEBSOCKET_SERVER         bool //websocket服务
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
