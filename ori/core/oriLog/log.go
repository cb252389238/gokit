package oriLog

import (
	"fmt"
	"ori/core/oriConfig"
	oriTools2 "ori/core/oriTools"
	"strings"
	"sync"
)

var (
	once      sync.Once
	OriLogger *LocalLogger
)

func NewLog() *LocalLogger {
	once.Do(func() {
		allConfig := oriConfig.GetHotConf()
		path := oriTools2.GetRootPath()
		dir := path + allConfig.LogPath
		_, err := oriTools2.MakeDir(dir)
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
		OriLogger = New() //实例化
	})
	return OriLogger
}

// LogTrace
//
//	@Description: 记录trace级别日志
//	@param format
//	@param v
func LogTrace(format string, v ...interface{}) {
	log := NewLog()
	log.Trace(format, v...)
}

// LogDebug
//
//	@Description: 记录debug级别日志
//	@param format
//	@param v
func LogDebug(format string, v ...interface{}) {
	log := NewLog()
	log.Debug(format, v...)
}

// LogInfo
//
//	@Description: 记录info级别日志
//	@param format
//	@param v
func LogInfo(format string, v ...interface{}) {
	log := NewLog()
	log.Info(format, v...)
}

// LogWarn
//
//	@Description: 记录warn级别日志
//	@param format
//	@param v
func LogWarn(format string, v ...interface{}) {
	log := NewLog()
	log.Warn(format, v...)
}

// LogError
//
//	@Description: 记录error级别日志
//	@param format
//	@param v
func LogError(format string, v ...interface{}) {
	log := NewLog()
	log.Error(format, v...)
}

// LogCrit
//
//	@Description: 记录crit级别日志
//	@param format
//	@param v
func LogCrit(format string, v ...interface{}) {
	log := NewLog()
	log.Crit(format, v...)
}

// LogAlert
//
//	@Description: 记录alert级别日志
//	@param format
//	@param v
func LogAlert(format string, v ...interface{}) {
	log := NewLog()
	log.Alert(format, v...)
}

// LogEmer
//
//	@Description: 记录emer级别日志
//	@param format
//	@param v
func LogEmer(format string, v ...interface{}) {
	log := NewLog()
	log.Emer(format, v...)
}

// LogPanic
//
//	@Description: 记录emer级别日志并抛出panic
//	@param format
//	@param v
func LogPanic(format string, v ...interface{}) {
	log := NewLog()
	log.Panic(format, v...)
}

// LogFatal
//
//	@Description: 记录emer级别日志并退出
//	@param format
//	@param v
func LogFatal(format string, v ...interface{}) {
	log := NewLog()
	log.Fatal(format, v...)
}
