package logger

import (
	"fmt"
	"os"
	"path"
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
}

type logger struct {
	*logrus.Logger
	prefix string
}

func New(prefix string, console bool) Logger {
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
	return &logger{l, prefix}
}

func (l *logger) Log(args ...interface{}) {
	l.Println(args)
}

func (l *logger) Error(args ...interface{}) {
	l.Errorln(args)
}

func (l *logger) Fatal(args ...interface{}) {
	l.Fatalln(args)
}

func (l *logger) Trace(args ...interface{}) {
	l.Traceln(args)
}

func (l *logger) Debug(args ...interface{}) {
	l.Debugln(args)
}
