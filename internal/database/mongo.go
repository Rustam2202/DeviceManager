package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Mongo struct {
	DB *mongo.Database
}

func MustConnectMongo(ctx context.Context, cfg *MongoDbConfig) *Mongo {
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d/%s", cfg.Host, cfg.Port, cfg.Name))
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		// logger.Logger.Error("Failed to create connection to Mongo database.", zap.Error(err))
		// return nil, err
		panic(err.Error())
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		// logger.Logger.Error("Failed to connection database.", zap.Error(err))
		// return nil, err
		panic(err.Error())
	}
	db := client.Database(cfg.Name)
	return &Mongo{DB: db}
}
