package onex

import (
	"github.com/moweilong/mo/log"
)

// moLogger is a logger that implements the Logger interface.
// It uses the log package to log error messages with additional context.
type moLogger struct{}

// NewLogger creates and returns a new instance of moLogger.
func NewLogger() *moLogger {
	return &moLogger{}
}

// Error logs an error message with the provided context using the log package.
func (l *moLogger) Error(err error, msg string, kvs ...any) {
	log.Errorw(err, msg, kvs...)
}
