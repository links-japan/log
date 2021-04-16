package log

import (
	"github.com/evalphobia/logrus_sentry"
	"github.com/getsentry/sentry-go"
	"github.com/sirupsen/logrus"
	"os"
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

func WithField(key string, value interface{}) *logrus.Entry { return log.WithField(key, value) }
func WithFields(fields logrus.Fields) *logrus.Entry         { return log.WithFields(fields) }
func WithError(err error) *logrus.Entry                     { return log.WithError(err) }
func Debugf(format string, args ...interface{})             { log.Debugf(format, args...) }
func Infof(format string, args ...interface{})              { log.Infof(format, args...) }
func Printf(format string, args ...interface{})             { log.Printf(format, args...) }
func Warnf(format string, args ...interface{})              { log.Warnf(format, args...) }
func Warningf(format string, args ...interface{})           { log.Warningf(format, args...) }
func Errorf(format string, args ...interface{})             { log.Errorf(format, args...) }
func Fatalf(format string, args ...interface{})             { log.Fatalf(format, args...) }
func Panicf(format string, args ...interface{})             { log.Panicf(format, args...) }
func Debug(args ...interface{})                             { log.Debug(args...) }
func Info(args ...interface{})                              { log.Info(args...) }
func Print(args ...interface{})                             { log.Print(args...) }
func Warn(args ...interface{})                              { log.Warn(args...) }
func Warning(args ...interface{})                           { log.Warning(args...) }
func Error(args ...interface{})                             { log.Error(args...) }
func Fatal(args ...interface{})                             { log.Fatal(args...) }
func Panic(args ...interface{})                             { log.Panic(args...) }
func Debugln(args ...interface{})                           { log.Debugln(args...) }
func Infoln(args ...interface{})                            { log.Infoln(args...) }
func Println(args ...interface{})                           { log.Println(args...) }
func Warnln(args ...interface{})                            { log.Warnln(args...) }
func Warningln(args ...interface{})                         { log.Warningln(args...) }
func Errorln(args ...interface{})                           { log.Errorln(args...) }
func Fatalln(args ...interface{})                           { log.Fatalln(args...) }
func Panicln(args ...interface{})                           { log.Panicln(args...) }
