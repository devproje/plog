package level

type Level int

const (
	Panic Level = iota
	Fatal
	Error
	Warn
	Info
	Debug
	Trace
)
