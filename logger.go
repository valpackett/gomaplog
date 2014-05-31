package gomaplog

import (
	"io"
	"os"
	"time"
)

type Logger struct {
	Formatter Formatter
	Writer    io.Writer
	Host      string
}

func (logger *Logger) LogE(event LogEvent) error {
	bytes, err := logger.Formatter.Format(event)
	if err != nil {
		return err
	}
	_, err = logger.Writer.Write(bytes)
	if err != nil {
		return err
	}
	_, err = logger.Writer.Write([]byte{'\n'})
	if err != nil {
		return err
	}
	return nil
}

func StdoutLogger(formatter Formatter) *Logger {
	return &Logger{Formatter: formatter, Writer: os.Stdout}
}

func (logger *Logger) LogL(level LogLevel, message, long_message string, extras Extras) error {
	return logger.LogE(LogEvent{Level: level, Host: logger.Host, Message: message, LongMessage: long_message, Timestamp: time.Now(), Extras: extras})
}

func (logger *Logger) Log(level LogLevel, message string, extras Extras) error {
	return logger.LogL(level, message, "", extras)
}
