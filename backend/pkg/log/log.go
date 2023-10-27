package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	logger *zap.SugaredLogger
}

type Field struct {
	zapcore.Field
}

func ProvideLog() *Logger {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sl := logger.Sugar()
	return &Logger{logger: sl}
}

func (log *Logger) Info(template string, args ...interface{}) {
	log.logger.Infof(template, args...)
}
