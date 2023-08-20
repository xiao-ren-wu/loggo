package loggo_test

import (
	"context"
	"github.com/xiao-ren-wu/loggo"
	"github.com/xiao-ren-wu/loggo/logger"
	"testing"
)

func TestWithCtxVal(t *testing.T) {
	logger, err := logger.NewLoggo(
		logger.WithCtxValue(func(ctx context.Context) map[string]interface{} {
			return map[string]interface{}{
				"logID":   ctx.Value("logID"),
				"traceID": ctx.Value("traceID"),
				"caller":  ctx.Value("caller"),
			}
		}),
		logger.WithRotateLogs(&logger.RotateLogsConfig{
			LogFilePrefix: "rotatelog",
			RotationSize:  10,
			RotationCount: 0,
			RotationTime:  0,
			MaxAge:        0,
		}),
		logger.WithReportCaller())
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.WithValue(context.Background(), "logID", "123545")
	ctx = context.WithValue(ctx, "traceID", "2023")
	ctx = context.WithValue(ctx, "caller", "test.service")
	//logger.CtxFatal(ctx, "helloyyyyyyyyyyyyyyyuggggggggggggggggggggggggggggggggggg")
	logger.CtxInfo(ctx, "helloyyyyyyyyyyyyyyyuggggggggggggggggggggggggggggggggggg")
	logger.Error("helloyyyyyyyyyyyyyyyuggggggggggggggggggggggggggggggggggg")
	logger.Warn("helloyyyyyyyyyyyyyyyuggggggggggggggggggggggggggggggggggg")
}

func TestLoggo(t *testing.T) {
	//log, err := logger.NewLoggo()
	//if err != nil {
	//	t.Fatal(err)
	//}
	ctx := context.WithValue(context.Background(), "logID", "123545")
	ctx = context.WithValue(ctx, "traceID", "2023")
	ctx = context.WithValue(ctx, "caller", "test.service")
	//loggo.SetLoggo(log)
	loggo.CtxInfo(ctx, "hello world")
	loggo.CtxInfol(ctx, "lo", "helloworld")
}
