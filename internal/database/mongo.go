package database

import (
	"device-manager/internal/logger"
	"fmt"

	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/zap"
)

type Mongo struct {
	DB *mongo.Database
}

func NewMongo(ctx context.Context, cfg *MongoDbConfig) (*Mongo, error) {
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d/%s", cfg.Host, cfg.Port, cfg.Name))
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		logger.Logger.Error("Failed to create connection to Mongo database.", zap.Error(err))
		return nil, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		logger.Logger.Error("Failed to connection database.", zap.Error(err))
		return nil, err
	}
	db := client.Database(cfg.Name)
	return &Mongo{DB: db}, nil
}
