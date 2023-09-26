package repository

import (
	"context"
	"device-manager/internal/database"
	"device-manager/internal/domain"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EventRepository struct {
	CollectionName string
	MongoDB        *database.DataBaseMongo
}

func NewEventRepository(mdb *database.DataBaseMongo) *EventRepository {
	return &EventRepository{CollectionName: "events", MongoDB: mdb}
}

func (r *EventRepository) Create(ctx context.Context, event *domain.Event) error {
	eventsCollection := r.MongoDB.MDB.Collection(r.CollectionName)
	result, err := eventsCollection.InsertOne(ctx, event)
	if err != nil {
		return err
	}
	event.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (r *EventRepository) Get(ctx context.Context, deviceId primitive.ObjectID, begin, end time.Time) ([]domain.Event, error) {
	filter := bson.M{
		"device_id":  deviceId,
		"created_at": bson.M{"$gte": begin, "$lte": end},
	}
	cursor, err := r.MongoDB.MDB.Collection(r.CollectionName).Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var events []domain.Event
	for cursor.Next(ctx) {
		var event domain.Event
		err := cursor.Decode(&event)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}
