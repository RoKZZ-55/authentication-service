package main

import (
	"authentication-service/config"
	"authentication-service/pkg/logger/sl"
)

func main() {
	cfg := config.GetConfig()
	log := sl.GetLogger(cfg.LogLevel)
	log.Info(
		"starting authentication-service",
		"LogLevel", cfg.LogLevel,
	)
}
