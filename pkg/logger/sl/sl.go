package sl

import (
	"os"

	"golang.org/x/exp/slog"
)

const (
	levelDebug = "debug"
	levelInfo  = "info"
)

func Err(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}

func GetLogger(LogLevel string) *slog.Logger {
	var log *slog.Logger

	switch LogLevel {
	case levelDebug:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case levelInfo:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	default:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
		log.Error("config log level is incorrect", "default info log levl is used", levelInfo)
	}
	return log
}
