package log

import (
	"fmt"
)

//LogLevel represents one of 'debug/info/warn/error/fatal'
type LogLevel int64

const (
	DEBUG LogLevel = iota //Debug log level
	INFO  LogLevel = iota //Info log level
	WARN  LogLevel = iota //Warn log level
	ERROR LogLevel = iota //Error log level
	FATAL LogLevel = iota //Fatal log level
)

// Info logs at the INFO level
func (l *Logger) Info(v string)                            { l.doMsg(INFO, v) }
func (l *Logger) Infof(format string, args ...interface{}) { l.doMsg(INFO, fmt.Sprintf(format, args)) }

// Debug logs at the DEBUG level
func (l *Logger) Debug(v string)                            { l.doMsg(DEBUG, v) }
func (l *Logger) Debugf(format string, args ...interface{}) { l.doMsg(DEBUG, fmt.Sprintf(format, args)) }

// Warn logs at the WARN level
func (l *Logger) Warn(v string)                            { l.doMsg(WARN, v) }
func (l *Logger) Warnf(format string, args ...interface{}) { l.doMsg(WARN, fmt.Sprintf(format, args)) }

// Error logs at the ERROR level
func (l *Logger) Error(v string)                            { l.doMsg(ERROR, v) }
func (l *Logger) Errorf(format string, args ...interface{}) { l.doMsg(ERROR, fmt.Sprintf(format, args)) }

// Fatal logs at the FATAL level
func (l *Logger) Fatal(v string)                            { l.doMsg(FATAL, v) }
func (l *Logger) Fatalf(format string, args ...interface{}) { l.doMsg(FATAL, fmt.Sprintf(format, args)) }

// Log at an arbitrary level
func (l *Logger) Log(level LogLevel, v string) { l.doMsg(level, v) }
