package utils

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"
)

var Logger *slog.Logger

type myHandler struct{}

func (handler myHandler) Enabled(ctx context.Context, level slog.Level) bool {
	switch level {
	case slog.LevelInfo:
		fallthrough
	case slog.LevelDebug:
		fallthrough
	case slog.LevelError:
		fallthrough
	case slog.LevelWarn:
		return true
	default:
		return false
	}
}

func (handler myHandler) Handle(ctx context.Context, record slog.Record) error {
	message := record.Message
	level := record.Level
	timeStamp := record.Time.Format(time.RFC1123Z)

	record.Attrs(func(a slog.Attr) bool {
		message += a.Key + ":" + a.Value.String()
		return true
	})

	switch level {
	case slog.LevelInfo:
		fallthrough
	case slog.LevelDebug:
		fallthrough
	case slog.LevelWarn:
		fmt.Fprintf(os.Stdout, "[%s] %s: %s\n", level, timeStamp, message)
	case slog.LevelError:
		fmt.Fprintf(os.Stderr, "[%s] %s: %s\n", level, timeStamp, message)
	default:
		fmt.Fprintf(os.Stderr, "this is not supposed to be logged")
	}
	return nil
}

func (handler myHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	panic("not implemented")
}

func (handler myHandler) WithGroup(name string) slog.Handler {
	panic("not implemented")
}

func NewLogger(tag string) {
	Logger = slog.New(myHandler{})
}
