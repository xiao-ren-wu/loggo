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
