package logger

import (
	"log"
	"os"
)

// Logger is a reusable logging utility
type Logger struct {
	logger *log.Logger
}

// NewLogger creates a new logger instance
func NewLogger(prefix string) *Logger {
	return &Logger{
		logger: log.New(os.Stdout, prefix, log.LstdFlags),
	}
}

// Info logs an info message
func (l *Logger) Info(msg string) {
	l.logger.Println("[INFO]", msg)
}

// Error logs an error message
func (l *Logger) Error(msg string) {
	l.logger.Println("[ERROR]", msg)
}
