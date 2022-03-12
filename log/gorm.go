package log

import (
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"time"
)

type GormLogger struct {
	logger.Config
	LogLevel  logger.LogLevel
	ZapLogger *zap.SugaredLogger
}

func NewGormLogger(logConfig logger.Config) (logger.Interface, error) {
	zapLogger, err := zap.NewProduction()

	if err != nil {
		return nil, err
	}

	return &GormLogger{
		Config:    logConfig,
		LogLevel:  logger.Info,
		ZapLogger: zapLogger.Sugar(),
	}, nil
}

func (g *GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	newLogger := *g
	newLogger.LogLevel = level
	return &newLogger
}

func (g GormLogger) Info(ctx context.Context, s string, i ...interface{}) {
	if g.LogLevel >= logger.Info {
		g.ZapLogger.Infow(s, i...)
	}
}

func (g GormLogger) Warn(ctx context.Context, s string, i ...interface{}) {
	if g.LogLevel >= logger.Warn {
		g.ZapLogger.Warnw(s, i...)
	}
}

func (g GormLogger) Error(ctx context.Context, s string, i ...interface{}) {
	if g.LogLevel >= logger.Error {
		g.ZapLogger.Errorw(s, i...)
	}
}

func (g GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if g.LogLevel <= logger.Silent {
		return
	}

	elapsed := time.Since(begin)
	switch {
	case err != nil && g.LogLevel >= logger.Error && (!errors.Is(err, gorm.ErrRecordNotFound) || !g.IgnoreRecordNotFoundError):
		sql, rows := fc()
		if rows == -1 {
			g.ZapLogger.Errorw(err.Error(),
				zap.String("sql", sql),
				zap.Float64("time", float64(elapsed.Nanoseconds())/1e6),
				zap.String("file", utils.FileWithLineNum()),
			)
		} else {
			g.ZapLogger.Errorw(err.Error(),
				zap.String("sql", sql),
				zap.Float64("time", float64(elapsed.Nanoseconds())/1e6),
				zap.String("file", utils.FileWithLineNum()),
				zap.Int64("rows", rows),
			)
		}
	case elapsed > g.SlowThreshold && g.SlowThreshold != 0 && g.LogLevel >= logger.Warn:
		sql, rows := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", g.SlowThreshold)
		if rows == -1 {
			g.ZapLogger.Warnw(slowLog,
				zap.String("sql", sql),
				zap.Float64("time", float64(elapsed.Nanoseconds())/1e6),
				zap.String("file", utils.FileWithLineNum()),
			)
		} else {
			g.ZapLogger.Warnw(slowLog,
				zap.String("sql", sql),
				zap.Float64("time", float64(elapsed.Nanoseconds())/1e6),
				zap.String("file", utils.FileWithLineNum()),
				zap.Int64("rows", rows),
			)
		}
	case g.LogLevel == logger.Info:
		sql, rows := fc()
		if rows == -1 {
			g.ZapLogger.Infow("TRACE",
				zap.String("sql", sql),
				zap.Float64("time", float64(elapsed.Nanoseconds())/1e6),
				zap.String("file", utils.FileWithLineNum()),
			)
		} else {
			g.ZapLogger.Warnw("TRACE",
				zap.String("sql", sql),
				zap.Float64("time", float64(elapsed.Nanoseconds())/1e6),
				zap.String("file", utils.FileWithLineNum()),
				zap.Int64("rows", rows),
			)
		}
	}
}
