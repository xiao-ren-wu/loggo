package loggo

import (
	"context"
	"github.com/xiao-ren-wu/loggo/api"
	"github.com/xiao-ren-wu/loggo/logger"
)

var loggo api.Loggo

func init() {
	loggo, _ = logger.NewLoggo()
}

func SetLoggo(log api.Loggo) {
	loggo = log
}

func CtxFatal(ctx context.Context, format string, v ...interface{}) {
	loggo.CtxFatal(ctx, format, v...)
}
func CtxError(ctx context.Context, format string, v ...interface{}) {
	loggo.CtxError(ctx, format, v...)
}
func CtxWarn(ctx context.Context, format string, v ...interface{}) {
	loggo.CtxWarn(ctx, format, v...)
}
func CtxInfo(ctx context.Context, format string, v ...interface{}) {
	loggo.CtxInfo(ctx, format, v...)
}
func CtxDebug(ctx context.Context, format string, v ...interface{}) {
	loggo.CtxDebug(ctx, format, v...)
}
func CtxTrace(ctx context.Context, format string, v ...interface{}) {
	loggo.CtxTrace(ctx, format, v...)
}

func Fatal(format string, v ...interface{}) {
	loggo.Fatal(format, v...)
}
func Error(format string, v ...interface{}) {
	loggo.Error(format, v...)
}
func Warn(format string, v ...interface{}) {
	loggo.Warn(format, v...)
}
func Info(format string, v ...interface{}) {
	loggo.Info(format, v...)
}
func Debug(format string, v ...interface{}) {
	loggo.Debug(format, v...)
}
func Trace(format string, v ...interface{}) {
	loggo.Trace(format, v...)
}

func CtxflFatal(ctx context.Context, location string, format string, v ...interface{}) {
	loggo.CtxflFatal(ctx, location, format, v...)
}
func CtxErrorl(ctx context.Context, location string, format string, v ...interface{}) {
	loggo.CtxErrorl(ctx, location, format, v...)
}
func CtxWarnl(ctx context.Context, location string, format string, v ...interface{}) {
	loggo.CtxWarnl(ctx, location, format, v...)
}
func CtxInfol(ctx context.Context, location string, format string, v ...interface{}) {
	loggo.CtxInfol(ctx, location, format, v...)
}
func CtxDebugl(ctx context.Context, location string, format string, v ...interface{}) {
	loggo.CtxDebugl(ctx, location, format, v...)
}
func CtxTracel(ctx context.Context, location string, format string, v ...interface{}) {
	loggo.CtxTracel(ctx, location, format, v...)
}
