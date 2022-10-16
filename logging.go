package plog

import (
	"fmt"
	"github.com/devproje/plog/level"
)

func Log(lev level.Level, v ...any) {
	logging(lev, fmt.Sprint(v...))
}

func Info(v ...any) {
	Log(level.Info, v...)
}

func Logf(lev level.Level, format string, args ...any) {
	logging(lev, fmt.Sprintf(format, args...))
}

func Infof(format string, args ...any) {
	Logf(level.Info, format, args...)
}
