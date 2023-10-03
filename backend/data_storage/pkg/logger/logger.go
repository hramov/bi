package logger

import (
	"runtime"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

type LogLevel int

const (
	Debug   LogLevel = 0
	Info    LogLevel = 1
	Warning LogLevel = 2
	Error   LogLevel = 3
)

type Logger struct {
	log   *logrus.Entry
	level LogLevel
}

func New(service string, level LogLevel) *Logger {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		panic("Could not get context info for logger!")
	}

	filename := file[strings.LastIndex(file, "/")+1:] + ":" + strconv.Itoa(line)
	funcname := runtime.FuncForPC(pc).Name()
	fn := funcname[strings.LastIndex(funcname, ".")+1:]

	log := logrus.WithField("file", filename).WithField("function", fn).WithField("service", service)
	return &Logger{log, level}
}

func (l *Logger) Debug(msg string) {
	l.log.Debugln(msg)
}

func (l *Logger) Info(msg string) {
	l.log.Infoln(msg)
}

func (l *Logger) Warning(msg string) {
	l.log.Warningln(msg)
}

func (l *Logger) Error(msg string) {
	l.log.Errorln(msg)
}
