package logs

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"runtime"
	"time"
)

type Loggers struct {
	*config
	log *logrus.Logger
}

func NewLogger(ops ...OpsFunc) (*Loggers, error) {
	conf := config{
		reportCaller: false,
		timeFormat:   time.RFC3339,
		consoleColor: true,
		rotateLogsConfig: &RotateLogsConfig{
			LogFilePrefix: "log",
		},
	}
	for _, opsFunc := range ops {
		opsFunc(&conf)
	}
	logger := logrus.New()
	if conf.parseCtxFunc != nil {
		logger.AddHook(&ctxHook{
			parseCtxFunc: conf.parseCtxFunc,
		})
	}

	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:     conf.consoleColor,
		DisableColors:   false,
		FullTimestamp:   true,
		TimestampFormat: conf.timeFormat,
	})

	rotateLogsHook, err := NewRotateLogsHook(conf.rotateLogsConfig, &logrus.TextFormatter{
		ForceColors:     conf.consoleColor,
		DisableColors:   false,
		FullTimestamp:   true,
		TimestampFormat: conf.timeFormat,
	})
	if err != nil {
		return nil, err
	}

	logger.AddHook(rotateLogsHook)

	if conf.setLevel {
		logger.SetLevel(conf.level)
	}

	return &Loggers{
		config: &conf,
		log:    logger,
	}, nil
}

func (l *Loggers) CtxFatal(ctx context.Context, format string, v ...interface{}) {
	var log logrus.Ext1FieldLogger = l.log.WithContext(ctx)
	if l.reportCaller {
		_, file, line, _ := runtime.Caller(1)
		log = log.WithField("position", fmt.Sprintf("%s.%d", file, line))
	}
	log.Fatalf(format, v...)
}

func (l *Loggers) CtxError(ctx context.Context, format string, v ...interface{}) {
	var log logrus.Ext1FieldLogger = l.log.WithContext(ctx)
	if l.reportCaller {
		_, file, line, _ := runtime.Caller(1)
		log = log.WithField("position", fmt.Sprintf("%s.%d", file, line))
	}
	log.Errorf(format, v...)
}

func (l *Loggers) CtxWarn(ctx context.Context, format string, v ...interface{}) {
	var log logrus.Ext1FieldLogger = l.log.WithContext(ctx)
	if l.reportCaller {
		_, file, line, _ := runtime.Caller(1)
		log = log.WithField("position", fmt.Sprintf("%s.%d", file, line))
	}
	log.Warnf(format, v...)
}

func (l *Loggers) CtxInfo(ctx context.Context, format string, v ...interface{}) {
	var log logrus.Ext1FieldLogger = l.log.WithContext(ctx)
	if l.reportCaller {
		_, file, line, _ := runtime.Caller(1)
		log = log.WithField("position", fmt.Sprintf("%s.%d", file, line))
	}
	log.Infof(format, v...)
}

func (l *Loggers) CtxDebug(ctx context.Context, format string, v ...interface{}) {
	var log logrus.Ext1FieldLogger = l.log.WithContext(ctx)
	if l.reportCaller {
		_, file, line, _ := runtime.Caller(1)
		log = log.WithField("position", fmt.Sprintf("%s.%d", file, line))
	}
	log.Debugf(format, v...)
}

func (l *Loggers) CtxTrace(ctx context.Context, format string, v ...interface{}) {
	var log logrus.Ext1FieldLogger = l.log.WithContext(ctx)
	if l.reportCaller {
		_, file, line, _ := runtime.Caller(1)
		log = log.WithField("position", fmt.Sprintf("%s.%d", file, line))
	}
	log.Tracef(format, v...)
}

func (l *Loggers) Fatal(format string, v ...interface{}) {
	var log logrus.Ext1FieldLogger = l.log
	if l.reportCaller {
		_, file, line, _ := runtime.Caller(1)
		log = log.WithField("position", fmt.Sprintf("%s.%d", file, line))
	}
	log.Fatalf(format, v...)
}

func (l *Loggers) Error(format string, v ...interface{}) {
	var log logrus.Ext1FieldLogger = l.log
	if l.reportCaller {
		_, file, line, _ := runtime.Caller(1)
		log = log.WithField("position", fmt.Sprintf("%s.%d", file, line))
	}
	log.Errorf(format, v...)
}

func (l *Loggers) Warn(format string, v ...interface{}) {
	var log logrus.Ext1FieldLogger = l.log
	if l.reportCaller {
		_, file, line, _ := runtime.Caller(1)
		log = log.WithField("position", fmt.Sprintf("%s.%d", file, line))
	}
	log.Warnf(format, v...)
}

func (l *Loggers) Info(format string, v ...interface{}) {
	var log logrus.Ext1FieldLogger = l.log
	if l.reportCaller {
		_, file, line, _ := runtime.Caller(1)
		log = log.WithField("position", fmt.Sprintf("%s.%d", file, line))
	}
	log.Infof(format, v...)
}

func (l *Loggers) Debug(format string, v ...interface{}) {
	var log logrus.Ext1FieldLogger = l.log
	if l.reportCaller {
		_, file, line, _ := runtime.Caller(1)
		log = log.WithField("position", fmt.Sprintf("%s.%d", file, line))
	}
	log.Debugf(format, v...)
}

func (l *Loggers) Trace(format string, v ...interface{}) {
	var log logrus.Ext1FieldLogger = l.log
	if l.reportCaller {
		_, file, line, _ := runtime.Caller(1)
		log = log.WithField("position", fmt.Sprintf("%s.%d", file, line))
	}
	log.Tracef(format, v...)
}
