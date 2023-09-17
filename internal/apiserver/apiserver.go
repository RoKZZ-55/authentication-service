package apiserver

import (
	"authentication-service/config"
	"authentication-service/internal/handler"
	"authentication-service/internal/storage"
	"authentication-service/pkg/db/mongodb"
	"context"
	"log/slog"
	"net/http"
)

func Run(cfg *config.Config) error {
	slog.Info(
		"init monogdb",
		slog.String("Host", cfg.Host),
		slog.String("Port", cfg.Port),
		slog.String("DBName", cfg.DBName),
	)

	ctx := context.Background()

	// connect to mongodb and check connection
	db, err := mongodb.New(ctx, &cfg.MongoDB)
	if err != nil {
		return err
	}

	defer func() {
		if err := db.Client().Disconnect(ctx); err != nil {
			slog.Error("database disconnect error", err)
		}
	}()

	storage := storage.New(db, cfg)
	handler := handler.New(storage)

	// getting router and routes for server
	srv := New(handler)
	slog.Info(
		"server start",
		slog.String("BindAddr", cfg.BindAddr),
	)

	return http.ListenAndServe(cfg.BindAddr, srv)
}
