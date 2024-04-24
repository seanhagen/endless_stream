// Package observability ...
package observability

import (
	"context"
	"io"
	"log/slog"
	"os"
	"time"
)

type Logger interface {
	Debug(context.Context, string, ...slog.Attr)
	Info(context.Context, string, ...slog.Attr)
	Error(context.Context, string, ...slog.Attr)
	Warn(context.Context, string, ...slog.Attr)
	WithAttrs(...slog.Attr) Logger
	WithGroup(string) Logger
}

type TestLogger interface {
	Logger

	GetOutput() io.Writer
}

type (
	Attr  = slog.Attr
	Value = slog.Value
)

func DurationValue(v time.Duration) Value {
	return slog.DurationValue(v)
}

func ToAttr(key string, value Value) Attr {
	return Attr{Key: key, Value: value}
}

func ErrorAttr(err error) Attr {
	return Attr{
		Key:   "error",
		Value: slog.AnyValue(err),
	}
}

func SigAttr(sig os.Signal) Attr {
	return Attr{
		Key:   "signal",
		Value: slog.StringValue(sig.String()),
	}
}

func DurationAttr(dur time.Duration) Attr {
	return Attr{
		Key:   "duration",
		Value: slog.DurationValue(dur),
	}
}

func TypeAttr(t string) Attr {
	return Attr{
		Key:   "type",
		Value: slog.StringValue(t),
	}
}
