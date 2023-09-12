package mongodb

import (
	"authentication-service/config"
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ErrCreateClientMongodb = errors.New("failed to create client to mongodb due to error")

func New(ctx context.Context, cfg *config.MongoDB) (*mongo.Database, error) {
	dbURL := fmt.Sprintf("mongodb://%s:%s@%s:%s", cfg.User, cfg.Password, cfg.Host, cfg.Port)

	// Connect mongodb
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbURL))
	if err != nil {
		return nil, fmt.Errorf("%w %w", ErrCreateClientMongodb, err)
	}

	// Ping mongodb
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("%w %w", ErrCreateClientMongodb, err)
	}

	return client.Database(cfg.DBName), nil
}
