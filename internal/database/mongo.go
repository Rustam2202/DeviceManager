package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DataBaseMongo struct {
	MDB *mongo.Database
}

func NewMongo(cfg *MongoDbConfig) *DataBaseMongo {
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d", cfg.Host, cfg.Port))
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
	}
	// defer client.Disconnect(context.Background())

	err = client.Ping(context.Background(), nil)
	if err != nil {
	}
	db := client.Database(cfg.Name)
	return &DataBaseMongo{MDB: db}
}
