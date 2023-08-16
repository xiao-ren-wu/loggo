package logger

import (
	"context"
	"github.com/sirupsen/logrus"
)

type CtxHook struct {
	abstractHook
	parseCtxFunc func(ctx context.Context) map[string]interface{}
}

func (ch *CtxHook) Fire(entry *logrus.Entry) error {
	if entry.Context != nil {
		for k, v := range ch.parseCtxFunc(entry.Context) {
			entry.Data[k] = v
		}
	}
	return nil
}
