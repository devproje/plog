package plog

import (
	"fmt"

	"github.com/devproje/plog/color"
)

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

func (pl *Plog) prefix(t tag) string {
	i := fmt.Sprintf("%s%s", t, color.RESET)
	if !pl.DisableTimestamp {
		i += pl.timestamp()
	}

	return i
}
