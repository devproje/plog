package plog

import (
	"fmt"
	"github.com/devproje/plog/color"
	"github.com/devproje/plog/level"
	"io"
	"log"
	"os"
	"time"
)

var (
	logger *log.Logger
)

func init() {
	logger = log.New(os.Stdout, "", 0)
}

func indicator(c color.Color, lev level.Level) string {
	return fmt.Sprintf("%s%s%s[%s] ", c, lev, color.RESET, timestamp())
}

func timestamp() string {
	return time.Now().Format(time.RFC3339)
}

func logging(lev level.Level, str string) {
	switch lev {
	case level.Info:
		logger.Printf("%s %s", indicator(color.CYAN, level.Info), str)
	case level.Fine:
		logger.Printf("%s %s", indicator(color.CYAN, level.Fine), str)
	case level.Error:
		logger.Printf("%s %s", indicator(color.RED, level.Error), str)
	case level.Warn:
		logger.Printf("%s %s", indicator(color.YELLOW, level.Warn), str)
	}
}

func SetOutput(writers ...io.Writer) {
	logger.SetOutput(io.MultiWriter(writers...))
}
