package storage

import (
	"authentication-service/config"

	"go.mongodb.org/mongo-driver/mongo"
)

type Storage struct {
	db                    *mongo.Database
	cfg                   *config.Config
	authenticationStorage *AuthenticationStorage
}

func New(db *mongo.Database, cfg *config.Config) *Storage {
	return &Storage{
		db:  db,
		cfg: cfg,
	}
}
