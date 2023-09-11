package main

import (
	"authentication-service/config"
	"authentication-service/internal/apiserver"
	"authentication-service/pkg/logger/sl"

	"golang.org/x/exp/slog"
)

func main() {
	cfg := config.GetConfig()

	log := sl.GetLogger(cfg.LogLevel)

	log.Info(
		"starting authentication-service",
		slog.String("LogLevel", cfg.LogLevel),
	)

	if err := apiserver.Run(cfg, log); err != nil {
		log.Error("apiserver error:", sl.Err(err))
	}
}
