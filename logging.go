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

func Warn(v ...any) {
	Log(level.Warn, v...)
}

func Error(v ...any) {
	Log(level.Error, v...)
}

func Fatal(v ...any) {
	Log(level.Fatal, v...)
}

func Logf(lev level.Level, format string, args ...any) {
	logging(lev, fmt.Sprintf(format, args...))
}

func Infof(format string, args ...any) {
	Logf(level.Info, format, args...)
}

func Warnf(format string, args ...any) {
	Logf(level.Warn, format, args...)
}

func Errorf(format string, args ...any) {
	Logf(level.Error, format, args...)
}

func Fatalf(format string, args ...any) {
	Logf(level.Fatal, format, args...)
}

func Logln(lev level.Level, v ...any) {
	logging(lev, fmt.Sprintln(v...))
}

func Infoln(v ...any) {
	Logln(level.Info, v...)
}

func Warnln(v ...any) {
	Logln(level.Warn, v...)
}

func Errorln(v ...any) {
	Logln(level.Error, v...)
}

func Fatalln(v ...any) {
	Logln(level.Fatal, v...)
}
