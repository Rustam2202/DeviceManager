package database

import (
	"device-manager/internal/logger"
	"fmt"
	"time"

	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/zap"
)

type DataBaseMongo struct {
	MDB *mongo.Database
}

func NewMongo(cfg *MongoDbConfig) (*DataBaseMongo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d/%s", cfg.Host, cfg.Port, cfg.Name))
	client, err := mongo.Connect(ctx, clientOptions)
	defer func() {
		cancel()
		if err := client.Disconnect(ctx); err != nil {
			logger.Logger.Error("Disconnect error.", zap.Error(err))
		}
	}()
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
	return &DataBaseMongo{MDB: db}, nil
}
