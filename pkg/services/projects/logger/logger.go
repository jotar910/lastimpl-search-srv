package logger

import (
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	filename = path.Join(fmt.Sprintf("%ctmp", os.PathSeparator), fmt.Sprintf("%d.log", time.Now().UnixMilli()))
	level    = logrus.TraceLevel
)

type Logger interface {
	Log(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Trace(args ...interface{})
	Debug(args ...interface{})
	WithPrefix(prefix string) Logger
	WithField(key string, value interface{}) Logger
}

type logrusLogger interface {
	Println(args ...interface{})
	Errorln(args ...interface{})
	Fatalln(args ...interface{})
	Traceln(args ...interface{})
	Debugln(args ...interface{})
	WithField(key string, value interface{}) *logrus.Entry
}

type logger struct {
	logrusLogger
	prefix []string
}

func New(prefix string, console bool) *logger {
	l := logrus.New()
	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = fmt.Sprintf("%s 2006-01-02 15:04:05.000000000", prefix)
	customFormatter.FullTimestamp = true
	l.SetFormatter(customFormatter)
	l.SetLevel(level)
	if !console {
		file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)
		if err == nil {
			l.Println("Running tool...")
			l.Printf("Logging to %q\n", filename)
			l.SetOutput(file)
		} else {
			l.Errorf("Failed to log to file, using default stdout/stderr")
		}
	}
	return &logger{l, []string{prefix}}
}

func (l *logger) Log(args ...interface{}) {
	l.print(l.Println, args)
}

func (l *logger) Error(args ...interface{}) {
	l.print(l.Errorln, args)
}

func (l *logger) Fatal(args ...interface{}) {
	l.print(l.Fatalln, args)
}

func (l *logger) Trace(args ...interface{}) {
	l.print(l.Traceln, args)
}

func (l *logger) Debug(args ...interface{}) {
	l.print(l.Debugln, args)
}

func (l *logger) WithPrefix(prefix string) Logger {
	return &logger{
		logrusLogger: l.logrusLogger,
		prefix:       append(l.prefix, prefix),
	}
}

func (l *logger) WithField(key string, value interface{}) Logger {
	return &logger{
		logrusLogger: l.logrusLogger.WithField(key, value),
		prefix:       l.prefix,
	}
}

func (l *logger) print(fn func(...interface{}), args []interface{}) {
	pLen := len(l.prefix)
	logs := make([]string, pLen+len(args))
	for i, p := range l.prefix {
		logs[i] = p
	}
	for i, arg := range args {
		logs[i+pLen] = fmt.Sprintf("%v", arg)
	}
	fn(strings.Join(logs, ": "))
}
