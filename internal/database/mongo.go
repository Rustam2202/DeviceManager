package database

import (
	"device-manager/internal/logger"
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
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d/", cfg.Host, cfg.Port))
	clientOptions := options.Client().ApplyURI("mongodb://0.0.0.0:27017")
	client, err := mongo.Connect(ctx, clientOptions)
	// defer func() {
	// 	cancel()
	// 	if err := client.Disconnect(ctx); err != nil {
	// 		logger.Logger.Error("Disconnect error.", zap.Error(err))
	// 	}
	// }()
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
