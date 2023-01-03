package plog

import (
	"fmt"

	"github.com/devproje/plog/level"
)

func (pl *Plog) Log(lev level.Level, v ...any) {
	pl.write(lev, fmt.Sprint(v...))
}

func (pl *Plog) Print(v ...any) {
	pl.Info(v...)
}

func (pl *Plog) Trace(v ...any) {
	pl.Log(level.Trace, v...)
}

func (pl *Plog) Debug(v ...any) {
	pl.Log(level.Debug, v...)
}

func (pl *Plog) Info(v ...any) {
	pl.Log(level.Info, v...)
}

func (pl *Plog) Warn(v ...any) {
	pl.Log(level.Warn, v...)
}

func (pl *Plog) Error(v ...any) {
	pl.Log(level.Error, v...)
}

func (pl *Plog) Fatal(v ...any) {
	pl.Log(level.Fatal, v...)
}

func (pl *Plog) Panic(v ...any) {
	pl.Log(level.Panic, v...)
}

func (pl *Plog) Logf(lev level.Level, format string, args ...any) {
	pl.write(lev, fmt.Sprintf(format, args...))
}

func (pl *Plog) Printf(format string, args ...any) {
	pl.Infof(format, args...)
}

func (pl *Plog) Tracef(format string, args ...any) {
	pl.Logf(level.Trace, format, args...)
}

func (pl *Plog) Debugf(format string, args ...any) {
	pl.Logf(level.Debug, format, args...)
}

func (pl *Plog) Infof(format string, args ...any) {
	pl.Logf(level.Info, format, args...)
}

func (pl *Plog) Warnf(format string, args ...any) {
	pl.Logf(level.Warn, format, args...)
}

func (pl *Plog) Errorf(format string, args ...any) {
	pl.Logf(level.Error, format, args...)
}

func (pl *Plog) Fatalf(format string, args ...any) {
	pl.Logf(level.Fatal, format, args...)
}

func (pl *Plog) Panicf(format string, args ...any) {
	pl.Logf(level.Panic, format, args...)
}

func (pl *Plog) Logln(lev level.Level, v ...any) {
	pl.write(lev, fmt.Sprintln(v...))
}

func (pl *Plog) Println(v ...any) {
	pl.Infoln(v...)
}

func (pl *Plog) Traceln(v ...any) {
	pl.Logln(level.Trace, v...)
}

func (pl *Plog) Debugln(v ...any) {
	pl.Logln(level.Debug, v...)
}

func (pl *Plog) Infoln(v ...any) {
	pl.Logln(level.Info, v...)
}

func (pl *Plog) Warnln(v ...any) {
	pl.Logln(level.Warn, v...)
}

func (pl *Plog) Errorln(v ...any) {
	pl.Logln(level.Error, v...)
}

func (pl *Plog) Fatalln(v ...any) {
	pl.Logln(level.Fatal, v...)
}

func (pl *Plog) Panicln(v ...any) {
	pl.Logln(level.Panic, v...)
}
