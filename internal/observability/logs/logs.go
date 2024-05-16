//go:generate go-enum --marshal --lower --names --values --ptr

// Package logs ...
package logs

import (
	"context"
	"io"
	"log/slog"
	"os"

	"github.com/seanhagen/endless_stream/internal/observability"
)

const (
	LevelDebug = slog.LevelDebug
	LevelInfo  = slog.LevelInfo
	LevelWarn  = slog.LevelWarn
	LevelError = slog.LevelError
)

// LogType is the type used to switch the style of the logger. Options
// are text or JSON.
//
// ENUM(Text,JSON).
type LogType int

type Config struct {
	Out       io.Writer
	LogType   LogType
	Level     slog.Level
	AddSource bool
}

// handler ...
func (c *Config) handler() slog.Handler {
	opts := slog.HandlerOptions{
		AddSource: c.AddSource,
		Level:     c.Level,
	}

	switch c.LogType {
	case LogTypeText:
		return slog.NewTextHandler(c.Out, &opts)
	case LogTypeJSON:
		return slog.NewJSONHandler(c.Out, &opts)
	}

	return slog.Default().Handler()
}

func New(conf *Config) observability.Logger {
	if conf == nil {
		conf = &Config{
			Out:       os.Stdout,
			LogType:   LogTypeText,
			Level:     slog.LevelInfo,
			AddSource: false,
		}
	}

	return &logger{
		output: conf.Out,
		slog:   slog.New(conf.handler()),
	}
}

func DiscardLogger() observability.Logger {
	conf := &Config{
		Out: io.Discard,
	}
	return &logger{
		output: io.Discard,
		slog:   slog.New(conf.handler()),
	}
}

type logger struct {
	output io.Writer
	slog   *slog.Logger
}

// GetOutput ...
func (l *logger) GetOutput() io.Writer {
	return l.output
}

// Debug ...
func (l *logger) Debug(ctx context.Context, msg string, attrs ...slog.Attr) {
	l.slog.LogAttrs(ctx, slog.LevelDebug, msg, attrs...)
}

// Info ...
func (l *logger) Info(ctx context.Context, msg string, attrs ...slog.Attr) {
	l.slog.LogAttrs(ctx, slog.LevelInfo, msg, attrs...)
}

// Error ...
func (l *logger) Error(ctx context.Context, msg string, attrs ...slog.Attr) {
	l.slog.LogAttrs(ctx, slog.LevelError, msg, attrs...)
}

// Warn ...
func (l *logger) Warn(ctx context.Context, msg string, attrs ...slog.Attr) {
	l.slog.LogAttrs(ctx, slog.LevelWarn, msg, attrs...)
}

// WithAttrs ...
func (l *logger) WithAttrs(attrs ...slog.Attr) observability.Logger {
	return l
}

// WithGroup ...
func (l *logger) WithGroup(name string) observability.Logger {
	return l
}
