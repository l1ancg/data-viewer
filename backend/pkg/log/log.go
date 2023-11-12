package log

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.SugaredLogger

func init() {
	logMode := zapcore.InfoLevel
	syncer := zapcore.NewMultiWriteSyncer(getWriterSyncer(), zapcore.AddSync(os.Stdout))
	core := zapcore.NewCore(getEncoder(), syncer, logMode)
	Logger = zap.New(core).Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	{
		encoderConfig.LevelKey = "level"
		encoderConfig.MessageKey = "msg"
		encoderConfig.TimeKey = "time"
		encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
		encoderConfig.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
			encoder.AppendString(t.Local().Format("2006-01-02 15:04:05"))
		}
	}

	return zapcore.NewJSONEncoder(encoderConfig)
}

func getWriterSyncer() zapcore.WriteSyncer {
	stSeparator := string(filepath.Separator)
	stRootDir, _ := os.Getwd()
	stLogFilePath := stRootDir + stSeparator + "log" + stSeparator + time.Now().Format("2006-01-02") + ".log"
	fmt.Println(stLogFilePath)

	hook := lumberjack.Logger{
		Filename:   stLogFilePath,
		MaxSize:    100,
		MaxBackups: 3,
		MaxAge:     15,
		Compress:   true,
	}

	return zapcore.AddSync(&hook)
}
