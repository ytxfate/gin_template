package logger

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func InitLogger() {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = func(t time.Time, pae zapcore.PrimitiveArrayEncoder) {
		pae.AppendString(t.Format("2006-01-02 15:04:05.000"))
	}
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeCaller = zapcore.FullCallerEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)
	// fileEncoder := zapcore.NewConsoleEncoder(encoderConfig)
	core := zapcore.NewTee(
		// 标准输出日志
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel),
		// 文件记录日志  gopkg.in/natefinch/lumberjack.v2
		// zapcore.NewCore(fileEncoder, zapcore.AddSync(&lumberjack.Logger{
		// 	Filename:   "logs/api.log",
		// 	MaxSize:    1,
		// 	MaxBackups: 2,
		// 	MaxAge:     10,
		// 	Compress:   true,
		// }), zapcore.DebugLevel),
	)
	Logger = zap.New(core)

}
