package log

import (
	"encoding/json"
	"fmt"
	"github.com/evalphobia/logrus_sentry"
	"github.com/getsentry/sentry-go"
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
)

// Create a new instance of the logger. You can have any number of instances.
var log = logrus.New()

func Init() {
	dsn := os.Getenv("SENTRY_DSN")

	if len(dsn) == 0 {
		return
	}

	err := sentry.Init(sentry.ClientOptions{
		Dsn:   dsn,
		Debug: true,
	})
	if err != nil {
		return
	}

	hook, err := logrus_sentry.NewSentryHook(dsn, []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
	})
	hook.StacktraceConfiguration.Skip = 2
	hook.StacktraceConfiguration.Enable = true
	if err == nil {
		log.Hooks.Add(hook)
	}
}

func Info(args ...interface{}) {
	log.Info(args...)
}

func Fatal(args ...interface{}) {
	log.Fatal(args...)
}

func WithField(key string, value interface{}) *logrus.Entry {
	return log.WithField(key, value)
}

func Printf(format string, args ...interface{}) {
	log.Printf(format, args...)
}

func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

func Infoln(args ...interface{}) {
	log.Infoln(args...)
}

func JSON(v interface{}) {
	msg, _ := json.MarshalIndent(v, "", "    ")
	log.Info(string(msg))
}

func Warn(msg string) {
	log.WithFields(logrus.Fields{
		"func": trace(1),
	}).Warn(msg)
}

func Warnf(format string, args ...interface{}) {
	log.WithFields(logrus.Fields{
		"func": trace(1),
	}).Warnf(format, args...)
}

func Error(err error) {
	log.WithFields(logrus.Fields{
		"func": trace(1),
	}).Error(err)
}

func Error2(err error) {
	log.WithFields(logrus.Fields{
		"func": trace(2),
	}).Error(err)
}

func trace(n int) string {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[n]) // get the 2nd one, the caller
	file, line := f.FileLine(pc[n])
	return fmt.Sprintf("%s:%d %s\n", file, line, f.Name())
}
