package coreConfig

import (
	"apiServer/typedef"
	easy2 "apiServer/util/easy"
	"log"
	"os"
	"strings"
	"time"

	"github.com/spf13/viper"
)

var (
	hotConf                                    HotConf
	configPath                                 string
	configFullName, configSuffix, configPrefix string
	initialPath                                string
)

// Load
//
//	@Description:	载入配置文件
//	@param			path
func Load(path string) {
	defer updateFileModifyTime()
	initialPath = path
	configFullName, configSuffix, configPrefix = easy2.FileInfo(path)
	configPath = easy2.GetRootPath() + strings.Replace(path[:len(path)-len(configFullName)], "./", "", -1)
	var conf typedef.Config
	c := viper.New()
	c.SetConfigName(configPrefix)
	c.SetConfigType(configSuffix[1:])
	c.AddConfigPath(configPath)
	err := c.ReadInConfig()
	if err != nil {
		panic("配置文件载入错误：" + configPath)
	}
	err = c.Unmarshal(&conf)
	if err != nil {
		panic("配置文件解析错误" + err.Error())
	}
	hotConf.L.Lock()
	hotConf.Conf = conf
	hotConf.L.Unlock()
}

// Listen
//
//	@Description:	监听配置文件变化
//	@param			sec
func Listen(sec time.Duration) {
	configFile := configPath + configFullName
	//每隔三秒读取一次配置文件查看是否发生变化
	ticker := time.NewTicker(time.Second * sec)
	for _ = range ticker.C {
		f, err := os.Open(configFile)
		if err != nil {
			log.Printf("listenConfig is open coreConfig file err:%+v|path:%s", err, configFile)
		} else {
			fileInfo, err := f.Stat()
			if err != nil {
				log.Printf("listenConfig stat file error:%+v", err)
			} else {
				curModifyTime := fileInfo.ModTime().Unix()
				if curModifyTime > hotConf.LastModifyTime {
					Load(initialPath) //重新载入配置文件
					log.Printf("配置文件发生变化：%+v", hotConf.Conf)
				}
			}
			if err = f.Close(); err != nil {
				log.Printf("coreConfig file close fail error:%+v", err)
			}
		}
	}
}

// updateFileModifyTime
//
//	@Description:	更新文件更新时间
func updateFileModifyTime() {
	hotConf.L.Lock()
	hotConf.LastModifyTime = time.Now().Unix()
	hotConf.L.Unlock()
}

// GetHotConf
//
//	@Description:	获取热配置
//	@return			typedef.Config
func GetHotConf() typedef.Config {
	hotConf.L.RLock()
	defer hotConf.L.RUnlock()
	conf := hotConf.Conf
	return conf
}
