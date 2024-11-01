package coreDb

import (
	"apiServer/core/coreLog"
	"context"
	"gorm.io/gorm/logger"
	"time"
)

type LogrusLogger struct {
	log *coreLog.LocalLogger
}

func (l *LogrusLogger) LogMode(level logger.LogLevel) logger.Interface {
	return l
}

func (l *LogrusLogger) Info(ctx context.Context, s string, args ...interface{}) {
	coreLog.LogInfo(s, args)
}

func (l *LogrusLogger) Warn(ctx context.Context, s string, args ...interface{}) {
	coreLog.LogWarn(s, args)
}

func (l *LogrusLogger) Error(ctx context.Context, s string, args ...interface{}) {
	coreLog.LogError(s, args)
}

func (l *LogrusLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()
	if err != nil {
		coreLog.LogError("[%s] %s", err, sql)
	} else {
		coreLog.Info("[%.2fms] [rows:%d] %s", float64(elapsed.Nanoseconds())/1e6, rows, sql)
	}
}
