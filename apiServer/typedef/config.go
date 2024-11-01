package typedef

type Config struct {
	APP         string  //项目名称
	ENV         string  //环境值
	Debug       bool    //debug模式
	LogFileName string  //日志文件名
	LogPath     string  //日志目录
	LogLevel    string  //日志输出等级
	ImagePrefix string  //图片前缀
	Mysql       []Mysql //数据库
	Redis       []Redis //Redis
	Http        Http    //http
	JwtSecret   string  //token密钥
	Kafka       Kafka
	Pgsql       []Pgsql
	InnerIp     []string
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

type Http struct {
	Port uint16
}

type Kafka struct {
	Action      KafkaConf //从业者主题
	Gift        KafkaConf //礼物主题
	PrivateChat KafkaConf //私聊主题
	PublicChat  KafkaConf //公屏主题
}

type KafkaConf struct {
	Addr  []string
	Topic string
}

type Pgsql struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	Name     string
}
