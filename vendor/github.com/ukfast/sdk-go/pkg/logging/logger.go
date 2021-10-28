package logging

import "fmt"

var logger Logger

// Logger defines a logger interface with methods common to most logging
// libraries. A shim will be required for logging libraries which don't implement
// all / any of these methods
type Logger interface {
	Error(string)
	Warn(string)
	Info(string)
	Debug(string)
	Trace(string)
}

func SetLogger(l Logger) {
	logger = l
}

func Error(msg string) {
	if logger != nil {
		logger.Error(msg)
	}
}

func Errorf(msg string, v ...interface{}) {
	Error(fmt.Sprintf(msg, v...))
}

func Warn(msg string) {
	if logger != nil {
		logger.Warn(msg)
	}
}

func Warnf(msg string, v ...interface{}) {
	Warn(fmt.Sprintf(msg, v...))
}

func Info(msg string) {
	if logger != nil {
		logger.Info(msg)
	}
}

func Infof(msg string, v ...interface{}) {
	Info(fmt.Sprintf(msg, v...))
}

func Debug(msg string) {
	if logger != nil {
		logger.Debug(msg)
	}
}

func Debugf(msg string, v ...interface{}) {
	Debug(fmt.Sprintf(msg, v...))
}

func Trace(msg string) {
	if logger != nil {
		logger.Trace(msg)
	}
}

func Tracef(msg string, v ...interface{}) {
	Trace(fmt.Sprintf(msg, v...))
}
