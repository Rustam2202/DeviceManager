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
	CollectionName       string
	DeviceCollectionName string
	MongoDB              *database.DataBaseMongo
}

func NewEventRepository(mdb *database.DataBaseMongo) *EventRepository {
	return &EventRepository{CollectionName: "events", DeviceCollectionName: "devices", MongoDB: mdb}
}

func (r *EventRepository) Create(ctx context.Context, event *domain.Event) error {
	// devicesCollection := r.MongoDB.MDB.Collection(r.DeviceCollectionName)
	// devce := devicesCollection.FindOne(ctx, bson.M{"_id": event.DeviceId})
	// if devce.Err() != nil {
	// 	return fmt.Errorf("no device exist with '%s' id", event.DeviceId)
	// }
	eventsCollection := r.MongoDB.MDB.Collection(r.CollectionName)
	_, err := eventsCollection.InsertOne(ctx, event)
	if err != nil {
		return err
	}
	return nil
}

func (r *EventRepository) Get(ctx context.Context, deviceId primitive.ObjectID, begin, end time.Time) ([]domain.Event, error) {
	// devicesCollection := r.MongoDB.MDB.Collection(r.DeviceCollectionName)
	// devce := devicesCollection.FindOne(ctx, bson.M{"_id": deviceId})
	// if devce.Err() != nil {
	// 	return nil, fmt.Errorf("no device exist with '%s' id", deviceId)
	// }
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

