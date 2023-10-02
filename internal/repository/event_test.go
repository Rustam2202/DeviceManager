package repository

import (
	"context"
	"device-manager/internal/database"
	"device-manager/internal/domain"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestCreateEvent(t *testing.T) {
	ctx := context.Background()
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		eventsCollection = mt.Coll
		success := mtest.CreateSuccessResponse()
		mt.AddMockResponses(success)
		repo := NewEventRepository(&database.Mongo{
			DB: mt.DB,
		})
		err := repo.Create(ctx, &domain.Event{
			DeviceUUID: uuid.New().String(),
			Name:       "event name",
			CreatedAt:  time.Now(),
			Attributes: []interface{}{"text", 1, 0.99, true},
		})
		assert.Nil(t, err)
	})
	mt.Run("error", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{{Key: "error", Value: 0}})
		repo := NewEventRepository(&database.Mongo{
			DB: mt.DB,
		})
		err := repo.Create(ctx, &domain.Event{})
		assert.NotNil(t, err)
	})
}

func TestGetEvents(t *testing.T) {
	ctx := context.Background()
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		eventsCollection = mt.Coll
		uuid := uuid.New().String()
		find := mtest.CreateCursorResponse(1, "test.events", mtest.FirstBatch,
			bson.D{
				{Key: "_id", Value: primitive.NewObjectID()},
				{Key: "device_id", Value: uuid},
				{Key: "name", Value: "Event 1"},
				{Key: "created_at", Value: time.Now().Add(-5 * time.Second)},
				{Key: "attributes", Value: []interface{}{"text", 1, 0.99, true}},
			})
		getMore := mtest.CreateCursorResponse(1, "test.events", mtest.NextBatch,
			bson.D{
				{Key: "_id", Value: primitive.NewObjectID()},
				{Key: "device_id", Value: uuid},
				{Key: "name", Value: "Event 2"},
				{Key: "created_at", Value: time.Now().Add(-2 * time.Second)},
				{Key: "attributes", Value: []interface{}{"text 2", 2, 1.99, false}},
			})
		killCursors := mtest.CreateCursorResponse(0, "test.events", mtest.NextBatch)
		mt.AddMockResponses(find, getMore, killCursors)

		repo := NewEventRepository(&database.Mongo{
			DB: mt.DB,
		})

		events, err := repo.Get(ctx, uuid, time.Now().Add(-6*time.Second), time.Now())
		assert.NoError(t, err)
		assert.Len(t, events, 2)
		// assert.NotEmpty(t, events[0].ID)
		// assert.NotEmpty(t, events[1].ID)
		assert.Equal(t, "Event 1", events[0].Name)
		assert.Equal(t, "Event 2", events[1].Name)
		assert.True(t, events[0].CreatedAt.Before(events[1].CreatedAt))
	})

	mt.Run("error", func(mt *mtest.T) {
		eventsCollection = mt.Coll
		mt.AddMockResponses(bson.D{{Key: "error", Value: 0}})
		repo := NewEventRepository(&database.Mongo{
			DB: mt.DB,
		})
		response, err := repo.Get(ctx, uuid.New().String(), time.Now().Add(-2*time.Second), time.Now().Add(2*time.Second))
		assert.Nil(t, response)
		assert.NotNil(t, err)
	})
}
