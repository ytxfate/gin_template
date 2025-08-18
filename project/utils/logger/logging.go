package logger

import "go.uber.org/zap"

func Debug(msg string, fields ...zap.Field) {
	GetLogger().Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	GetLogger().Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	GetLogger().Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	GetLogger().Error(msg, fields...)
}

func DPanic(msg string, fields ...zap.Field) {
	GetLogger().DPanic(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	GetLogger().Panic(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	GetLogger().Fatal(msg, fields...)
}

func Debugf(template string, args ...interface{}) {
	GetLogger().Sugar().Debugf(template, args...)
}

func Infof(template string, args ...interface{}) {
	GetLogger().Sugar().Infof(template, args...)
}

func Warnf(template string, args ...interface{}) {
	GetLogger().Sugar().Warnf(template, args...)
}

func Errorf(template string, args ...interface{}) {
	GetLogger().Sugar().Errorf(template, args...)
}

func DPanicf(template string, args ...interface{}) {
	GetLogger().Sugar().DPanicf(template, args...)
}

func Panicf(template string, args ...interface{}) {
	GetLogger().Sugar().Panicf(template, args...)
}

func Fatalf(template string, args ...interface{}) {
	GetLogger().Sugar().Fatalf(template, args...)
}
