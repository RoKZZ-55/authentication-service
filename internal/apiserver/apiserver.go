package apiserver

import (
	"authentication-service/config"
	"authentication-service/internal/storage/storage"
	"authentication-service/pkg/db/mongodb"
	"context"
	"net/http"

	"golang.org/x/exp/slog"
)

func Run(cfg *config.Config, log *slog.Logger) error {
	log.Info(
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
			log.Error("database disconnect error", err)
		}
	}()

	storage.New(db)

	// getting router and routes for server
	srv := New(log)
	log.Info(
		"server start",
		slog.String("BindAddr", cfg.BindAddr),
	)

	return http.ListenAndServe(cfg.BindAddr, srv)
}
