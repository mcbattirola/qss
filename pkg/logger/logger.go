package logger

import (
	"fmt"
	"log"
	"os"
)

// Level is the level of the messages
// that should be logged
type Level uint8

var logLevel Level = InfoLevel

var fatalLog = log.New(os.Stderr, "[FATAL] ", log.Llongfile|log.Ltime|log.Lmsgprefix)
var errorLog = log.New(os.Stderr, "[ERROR] ", log.Llongfile|log.Ltime|log.Lmsgprefix)
var warningLog = log.New(os.Stderr, "[WARN] ", log.Llongfile|log.Ltime|log.Lmsgprefix)
var infoLog = log.New(os.Stderr, "[INFO] ", log.Lshortfile|log.Ltime|log.Lmsgprefix)
var debugLog = log.New(os.Stderr, "[DEBUG] ", log.Lshortfile|log.Ltime|log.Lmsgprefix)

const (
	// Fatal is for logging critical failures. Application will exit.
	FatalLevel Level = iota
	// ErrorLevel is for important error logs.
	ErrorLevel
	// WarnLevel is for relevant warning logs.
	WarnLevel
	// InfoLevel is for informational level logs.
	InfoLevel
	// DebugLevel is for development and debugging logs.
	DebugLevel
)

// SetLogLevel sets the log level that will be printed.
// Level is global for all the aplication.
func SetLogLevel(level Level) {
	logLevel = level
}

// Debug prints the message to stderr only in debug mode
func Debug(format string, a ...any) {
	if logLevel >= DebugLevel {
		debugLog.Output(2, fmt.Sprintf(format, a...))
	}
}

// Info prints the message to stderr
func Info(format string, a ...any) {
	if logLevel >= InfoLevel {
		infoLog.Output(2, fmt.Sprintf(format, a...))
	}
}

// Warn prints the message to stderr
func Warn(format string, a ...any) {
	if logLevel >= WarnLevel {
		warningLog.Output(2, fmt.Sprintf(format, a...))
	}
}

// Error prints the message to stderr
func Error(format string, a ...any) {
	if logLevel >= ErrorLevel {
		errorLog.Output(2, fmt.Sprintf(format, a...))
	}
}

// Fatal prints the message to stderr and terminates the process
func Fatal(format string, a ...any) {
	if logLevel >= FatalLevel {
		fatalLog.Output(2, fmt.Sprintf(format, a...))
		os.Exit(1)
	}
}
