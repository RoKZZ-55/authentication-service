package apiserver

import (
	"authentication-service/config"
	"authentication-service/internal/storage/storage"
	"authentication-service/pkg/db/mongodb"
	"context"

	"golang.org/x/exp/slog"
)

func Run(cfg *config.Config, log *slog.Logger) error {
	log.Info("init monogdb")

	ctx := context.Background()

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

	return nil
}
