package plog

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/devproje/plog/level"
)

type Plog struct {
	Logger           *log.Logger
	Level            level.Level
	DisableTimestamp bool
	TimestampFormat  string
}

func New() *Plog {
	return &Plog{
		log.New(os.Stdout, "", 0),
		level.Info,
		false,
		time.RFC3339,
	}
}

func fatalHandle() {
	os.Exit(1)
}

func (pl *Plog) logging(lev level.Level, message string) {
	if pl.Level < lev {
		if pl.Level == level.Fatal {
			fatalHandle()
		}

		return
	}

	switch lev {
	case level.Trace:
		pl.Logger.Printf("%s %s", pl.prefix(traceTag), message)
	case level.Debug:
		pl.Logger.Printf("%s %s", pl.prefix(debugTag), message)
	case level.Info:
		pl.Logger.Printf("%s %s", pl.prefix(infoTag), message)
	case level.Warn:
		pl.Logger.Printf("%s %s", pl.prefix(warnTag), message)
	case level.Error:
		pl.Logger.Printf("%s %s", pl.prefix(errorTag), message)
	case level.Fatal:
		pl.Logger.Printf("%s %s", pl.prefix(fatalTag), message)
		fatalHandle()
	case level.Panic:
		pl.Logger.Panicf("%s %s", pl.prefix(panicTag), message)
	}
}

func (pl *Plog) SetOutput(writer io.Writer) {
	pl.Logger.SetOutput(writer)
}

func (pl *Plog) SetLevel(lev level.Level) {
	pl.Level = lev
}
