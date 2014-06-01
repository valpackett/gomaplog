package gomaplog

import (
	"time"
)

// The TermColors trick was stolen from https://github.com/spacemonkeygo/spacelog
type TermColors struct{}

func (TermColors) Reset() string     { return "\x1b[0m" }
func (TermColors) Bold() string      { return "\x1b[1m" }
func (TermColors) Underline() string { return "\x1b[4m" }
func (TermColors) Black() string     { return "\x1b[30m" }
func (TermColors) Red() string       { return "\x1b[31m" }
func (TermColors) Green() string     { return "\x1b[32m" }
func (TermColors) Yellow() string    { return "\x1b[33m" }
func (TermColors) Blue() string      { return "\x1b[34m" }
func (TermColors) Magenta() string   { return "\x1b[35m" }
func (TermColors) Cyan() string      { return "\x1b[36m" }
func (TermColors) White() string     { return "\x1b[37m" }

type Extras map[string]interface{}

type LogEvent struct {
	Level       LogLevel
	Host        string
	Message     string
	LongMessage string
	Timestamp   time.Time
	Extras      Extras
	TermColors
}

func (event *LogEvent) LevelColor() string {
	switch event.Level {
	case Emergency, Alert, Critical, Error:
		return event.Red() + event.Bold()
	case Warning, Notice:
		return event.Red()
	case Info, Debug:
		return event.Blue()
	}
	return ""
}

func (event *LogEvent) TimeRFC3339() string {
	return event.Timestamp.Format(time.RFC3339)
}
