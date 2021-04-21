package logger

import (
	"fmt"
	"io"
)

type Option func(logger) logger

func WithStdDriver(writer io.Writer, flag int) Option {
	return func(l logger) logger {
		l.Driver = newStdDriver(writer, fmt.Sprintf("[%s] ", l.Prefix), flag)
		return l
	}
}
