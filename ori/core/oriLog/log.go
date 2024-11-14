package oriLog

import (
	"fmt"
	"ori/core/oriConfig"
	"ori/core/oriTools/easy"
	"strings"
	"sync"
)

var (
	oriLogOnce sync.Once
	oriLogger  *LocalLogger
)

func NewLog() {
	oriLogOnce.Do(func() {
		allConfig := oriConfig.GetHotConf()
		path := easy.GetRootPath()
		dir := path + allConfig.LogPath
		_, err := easy.MakeDir(dir)
		if err != nil {
			panic(fmt.Sprintf("创建日志目录失败 path:%s", dir))
		}
		filename := strings.Replace(dir+"/"+allConfig.LogFileName, "\\", "/", -1)
		logConf := &logConfig{
			TimeFormat: "2006-01-02 15:04:05",
			Console: &consoleLogger{
				Level:    allConfig.LogLevel, //控制台日志等级
				Colorful: false,              //是否输出颜色
				Switch:   true,
			},
			File: &fileLogger{
				Filename:   filename + ".log",
				Append:     true,   //是否日志追加
				MaxLines:   100000, //日志最大行数
				MaxSize:    10,
				Daily:      true, //是否按天分割
				MaxDays:    -1,   //日志文件最大有效期
				Level:      allConfig.LogLevel,
				PermitMask: "0660", //权限
			},
			Conn:   nil,
			Format: "json", //日志格式
		}
		err = SetLogger(logConf) //设置配置
		if err != nil {
			panic(err)
		}
		oriLogger = New() //实例化
	})
}
