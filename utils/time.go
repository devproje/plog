package utils

import "strings"

// TimeFormatter string time custom formatter
func TimeFormatter(f string) string {
	year := "2006"
	month := "01"
	day := "02"
	hour := "15"
	minute := "04"
	second := "05"
	a := "PM"
	timezone := "Z07:00"

	f = strings.ReplaceAll(f, "YYYY", year)
	f = strings.ReplaceAll(f, "MM", month)
	f = strings.ReplaceAll(f, "DD", day)
	f = strings.ReplaceAll(f, "hh", hour)
	f = strings.ReplaceAll(f, "mm", minute)
	f = strings.ReplaceAll(f, "ss", second)
	f = strings.ReplaceAll(f, "a", a)
	f = strings.ReplaceAll(f, "ZZZ", timezone)

	return f
}
