package storage

import "go.mongodb.org/mongo-driver/mongo"

type Storage struct {
	db                    *mongo.Database
	authenticationStorage *AuthenticationStorage
}

func New(db *mongo.Database) *Storage {
	return &Storage{
		db: db,
	}
}
