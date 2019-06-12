package httplog

import (
	"context"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	lvl := zerolog.DebugLevel
	if level := os.Getenv("LOG_LEVEL"); len(level) > 0 {
		switch level {
		case "Info":
			lvl = zerolog.InfoLevel
		case "Warn":
			lvl = zerolog.WarnLevel
		case "Error":
			lvl = zerolog.ErrorLevel
		case "Fatal":
			lvl = zerolog.FatalLevel
		case "Panic":
			lvl = zerolog.PanicLevel
		case "Disabled":
			lvl = zerolog.Disabled
		}
	}
	log.Logger = zerolog.New(os.Stdout).With().Timestamp().Logger().Level(lvl)
}

// Debug ...
func Debug(ctx context.Context) *zerolog.Event {
	return hook(ctx, log.Debug())
}

// Info ...
func Info(ctx context.Context) *zerolog.Event {
	return hook(ctx, log.Info())
}

// Warn ...
func Warn(ctx context.Context) *zerolog.Event {
	return hook(ctx, log.Warn())
}

// Error ...
func Error(ctx context.Context) *zerolog.Event {
	return hook(ctx, log.Error())
}

// Fatal ...
func Fatal(ctx context.Context) *zerolog.Event {
	return hook(ctx, log.Fatal())
}

func hook(ctx context.Context, e *zerolog.Event) *zerolog.Event {
	reqID := GetRequestID(ctx)
	return e.Str("reqID", reqID)
}
