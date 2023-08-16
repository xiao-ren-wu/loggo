package api

import "context"

type Loggo interface {
	CtxLoggerInterface
	LoggerInterface
}

type CtxLoggerInterface interface {
	CtxFatal(ctx context.Context, format string, v ...interface{})
	CtxError(ctx context.Context, format string, v ...interface{})
	CtxWarn(ctx context.Context, format string, v ...interface{})
	CtxInfo(ctx context.Context, format string, v ...interface{})
	CtxDebug(ctx context.Context, format string, v ...interface{})
	CtxTrace(ctx context.Context, format string, v ...interface{})
}

type LoggerInterface interface {
	Fatal(format string, v ...interface{})
	Error(format string, v ...interface{})
	Warn(format string, v ...interface{})
	Info(format string, v ...interface{})
	Debug(format string, v ...interface{})
	Trace(format string, v ...interface{})
}
