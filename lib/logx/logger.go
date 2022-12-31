package logx

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	DefaultTimeFormat = "2006-01-02 15:04:05"
)

type Logger struct {
	timeFormat string
	level      level
	Writer     io.Writer
}

func NewLogger(opts ...Option) *Logger {
	l := &Logger{
		Writer:     os.Stdout,
		timeFormat: DefaultTimeFormat,
		level:      infoLevel,
	}

	for _, opt := range opts {
		opt.Apply(l)
	}
	return l
}

func (l *Logger) log(level level, msg string, params ...any) {
	if level < l.level {
		return
	}

	if len(params) > 0 {
		msg = fmt.Sprintf(msg, params...)
	}
	msg = fmt.Sprintf("%s [%s] %s \n", time.Now().Format(l.timeFormat), level.String(), msg)
	l.Writer.Write([]byte(msg))
}

func (l *Logger) Debug(msg string, params ...any) {
	l.log(debugLevel, msg, params...)
}

func (l *Logger) Info(msg string, params ...any) {
	l.log(infoLevel, msg, params...)
}

func (l *Logger) Error(msg string, params ...any) {
	l.log(errorLevel, msg, params...)
}

func (l *Logger) Warn(msg string, params ...any) {
	l.log(warnLevel, msg, params...)
}
