package repository

import (
	"context"
	"device-manager/internal/database"
	"device-manager/internal/domain"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type EventRepository struct {
	CollectionName string
	MongoDB        *database.Mongo
}

func NewEventRepository(mdb *database.Mongo) *EventRepository {
	return &EventRepository{CollectionName: "events", MongoDB: mdb}
}

func (r *EventRepository) Create(ctx context.Context, event *domain.Event) error {
	eventsCollection := r.MongoDB.DB.Collection(r.CollectionName)
	_, err := eventsCollection.InsertOne(ctx, event)
	if err != nil {
		return err
	}
	return nil
}

func (r *EventRepository) Get(ctx context.Context, uuid string, begin, end time.Time) ([]domain.Event, error) {
	filter := bson.M{
		"device_id":  uuid,
		"created_at": bson.M{"$gte": begin, "$lte": end},
	}
	cursor, err := r.MongoDB.DB.Collection(r.CollectionName).Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var events []domain.Event
	if err = cursor.All(ctx, &events); err != nil {
		return nil, err
	}
	return events, nil
}
