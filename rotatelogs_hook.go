package loggo

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"time"
)

type RotateLogsConfig struct {
	LogFilePrefix string
	RotationSize  int64
	RotationCount uint
	RotationTime  time.Duration
	MaxAge        time.Duration
}

func (c *RotateLogsConfig) genOps() (ops []rotatelogs.Option) {
	if c.MaxAge > 0 {
		ops = append(ops, rotatelogs.WithMaxAge(c.MaxAge))
	}
	if c.RotationSize > 0 {
		ops = append(ops, rotatelogs.WithRotationSize(c.RotationSize))
	}
	if c.RotationCount > 0 {
		ops = append(ops, rotatelogs.WithRotationCount(c.RotationCount))
	}
	if c.RotationTime > 0 {
		ops = append(ops, rotatelogs.WithRotationTime(c.RotationTime))
	}
	return
}

func NewRotateLogsHook(ch *RotateLogsConfig, formatter logrus.Formatter) (logrus.Hook, error) {
	logs, err := rotatelogs.New(fmt.Sprintf("%s-%s", ch.LogFilePrefix, time.Now().Format("20060102150405")), ch.genOps()...)
	if err != nil {
		return nil, err
	}
	lfsHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: logs,
		logrus.InfoLevel:  logs,
		logrus.WarnLevel:  logs,
		logrus.ErrorLevel: logs,
		logrus.FatalLevel: logs,
		logrus.PanicLevel: logs,
	}, formatter)
	return lfsHook, nil
}
