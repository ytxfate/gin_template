package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func InitLogger() {
	Logger = zap.New(zapcore.NewCore(
		zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
		// gopkg.in/natefinch/lumberjack.v2
		// zapcore.AddSync(&lumberjack.Logger{
		// 	Filename:   "./test.log",
		// 	MaxSize:    1,
		// 	MaxBackups: 5,
		// 	MaxAge:     30,
		// 	Compress:   true,
		// }),
		zapcore.AddSync(os.Stdout),
		zapcore.DebugLevel,
	))
}
