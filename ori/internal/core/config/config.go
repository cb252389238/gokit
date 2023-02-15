package config

import (
	"log"
	"ori/internal/core/oriCommon"
	"os"
	"time"

	"github.com/spf13/viper"
)

var (
	hotConf    HotConf
	configPath string
)

func init() {
	configPath = oriCommon.GetRootPath()
}

func Load() {
	defer updateFileModifyTime()
	var conf Config
	configFile := "config"
	configFileExt := "yaml"
	c := viper.New()
	c.SetConfigName(configFile)
	c.SetConfigType(configFileExt)
	c.AddConfigPath(configPath)
	err := c.ReadInConfig()
	if err != nil {
		panic("配置文件载入错误" + configPath + configFile + "." + configFileExt)
	}
	err = c.Unmarshal(&conf)
	if err != nil {
		panic("配置文件解析错误" + err.Error())
	}
	hotConf.L.Lock()
	hotConf.Conf = conf
	hotConf.L.Unlock()
}

func Listen(sec time.Duration) {
	configFile := configPath + "config.yaml"
	//每隔三秒读取一次配置文件查看是否发生变化
	ticker := time.NewTicker(time.Second * sec)
	for _ = range ticker.C {
		f, err := os.Open(configFile)
		if err != nil {
			log.Printf("listenConfig is open config file err:%+v|path:%s", err, configFile)
		} else {
			fileInfo, err := f.Stat()
			if err != nil {
				log.Printf("listenConfig stat file error:%+v", err)
			} else {
				curModifyTime := fileInfo.ModTime().Unix()
				if curModifyTime > hotConf.LastModifyTime {
					Load() //重新载入配置文件
					log.Printf("配置文件发生变化：%+v", hotConf.Conf)
				}
			}
			if err = f.Close(); err != nil {
				log.Printf("config file close fail error:%+v", err)
			}
		}
	}
}

func updateFileModifyTime() {
	hotConf.L.Lock()
	hotConf.LastModifyTime = time.Now().Unix()
	hotConf.L.Unlock()
}

/*
*获取热配置
 */
func GetHotConf() Config {
	hotConf.L.RLock()
	defer hotConf.L.RUnlock()
	conf := hotConf.Conf
	return conf
}
