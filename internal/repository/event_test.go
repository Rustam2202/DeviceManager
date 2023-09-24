package repository

import (
	"context"
	"device-manager/internal/database"
	"device-manager/internal/domain"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestCreateEvent(t *testing.T) {
	ctx := context.Background()
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		eventsCollection = mt.Coll
		id := primitive.NewObjectID()
		deviceId := primitive.NewObjectID()
		mt.AddMockResponses(mtest.CreateSuccessResponse())
		repo := NewEventRepository(&database.DataBaseMongo{
			MDB: mt.DB,
		})
		err := repo.Create(ctx, &domain.Event{
			ID:         id,
			DeviceId:   deviceId,
			Name:       "event name",
			CreatedAt:  time.Now(),
			Attributes: []interface{}{"text", 1, 0.99, true},
		})
		assert.Nil(t, err)
	})
	mt.Run("error", func(mt *mtest.T) {
		eventsCollection = mt.Coll
		id := primitive.NewObjectID()
		deviceId := primitive.NewObjectID()
		mt.AddMockResponses(bson.D{{Key: "error", Value: 0}})
		repo := NewEventRepository(&database.DataBaseMongo{
			MDB: mt.DB,
		})
		err := repo.Create(ctx, &domain.Event{
			ID:         id,
			DeviceId:   deviceId,
			Name:       "event name",
			CreatedAt:  time.Now(),
			Attributes: []interface{}{"text", 1, 0.99, true},
		})
		assert.NotNil(t, err)
	})
	mt.Run("dublicate", func(mt *mtest.T) {
		eventsCollection = mt.Coll
		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Index:   1,
			Code:    11000,
			Message: "duplicate key error",
		}))
		repo := NewEventRepository(&database.DataBaseMongo{
			MDB: mt.DB,
		})
		err := repo.Create(ctx, &domain.Event{})
		assert.NotNil(t, err)
		assert.True(t, mongo.IsDuplicateKeyError(err))
	})
}

func TestGetEvents(t *testing.T) {
	ctx := context.Background()
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		eventsCollection = mt.Coll
		deviceId := primitive.NewObjectID()

		find := mtest.CreateCursorResponse(1, "test.events", mtest.FirstBatch,
			bson.D{
				{Key: "_id", Value: primitive.NewObjectID()},
				{Key: "device_id", Value: deviceId},
				{Key: "name", Value: "Event 1"},
				{Key: "created_at", Value: time.Now().Add(-5 * time.Second)},
				{Key: "attributes", Value: []interface{}{"text", 1, 0.99, true}},
			})
		getMore := mtest.CreateCursorResponse(1, "test.events", mtest.NextBatch,
			bson.D{
				{Key: "_id", Value: primitive.NewObjectID()},
				{Key: "device_id", Value: deviceId},
				{Key: "name", Value: "Event 2"},
				{Key: "created_at", Value: time.Now().Add(-2 * time.Second)},
				{Key: "attributes", Value: []interface{}{"text 2", 2, 1.99, false}},
			})
		killCursors := mtest.CreateCursorResponse(0, "test.events", mtest.NextBatch)
		mt.AddMockResponses(find, getMore, killCursors)

		repo := NewEventRepository(&database.DataBaseMongo{
			MDB: mt.DB,
		})

		events, err := repo.Get(ctx, deviceId, time.Now().Add(-6*time.Second), time.Now())
		assert.NoError(t, err)
		assert.Len(t, events, 2)
		assert.NotEmpty(t, events[0].ID)
		assert.NotEmpty(t, events[1].ID)
		assert.Equal(t, "Event 1", events[0].Name)
		assert.Equal(t, "Event 2", events[1].Name)
		assert.True(t, events[0].CreatedAt.Before(events[1].CreatedAt))
	})

	mt.Run("error", func(mt *mtest.T) {
		eventsCollection = mt.Coll
		mt.AddMockResponses(bson.D{{Key: "error", Value: 0}})
		repo := NewEventRepository(&database.DataBaseMongo{
			MDB: mt.DB,
		})
		response, err := repo.Get(ctx, primitive.NewObjectID(), time.Now().Add(-2*time.Second), time.Now().Add(2*time.Second))
		assert.Nil(t, response)
		assert.NotNil(t, err)
	})
}
