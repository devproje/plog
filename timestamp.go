package plog

import (
	"fmt"
	"time"
)

func (pl *Plog) timestamp() string {
	return fmt.Sprintf("[%s]", time.Now().Format(pl.TimestampFormat))
}

// DisableTimeStamp	turn off timestamp
func (pl *Plog) SetTimestamp(t bool) {
	pl.DisableTimestamp = !t
}

// SetTimeFormat set time format
// If you want a using custom formatter, please use TimeFormatter(f string) function
func (pl *Plog) SetTimeFormat(f string) {
	pl.TimestampFormat = f
}
