package loggo

import (
	"context"
	"github.com/sirupsen/logrus"
)

type config struct {
	reportCaller     bool
	timeFormat       string
	consoleColor     bool
	parseCtxFunc     func(ctx context.Context) map[string]interface{}
	setLevel         bool
	level            logrus.Level
	rotateLogsConfig *RotateLogsConfig
}

type OpsFunc func(cfg *config)

func WithCtxValue(f func(ctx context.Context) map[string]interface{}) OpsFunc {
	return func(cfg *config) {
		cfg.parseCtxFunc = f
	}
}

func WithDisableColor() OpsFunc {
	return func(cfg *config) {
		cfg.consoleColor = false
	}
}

func WithReportCaller() OpsFunc {
	return func(cfg *config) {
		cfg.reportCaller = true
	}
}

type Level uint32

// These are the different logging levels. You can set the logging level to log
const (
	// PanicLevel level, highest level of severity. Logs and then calls panic with the
	// message passed to Debug, Info, ...
	PanicLevel Level = iota
	// FatalLevel level. Logs and then calls `logger.Exit(1)`. It will exit even if the
	// logging level is set to Panic.
	FatalLevel
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	ErrorLevel
	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel
	// InfoLevel level. General operational entries about what's going on inside the
	// application.
	InfoLevel
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DebugLevel
	// TraceLevel level. Designates finer-grained informational events than the Debug.
	TraceLevel
)

func WithLevel(level Level) OpsFunc {
	return func(cfg *config) {
		cfg.setLevel = true
		cfg.level = logrus.Level(level)
	}
}

func WithRotateLogs(conf *RotateLogsConfig) OpsFunc {
	return func(cfg *config) {
		cfg.rotateLogsConfig = conf
	}
}
