package loggo

import (
	"context"
	"testing"
)

func TestWithCtxVal(t *testing.T) {
	logger, err := NewLogger(
		WithCtxValue(func(ctx context.Context) map[string]interface{} {
			return map[string]interface{}{
				"logID":   ctx.Value("logID"),
				"traceID": ctx.Value("traceID"),
				"caller":  ctx.Value("caller"),
			}
		}),
		WithRotateLogs(&RotateLogsConfig{
			LogFilePrefix: "rotatelog",
			RotationSize:  10,
			RotationCount: 0,
			RotationTime:  0,
			MaxAge:        0,
		}),
		WithReportCaller())
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
