package coreLog

import (
	"apiServer/core/coreConfig"
	easy2 "apiServer/util/easy"
	"fmt"
	"strings"
	"sync"
)

var (
	oriLogOnce sync.Once
	OriLogger  *LocalLogger
)

func NewLog() *LocalLogger {
	oriLogOnce.Do(func() {
		allConfig := coreConfig.GetHotConf()
		path := easy2.GetRootPath()
		dir := path + allConfig.LogPath
		_, err := easy2.MakeDir(dir)
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
				MaxSize:    100,
				Daily:      true, //是否按天分割
				MaxDays:    -1,   //日志文件最大有效期
				Level:      allConfig.LogLevel,
				PermitMask: "0660", //权限
			},
			Conn:   nil,
			Format: "data", //日志格式
		}
		err = SetLogger(logConf) //设置配置
		if err != nil {
			panic(err)
		}
		OriLogger = New() //实例化
	})
	return OriLogger
}

// 记录trace级别日志
func LogTrace(format string, v ...interface{}) {
	log := NewLog()
	log.Trace(format, v...)
}

// 记录debug级别日志
func LogDebug(format string, v ...interface{}) {
	log := NewLog()
	log.Debug(format, v...)
}

// 记录info级别日志
func LogInfo(format string, v ...interface{}) {
	log := NewLog()
	log.Info(format, v...)
}

// 记录info级别日志
func LogWarn(format string, v ...interface{}) {
	log := NewLog()
	log.Warn(format, v...)
}

// 记录error级别日志
func LogError(format string, v ...interface{}) {
	log := NewLog()
	log.Error(format, v...)
}

// 记录crit级别日志
func LogCrit(format string, v ...interface{}) {
	log := NewLog()
	log.Crit(format, v...)
}

// 记录alert级别日志
func LogAlert(format string, v ...interface{}) {
	log := NewLog()
	log.Alert(format, v...)
}

// 记录emer级别日志
func LogEmer(format string, v ...interface{}) {
	log := NewLog()
	log.Emer(format, v...)
}

// 记录emer级别日志并抛出panic
func LogPanic(format string, v ...interface{}) {
	log := NewLog()
	log.Panic(format, v...)
}

// 记录emer级别日志并抛出panic
func LogFatal(format string, v ...interface{}) {
	log := NewLog()
	log.Fatal(format, v...)
}
