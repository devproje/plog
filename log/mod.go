package log

import (
	"io"

	"github.com/devproje/plog"
	"github.com/devproje/plog/level"
)

var p *plog.Plog = plog.New()

func SetOutput(writer io.Writer) {
	p.Logger.SetOutput(writer)
}

func SetLevel(lev level.Level) {
	p.Level = lev
}

func SetTimestamp(t bool) {
	p.DisableTimestamp = !t
}

func SetTimeFormat(f string) {
	p.TimestampFormat = f
}

func Log(lev level.Level, v ...any) {
	p.Log(lev, v...)
}

func Print(v ...any) {
	p.Print(v...)
}

func Trace(v ...any) {
	p.Trace(v...)
}

func Debug(v ...any) {
	p.Debug(v...)
}

func Info(v ...any) {
	p.Info(v...)
}

func Warn(v ...any) {
	p.Warn(v...)
}

func Error(v ...any) {
	p.Error(v...)
}

func Fatal(v ...any) {
	p.Fatal(v...)
}

func Panic(v ...any) {
	p.Panic(v...)
}

func Logf(lev level.Level, format string, args ...any) {
	p.Logf(lev, format, args...)
}

func Printf(format string, args ...any) {
	p.Printf(format, args...)
}

func Tracef(format string, args ...any) {
	p.Tracef(format, args...)
}

func Debugf(format string, args ...any) {
	p.Debugf(format, args...)
}

func Infof(format string, args ...any) {
	p.Infof(format, args...)
}

func Warnf(format string, args ...any) {
	p.Warnf(format, args...)
}

func Errorf(format string, args ...any) {
	p.Errorf(format, args...)
}

func Fatalf(format string, args ...any) {
	p.Fatalf(format, args...)
}

func Panicf(format string, args ...any) {
	p.Panicf(format, args...)
}

func Logln(lev level.Level, v ...any) {
	p.Logln(lev, v...)
}

func Println(v ...any) {
	p.Println(v...)
}

func Traceln(v ...any) {
	p.Traceln(v...)
}

func Debugln(v ...any) {
	p.Debugln(v...)
}

func Infoln(v ...any) {
	p.Infoln(v...)
}

func Warnln(v ...any) {
	p.Warnln(v...)
}

func Errorln(v ...any) {
	p.Errorln(v...)
}

func Fatalln(v ...any) {
	p.Fatalln(v...)
}

func Panicln(v ...any) {
	p.Panicln(v...)
}
