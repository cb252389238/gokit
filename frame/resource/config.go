package resource

import (
	"frame/types"
	"frame/util/common"
	"os"
	"time"

	"github.com/spf13/viper"
)

var (
	hotConf    types.HotConf
	configPath string
)

func init() {
	configPath = common.GetRootPath()
}

func LoadConfig() {
	defer updateFileModifyTime()
	var conf types.Config
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

func ListenConfig() {
	configFile := configPath + "config.yaml"
	//每隔三秒读取一次配置文件查看是否发生变化
	ticker := time.NewTicker(time.Second * 10)
	for _ = range ticker.C {
		f, err := os.Open(configFile)
		if err != nil {
			LogError("listenConfig is open config file err:%+v|path:%s", err, configFile)
		} else {
			fileInfo, err := f.Stat()
			if err != nil {
				LogError("listenConfig stat file error:%+v", err)
			} else {
				curModifyTime := fileInfo.ModTime().Unix()
				if curModifyTime > hotConf.LastModifyTime {
					LoadConfig() //重新载入配置文件
					LogInfo("配置文件发生变化：%+v", hotConf.Conf)
				}
			}
			if err = f.Close(); err != nil {
				LogError("config file close fail error:%+v", err)
			}
		}
	}
}

func updateFileModifyTime() {
	hotConf.L.Lock()
	hotConf.LastModifyTime = time.Now().Unix()
	hotConf.L.Unlock()
}

//GetAllHotConf
//  @Description: 获取所有热配置
//  @return types.Config
//
func GetAllHotConf() types.Config {
	hotConf.L.RLock()
	defer hotConf.L.RUnlock()
	conf := hotConf.Conf
	return conf
}
