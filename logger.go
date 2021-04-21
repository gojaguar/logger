package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

// Verbosity represents a verbosity level for the logger.
type Verbosity uint8

const (
	// VerbosityError only prints Logger.Error messages.
	VerbosityError Verbosity = 1
	// VerbosityWarn prints Logger.Error and Logger.Warn messages.
	VerbosityWarn Verbosity = 2
	// VerbosityInfo prints Logger.Error, Logger.Warn and Logger.Info messages.
	VerbosityInfo Verbosity = 3
	// VerbosityDebug prints messages for all levels.
	VerbosityDebug Verbosity = 4
)

// Logger has methods to print logging messages to different levels of importance.
type Logger interface {
	// Debug prints debug messages, usually used by developers.
	Debug(message string)
	// Info prints information messages, usually used to log important messages that are not necessarily an error.
	Info(message string)
	// Warn prints messages that are a warning, they should not compromise the execution of the program, but those
	// messages involve a certain risk for the system.
	Warn(message string)
	// Error prints messages that are severe and should be addressed as soon as possible.
	Error(message string)
}

// logger is a Logger implementation.
type logger struct {
	// Verbosity is the verbosity that the logger is allowed to print.
	Verbosity Verbosity
	// Prefix is the prefix of the logging message.
	// Example:
	//		2021/04/20 00:32:27 main.go:15 [PREFIX] [INFO] Hello, this is a message.
	Prefix string
	// Driver is in charge of delivering the actual log message to the operative system or third party services.
	Driver Driver
}

func (l *logger) Debug(message string) {
	if l.Verbosity < VerbosityDebug {
		return
	}
	l.log("DEBUG", message)
}

func (l *logger) Info(message string) {
	if l.Verbosity < VerbosityInfo {
		return
	}
	l.log("INFO", message)
}

func (l *logger) Warn(message string) {
	if l.Verbosity < VerbosityWarn {
		return
	}
	l.log("WARNING", message)
}

func (l *logger) Error(message string) {
	if l.Verbosity < VerbosityError {
		return
	}
	l.log("ERROR", message)
}

func (l *logger) log(tag string, message string) {
	_, fn, line, _ := runtime.Caller(2)

	l.Driver.Println(fmt.Sprintf("[ %s:%d ] [%s] %s", filepath.Base(fn), line, tag, message))
}

// NewLogger initializes a new Logger implementation.
// It can receive a set of options that allow customizing the underlying logger.
// 	If no options are provided, the following configuration will be used:
// 	* Verbosity: VerbosityDebug.
// 	* Driver: DriverStd with the following flags: log.Ldate | log.Ltime | log.LUTC
// 	* Prefix: [LOG]
func NewLogger(opts ...Option) Logger {
	l := logger{
		Verbosity: VerbosityDebug,
		Driver:    newStdDriver(os.Stdout, "[LOG] ", log.Ldate|log.Ltime|log.LUTC),
		Prefix:    "LOG",
	}
	for _, opt := range opts {
		l = opt(l)
	}
	return &l
}
