package plog

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/devproje/plog/color"
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

func (pl *Plog) write(lev level.Level, message string) {
	type tag string
	const (
		traceTag = tag(color.GRAY + "TRAC")
		debugTag = tag(color.GRAY + "DEBU")
		infoTag  = tag(color.CYAN + "INFO")
		warnTag  = tag(color.YELLOW + "WARN")
		errorTag = tag(color.RED + "ERRO")
		fatalTag = tag(color.RED + "FATA")
		panicTag = tag(color.RED + "PANI")
	)

	var prefix = func(t tag) string {
		i := fmt.Sprintf("%s%s", t, color.RESET)
		if !pl.DisableTimestamp {
			i += pl.timestamp()
		}

		return i
	}

	if pl.Level < lev {
		if pl.Level == level.Fatal {
			os.Exit(1)
		}

		return
	}

	switch lev {
	case level.Trace:
		pl.Logger.Printf("%s %s", prefix(traceTag), message)
	case level.Debug:
		pl.Logger.Printf("%s %s", prefix(debugTag), message)
	case level.Info:
		pl.Logger.Printf("%s %s", prefix(infoTag), message)
	case level.Warn:
		pl.Logger.Printf("%s %s", prefix(warnTag), message)
	case level.Error:
		pl.Logger.Printf("%s %s", prefix(errorTag), message)
	case level.Fatal:
		pl.Logger.Printf("%s %s", prefix(fatalTag), message)
		os.Exit(1)
	case level.Panic:
		pl.Logger.Panicf("%s %s", prefix(panicTag), message)
	}
}

func (pl *Plog) SetOutput(writer io.Writer) {
	pl.Logger.SetOutput(writer)
}

func (pl *Plog) SetLevel(lev level.Level) {
	pl.Level = lev
}
