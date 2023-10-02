package database

import (
	"context"
	"device-manager/internal/logger"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/zap"
)

type Mongo struct {
	DB *mongo.Database
}

func MustConnectMongo(ctx context.Context, cfg *MongoDbConfig) *Mongo {
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d/%s", cfg.Host, cfg.Port, cfg.Name))
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err.Error())
	}
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err.Error())
	}
	db := client.Database(cfg.Name)
	if err = db.CreateCollection(ctx, "devices"); err != nil {
		logger.Logger.Error("create collection error", zap.Error(err))
	}
	if err = db.CreateCollection(ctx, "events"); err != nil {
		logger.Logger.Error("create collection error", zap.Error(err))
	}
	devicesIndexes := []mongo.IndexModel{
		{Keys: bson.M{"language": 1}},
		{Keys: bson.M{"email": 1}},
		{Keys: bson.M{"location": "2dsphere"}},
	}
	if _, err = db.Collection("devices").Indexes().CreateMany(ctx, devicesIndexes); err != nil {
		logger.Logger.Error("create indexes error", zap.Error(err))
	}
	eventsIndexes := []mongo.IndexModel{
		{Keys: bson.M{"name": 1}},
		{Keys: bson.M{"created_at": 1}},
	}
	if _, err = db.Collection("events").Indexes().CreateMany(ctx, eventsIndexes); err != nil {
		logger.Logger.Error("create indexes error", zap.Error(err))
	}
	return &Mongo{DB: db}
}
