package loggers

import (
	"log/slog"
	"os"
)

// Logger is a struct for logging
type Logger struct {
	*slog.Logger
}

// NewLogger returns a new Logger
func NewLogger() *Logger {
	return &Logger{
		Logger: slog.New(slog.NewTextHandler(os.Stdout, nil)),
	}
}

// Info logs an info message
func (l *Logger) Info(message string, args ...any) {
	l.Logger.Info(message, args...)
}

// Error logs an error message
func (l *Logger) Error(message string, args ...any) {
	l.Logger.Error(message, args...)
}
