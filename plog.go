package plog

import (
	"io"
	"log"
	"os"

	"github.com/devproje/plog/level"
)

var (
	logger      *log.Logger
	loggerLevel = level.Info
)

func init() {
	logger = log.New(os.Stdout, "", 0)
}

func fatalHandle() {
	os.Exit(1)
}

func logging(lev level.Level, message string) {
	if loggerLevel < lev {
		if loggerLevel == level.Fatal {
			fatalHandle()
		}

		return
	}

	switch lev {
	case level.Trace:
		logger.Printf("%s %s", prefix(traceTag), message)
	case level.Debug:
		logger.Printf("%s %s", prefix(debugTag), message)
	case level.Info:
		logger.Printf("%s %s", prefix(infoTag), message)
	case level.Warn:
		logger.Printf("%s %s", prefix(warnTag), message)
	case level.Error:
		logger.Printf("%s %s", prefix(errorTag), message)
	case level.Fatal:
		logger.Printf("%s %s", prefix(fatalTag), message)
		fatalHandle()
	case level.Panic:
		logger.Panicf("%s %s", prefix(panicTag), message)
	}
}

func SetOutput(writers ...io.Writer) {
	logger.SetOutput(io.MultiWriter(writers...))
}

func SetLevel(lev level.Level) {
	loggerLevel = lev
}
