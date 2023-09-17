package main

import (
	"authentication-service/config"
	"authentication-service/internal/apiserver"
	"authentication-service/pkg/logger/sl"
	"log/slog"
)

func main() {
	cfg := config.GetConfig()

	if err := sl.GetLogger(cfg.LogLevel); err != nil {
		slog.Error("failed to init the logger", sl.Err(err))
		return
	}

	slog.Info(
		"starting authentication-service",
		slog.String("LogLevel", cfg.LogLevel),
	)

	if err := apiserver.Run(cfg); err != nil {
		slog.Error("error in the http.ListenAndServe function in the apiserver package", sl.Err(err))
	}
}
