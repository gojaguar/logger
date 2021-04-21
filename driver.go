package logger

import (
	"io"
	"log"
)

// Driver represents a generic log printer.
type Driver interface {
	// Println prints a log and appends a new line.
	Println(v ...interface{})
}

var _ Driver = (*DriverStd)(nil)

// DriverStd uses the standard log.Logger as log printer.
type DriverStd struct {
	*log.Logger
}

// newStdDriver initializes a new DriverStd implementation using log.Logger.
// All arguments accepted by this function should match the rules described in log.New.
func newStdDriver(w io.Writer, prefix string, flag int) *DriverStd {
	l := log.New(w, prefix, flag)
	return &DriverStd{Logger: l}
}
