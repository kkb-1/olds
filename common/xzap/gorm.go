package xzap

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type Log struct {
	Logger *zap.Logger
	Level  logger.LogLevel
}

func GetGormLog(logger *zap.Logger) Log {
	return Log{
		Logger: logger,
	}
}

func (l Log) LogMode(level logger.LogLevel) logger.Interface {
	l.Level = level
	return l
}

func (l Log) Info(ctx context.Context, format string, value ...interface{}) {
	if l.Level < logger.Info {
		return
	}

	l.Logger.Sugar().Infof(format, value)
}

func (l Log) Warn(ctx context.Context, format string, value ...interface{}) {
	if l.Level < logger.Warn {
		return
	}

	l.Logger.Sugar().Warnf(format, value)
}

func (l Log) Error(ctx context.Context, format string, value ...interface{}) {
	if l.Level < logger.Error {
		return
	}

	l.Logger.Sugar().Errorf(format, value)
}

func (l Log) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	// 获取运行时间
	elapsed := time.Since(begin)
	// 获取 SQL 请求和返回条数
	sql, rows := fc()

	// 通用字段
	logFields := []zap.Field{
		zap.String("sql", sql),
		zap.Duration("time", elapsed),
		zap.Int64("rows", rows),
	}

	// Gorm 错误
	if err != nil {
		// 记录未找到的错误使用 warning 等级
		if errors.Is(err, gorm.ErrRecordNotFound) {
			l.Logger.Warn("Database ErrRecordNotFound", logFields...)
		} else {
			// 其他错误使用 error 等级
			logFields = append(logFields, zap.Error(err))
			l.Logger.Error("Database Error", logFields...)
		}
	}

	// 慢查询日志
	if elapsed > (200 * time.Millisecond) {
		l.Logger.Warn("Database Slow Log", logFields...)
	}

	// 记录所有 SQL 请求
	l.Logger.Debug("Database Query", logFields...)
}
