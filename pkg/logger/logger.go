package logger

import (
	"log/slog"
	"os"
)

// Logger wraps slog.Logger for consistent logging across the application
type Logger struct {
	*slog.Logger
}

// New creates a new logger instance
func New(level string) *Logger {
	var logLevel slog.Level
	switch level {
	case "debug":
		logLevel = slog.LevelDebug
	case "info":
		logLevel = slog.LevelInfo
	case "warn":
		logLevel = slog.LevelWarn
	case "error":
		logLevel = slog.LevelError
	default:
		logLevel = slog.LevelInfo
	}

	opts := &slog.HandlerOptions{
		Level: logLevel,
	}

	handler := slog.NewJSONHandler(os.Stdout, opts)
	logger := slog.New(handler)

	return &Logger{Logger: logger}
}

// WithService adds service name to all log entries
func (l *Logger) WithService(service string) *Logger {
	return &Logger{Logger: l.Logger.With("service", service)}
}

// WithRequestID adds request ID to log entries for tracing
func (l *Logger) WithRequestID(requestID string) *Logger {
	return &Logger{Logger: l.Logger.With("request_id", requestID)}
}
