package logger

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/xiao-ren-wu/loggo/api"
	"runtime"
	"time"
)

type LoggoImpl struct {
	*Config
	log *logrus.Logger
}

func NewLoggo(ops ...OpsFunc) (api.Loggo, error) {
	conf := Config{
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
		logger.AddHook(&CtxHook{
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

	return &LoggoImpl{
		Config: &conf,
		log:    logger,
	}, nil
}

func (l *LoggoImpl) CtxFatal(ctx context.Context, format string, v ...interface{}) {
	var log logrus.Ext1FieldLogger = l.log.WithContext(ctx)
	if l.reportCaller {
		_, file, line, _ := runtime.Caller(l.skipPoint)
		log = log.WithField("position", fmt.Sprintf("%s.%d", file, line))
	}
	log.Fatalf(format, v...)
}

func (l *LoggoImpl) CtxError(ctx context.Context, format string, v ...interface{}) {
	var log logrus.Ext1FieldLogger = l.log.WithContext(ctx)
	if l.reportCaller {
		_, file, line, _ := runtime.Caller(l.skipPoint)
		log = log.WithField("position", fmt.Sprintf("%s.%d", file, line))
	}
	log.Errorf(format, v...)
}

func (l *LoggoImpl) CtxWarn(ctx context.Context, format string, v ...interface{}) {
	var log logrus.Ext1FieldLogger = l.log.WithContext(ctx)
	if l.reportCaller {
		_, file, line, _ := runtime.Caller(l.skipPoint)
		log = log.WithField("position", fmt.Sprintf("%s.%d", file, line))
	}
	log.Warnf(format, v...)
}

func (l *LoggoImpl) CtxInfo(ctx context.Context, format string, v ...interface{}) {
	var log logrus.Ext1FieldLogger = l.log.WithContext(ctx)
	if l.reportCaller {
		_, file, line, _ := runtime.Caller(l.skipPoint)
		log = log.WithField("position", fmt.Sprintf("%s.%d", file, line))
	}
	log.Infof(format, v...)
}

func (l *LoggoImpl) CtxDebug(ctx context.Context, format string, v ...interface{}) {
	var log logrus.Ext1FieldLogger = l.log.WithContext(ctx)
	if l.reportCaller {
		_, file, line, _ := runtime.Caller(l.skipPoint)
		log = log.WithField("position", fmt.Sprintf("%s.%d", file, line))
	}
	log.Debugf(format, v...)
}

func (l *LoggoImpl) CtxTrace(ctx context.Context, format string, v ...interface{}) {
	var log logrus.Ext1FieldLogger = l.log.WithContext(ctx)
	if l.reportCaller {
		_, file, line, _ := runtime.Caller(l.skipPoint)
		log = log.WithField("position", fmt.Sprintf("%s.%d", file, line))
	}
	log.Tracef(format, v...)
}

func (l *LoggoImpl) Fatal(format string, v ...interface{}) {
	var log logrus.Ext1FieldLogger = l.log
	if l.reportCaller {
		_, file, line, _ := runtime.Caller(l.skipPoint)
		log = log.WithField("position", fmt.Sprintf("%s.%d", file, line))
	}
	log.Fatalf(format, v...)
}

func (l *LoggoImpl) Error(format string, v ...interface{}) {
	var log logrus.Ext1FieldLogger = l.log
	if l.reportCaller {
		_, file, line, _ := runtime.Caller(l.skipPoint)
		log = log.WithField("position", fmt.Sprintf("%s.%d", file, line))
	}
	log.Errorf(format, v...)
}

func (l *LoggoImpl) Warn(format string, v ...interface{}) {
	var log logrus.Ext1FieldLogger = l.log
	if l.reportCaller {
		_, file, line, _ := runtime.Caller(l.skipPoint)
		log = log.WithField("position", fmt.Sprintf("%s.%d", file, line))
	}
	log.Warnf(format, v...)
}

func (l *LoggoImpl) Info(format string, v ...interface{}) {
	var log logrus.Ext1FieldLogger = l.log
	if l.reportCaller {
		_, file, line, _ := runtime.Caller(l.skipPoint)
		log = log.WithField("position", fmt.Sprintf("%s.%d", file, line))
	}
	log.Infof(format, v...)
}

func (l *LoggoImpl) Debug(format string, v ...interface{}) {
	var log logrus.Ext1FieldLogger = l.log
	if l.reportCaller {
		_, file, line, _ := runtime.Caller(l.skipPoint)
		log = log.WithField("position", fmt.Sprintf("%s.%d", file, line))
	}
	log.Debugf(format, v...)
}

func (l *LoggoImpl) Trace(format string, v ...interface{}) {
	var log logrus.Ext1FieldLogger = l.log
	if l.reportCaller {
		_, file, line, _ := runtime.Caller(l.skipPoint)
		log = log.WithField("position", fmt.Sprintf("%s.%d", file, line))
	}
	log.Tracef(format, v...)
}

func (l *LoggoImpl) CtxflFatal(ctx context.Context, location string, format string, v ...interface{}) {
	var log logrus.Ext1FieldLogger = l.log.WithContext(ctx)
	log = log.WithField("position", location)
	log.Fatalf(format, v...)
}

func (l *LoggoImpl) CtxErrorl(ctx context.Context, location string, format string, v ...interface{}) {
	var log logrus.Ext1FieldLogger = l.log.WithContext(ctx)
	log = log.WithField("position", location)
	log.Errorf(format, v...)
}

func (l *LoggoImpl) CtxWarnl(ctx context.Context, location string, format string, v ...interface{}) {
	var log logrus.Ext1FieldLogger = l.log.WithContext(ctx)
	log = log.WithField("position", location)
	log.Warnf(format, v...)
}

func (l *LoggoImpl) CtxInfol(ctx context.Context, location string, format string, v ...interface{}) {
	var log logrus.Ext1FieldLogger = l.log.WithContext(ctx)
	log = log.WithField("position", location)
	log.Infof(format, v...)
}

func (l *LoggoImpl) CtxDebugl(ctx context.Context, location string, format string, v ...interface{}) {
	var log logrus.Ext1FieldLogger = l.log.WithContext(ctx)
	log = log.WithField("position", location)
	log.Debugf(format, v...)
}

func (l *LoggoImpl) CtxTracel(ctx context.Context, location string, format string, v ...interface{}) {
	var log logrus.Ext1FieldLogger = l.log.WithContext(ctx)
	log = log.WithField("position", location)
	log.Tracef(format, v...)
}
