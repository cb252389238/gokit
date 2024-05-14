package oriDb

import (
	"context"
	"gorm.io/gorm/logger"
	"ori/core/oriLog"
	"time"
)

type LogrusLogger struct {
	log *oriLog.LocalLogger
}

func (l *LogrusLogger) LogMode(level logger.LogLevel) logger.Interface {
	return l
}

func (l *LogrusLogger) Info(ctx context.Context, s string, args ...interface{}) {
	oriLog.LogInfo(s, args)
}

func (l *LogrusLogger) Warn(ctx context.Context, s string, args ...interface{}) {
	oriLog.LogWarn(s, args)
}

func (l *LogrusLogger) Error(ctx context.Context, s string, args ...interface{}) {
	oriLog.LogError(s, args)
}

func (l *LogrusLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()
	if err != nil {
		oriLog.LogError("[%s] %s", err, sql)
	} else {
		oriLog.Info("[%.2fms] [rows:%d] %s", float64(elapsed.Nanoseconds())/1e6, rows, sql)
	}
}
