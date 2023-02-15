package config

import "sync"

type HotConf struct {
	Conf           Config
	L              sync.RWMutex
	LastModifyTime int64
}

type Config struct {
	ENV           string //环境值
	Debug         bool   //debug模式
	LogFileName   string
	LogPath       string
	LogLevel      string
	WebHookToken  string
	WebHookSecret string
	Mysql         []Mysql
	Redis         []Redis
	Websocket     Websocket //websocket 服务配置
	Http          Http
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
