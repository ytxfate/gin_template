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

func Debugf(template string, args ...any) {
	GetLogger().Sugar().Debugf(template, args...)
}

func Infof(template string, args ...any) {
	GetLogger().Sugar().Infof(template, args...)
}

func Warnf(template string, args ...any) {
	GetLogger().Sugar().Warnf(template, args...)
}

func Errorf(template string, args ...any) {
	GetLogger().Sugar().Errorf(template, args...)
}

func DPanicf(template string, args ...any) {
	GetLogger().Sugar().DPanicf(template, args...)
}

func Panicf(template string, args ...any) {
	GetLogger().Sugar().Panicf(template, args...)
}

func Fatalf(template string, args ...any) {
	GetLogger().Sugar().Fatalf(template, args...)
}

func NDebug(name string, msg string, fields ...zap.Field) {
	GetNameLogger(name).Debug(msg, fields...)
}

func NInfo(name string, msg string, fields ...zap.Field) {
	GetNameLogger(name).Info(msg, fields...)
}

func NWarn(name string, msg string, fields ...zap.Field) {
	GetNameLogger(name).Warn(msg, fields...)
}

func NError(name string, msg string, fields ...zap.Field) {
	GetNameLogger(name).Error(msg, fields...)
}

func NDPanic(name string, msg string, fields ...zap.Field) {
	GetNameLogger(name).DPanic(msg, fields...)
}

func NPanic(name string, msg string, fields ...zap.Field) {
	GetNameLogger(name).Panic(msg, fields...)
}

func NFatal(name string, msg string, fields ...zap.Field) {
	GetNameLogger(name).Fatal(msg, fields...)
}

func NDebugf(name string, template string, args ...any) {
	GetNameLogger(name).Sugar().Debugf(template, args...)
}

func NInfof(name string, template string, args ...any) {
	GetNameLogger(name).Sugar().Infof(template, args...)
}

func NWarnf(name string, template string, args ...any) {
	GetNameLogger(name).Sugar().Warnf(template, args...)
}

func NErrorf(name string, template string, args ...any) {
	GetNameLogger(name).Sugar().Errorf(template, args...)
}

func NDPanicf(name string, template string, args ...any) {
	GetNameLogger(name).Sugar().DPanicf(template, args...)
}

func NPanicf(name string, template string, args ...any) {
	GetNameLogger(name).Sugar().Panicf(template, args...)
}

func NFatalf(name string, template string, args ...any) {
	GetNameLogger(name).Sugar().Fatalf(template, args...)
}
