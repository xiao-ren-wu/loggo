# loggo

## quick start
1. install
~~~bash
go get github.com/xiao-ren-wu/loggo
~~~
2. create loggo
~~~go
	logger, err := loggo.NewLogger()
	if err != nil {
		panic(err)
	}
	logger.Info("hello world")
~~~
3. create with ops
~~~go
import (
	"context"
	"github.com/xiao-ren-wu/loggo"
	"sync"
)

var l *loggo.Loggers
var once sync.Once

func Logs() *loggo.Loggers {
	once.Do(func() {
		logger, err := loggo.NewLogger(
			loggo.WithReportCaller(),
			loggo.WithCtxValue(func(ctx context.Context) map[string]interface{} {
				return map[string]interface{}{
					"logid": ctx.Value("logid"),
				}
			}),
			loggo.WithRotateLogs(&loggo.RotateLogsConfig{
				LogFilePrefix: "logfile/loggo",
				RotationSize:  2000,
			}),
		)
		if err != nil {
			panic(err)
		}
		l = logger
	})
	return l
}

~~~
