package database

import (
	"device-manager/internal/logger"

	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type DataBaseMongo struct {
	MDB *mongo.Database
}

func NewMongo(cfg *MongoDbConfig) (*DataBaseMongo, error) {
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d", cfg.Host, cfg.Port))
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		logger.Logger.Error("Failed to create connection to Mongo database.", zap.Error(err))
		return nil, err
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		logger.Logger.Error("Failed to connection database.", zap.Error(err))
		return nil, err
	}
	db := client.Database(cfg.Name)
	return &DataBaseMongo{MDB: db}, nil
}
