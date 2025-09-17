package store

import (
	"context"

	"github.com/moweilong/mo/pkg/mlog"
)

// Logger is a logger that implements the Logger interface.
// It uses the log package to log error messages with additional context.
type Logger struct{}

// NewLogger creates and returns a new instance of Logger.
func NewLogger() *Logger {
	return &Logger{}
}

// Error logs an error message with the provided context using the log package.
func (l *Logger) Error(ctx context.Context, err error, msg string, kvs ...any) {
	kvs = append(kvs, err)
	mlog.W(ctx).Errorw(msg, kvs...)
}
