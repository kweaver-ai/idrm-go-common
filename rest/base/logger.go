package base

import "github.com/kweaver-ai/idrm-go-frame/core/logx/zapx"

type Logger interface {
	Debug(msg string, fields ...zapx.Field)
	Info(msg string, fields ...zapx.Field)
	Warn(msg string, fields ...zapx.Field)
	Error(msg string, fields ...zapx.Field)
	Fatal(msg string, fields ...zapx.Field)
	Trace(msg string, fields ...zapx.Field)

	Debugf(format string, v ...interface{})
	Infof(format string, v ...interface{})
	Warnf(format string, v ...interface{})
	Errorf(format string, v ...interface{})
	Fatalf(format string, v ...interface{})

	Debugw(msg string, keysAndValues ...interface{})
	Infow(msg string, keysAndValues ...interface{})
	Warnw(msg string, keysAndValues ...interface{})
	Errorw(msg string, keysAndValues ...interface{})
	Fatalw(msg string, keysAndValues ...interface{})
}
