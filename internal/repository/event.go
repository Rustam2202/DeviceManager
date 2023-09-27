package repository

import (
	"context"
	"device-manager/internal/database"
	"device-manager/internal/domain"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	result, err := eventsCollection.InsertOne(ctx, event)
	if err != nil {
		return err
	}
	event.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (r *EventRepository) Get(ctx context.Context, uuid uuid.UUID, begin, end time.Time) ([]domain.Event, error) {
	filter := bson.M{
		"device_id":  uuid,
		"created_at": bson.M{"$gte": begin, "$lte": end},
	}
	cursor, err := r.MongoDB.DB.Collection(r.CollectionName).Find(ctx, filter)
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
