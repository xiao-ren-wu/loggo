package logger

import (
	"context"
	"github.com/sirupsen/logrus"
)

type Config struct {
	reportCaller     bool
	skipPoint        int
	timeFormat       string
	consoleColor     bool
	parseCtxFunc     func(ctx context.Context) map[string]interface{}
	setLevel         bool
	level            logrus.Level
	rotateLogsConfig *RotateLogsConfig
}

type OpsFunc func(cfg *Config)

func WithCtxValue(f func(ctx context.Context) map[string]interface{}) OpsFunc {
	return func(cfg *Config) {
		cfg.parseCtxFunc = f
	}
}

func WithDisableColor() OpsFunc {
	return func(cfg *Config) {
		cfg.consoleColor = false
	}
}

func WithReportCaller(skipPoint ...int) OpsFunc {
	var skp = 1
	if len(skipPoint) > 0 {
		skp = skipPoint[0]
	}
	return func(cfg *Config) {
		cfg.reportCaller = true
		cfg.skipPoint = skp
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
	return func(cfg *Config) {
		cfg.setLevel = true
		cfg.level = logrus.Level(level)
	}
}

func WithRotateLogs(conf *RotateLogsConfig) OpsFunc {
	return func(cfg *Config) {
		cfg.rotateLogsConfig = conf
	}
}
