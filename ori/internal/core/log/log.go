package log

import (
	"fmt"
	"ori/internal/core/config"
	"ori/internal/core/oriTools"
	"strings"
	"sync"
)

var (
	once      sync.Once
	OriLogger *LocalLogger
)

func NewLog() *LocalLogger {
	once.Do(func() {
		allConfig := config.GetHotConf()
		path := oriTools.GetRootPath()
		dir := path + allConfig.LogPath
		_, err := oriTools.MakeDir(dir)
		if err != nil {
			panic(fmt.Sprintf("创建日志目录失败 path:%s", dir))
		}
		filename := strings.Replace(dir+"/"+allConfig.LogFileName+".logs", "\\", "/", -1)
		loggerJson := `{
    "TimeFormat": "2006-01-02 15:04:05",
    "Format": "json",
    "Console": {
        "level": "` + allConfig.LogLevel + `",
        "color": true
    },
    "File": {
        "filename": "` + filename + `",
        "level": "` + allConfig.LogLevel + `",
        "daily": true,
        "maxlines": 100000,
        "maxsize": 10,
        "maxdays": -1,
        "append": true,
        "permit": "0660"
    }
}`
		err = SetLogger(loggerJson)
		if err != nil {
			panic(err)
		}
		OriLogger = New()

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
