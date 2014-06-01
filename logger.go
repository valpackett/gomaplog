package gomaplog

import (
	"io"
	"os"
	"strings"
	"time"
)

type Logger struct {
	Formatter Formatter
	Writer    io.Writer
	Host      string
	MaxLevel  LogLevel
}

func (logger *Logger) LogE(event LogEvent) error {
	if event.Level <= logger.MaxLevel {
		bytes, err := logger.Formatter.Format(event)
		if err != nil {
			return err
		}
		_, err = logger.Writer.Write(append(bytes, '\n'))
		if err != nil {
			return err
		}
	}
	return nil
}

func StdoutLogger(formatter Formatter) *Logger {
	return &Logger{Formatter: formatter, Writer: os.Stdout, Host: "", MaxLevel: Debug}
}

func (logger *Logger) LogL(level LogLevel, message, long_message string, extras Extras) error {
	msgParts := strings.Split(message, "\n")
	longMsg := strings.Join(msgParts[1:], "\n")
	if long_message != "" && longMsg != "" {
		longMsg += "\n"
	}
	if long_message != "" {
		longMsg += long_message
	}
	return logger.LogE(LogEvent{Level: level, Host: logger.Host, Message: msgParts[0], LongMessage: longMsg, Timestamp: time.Now(), Extras: extras})
}

func (logger *Logger) Log(level LogLevel, message string, extras Extras) error {
	return logger.LogL(level, message, "", extras)
}

func (logger *Logger) EmergencyL(message, long_message string, extras Extras) error {
	return logger.LogL(Emergency, message, long_message, extras)
}

func (logger *Logger) Emergency(message string, extras Extras) error {
	return logger.LogL(Emergency, message, "", extras)
}

func (logger *Logger) AlertL(message, long_message string, extras Extras) error {
	return logger.LogL(Alert, message, long_message, extras)
}

func (logger *Logger) Alert(message string, extras Extras) error {
	return logger.LogL(Alert, message, "", extras)
}

func (logger *Logger) CriticalL(message, long_message string, extras Extras) error {
	return logger.LogL(Critical, message, long_message, extras)
}

func (logger *Logger) Critical(message string, extras Extras) error {
	return logger.LogL(Critical, message, "", extras)
}

func (logger *Logger) ErrorL(message, long_message string, extras Extras) error {
	return logger.LogL(Error, message, long_message, extras)
}

func (logger *Logger) Error(message string, extras Extras) error {
	return logger.LogL(Error, message, "", extras)
}

func (logger *Logger) WarningL(message, long_message string, extras Extras) error {
	return logger.LogL(Warning, message, long_message, extras)
}

func (logger *Logger) Warning(message string, extras Extras) error {
	return logger.LogL(Warning, message, "", extras)
}

func (logger *Logger) NoticeL(message, long_message string, extras Extras) error {
	return logger.LogL(Notice, message, long_message, extras)
}

func (logger *Logger) Notice(message string, extras Extras) error {
	return logger.LogL(Notice, message, "", extras)
}

func (logger *Logger) InfoL(message, long_message string, extras Extras) error {
	return logger.LogL(Info, message, long_message, extras)
}

func (logger *Logger) Info(message string, extras Extras) error {
	return logger.LogL(Info, message, "", extras)
}

func (logger *Logger) DebugL(message, long_message string, extras Extras) error {
	return logger.LogL(Debug, message, long_message, extras)
}

func (logger *Logger) Debug(message string, extras Extras) error {
	return logger.LogL(Debug, message, "", extras)
}
