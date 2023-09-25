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
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestCreateEvent(t *testing.T) {
	ctx := context.Background()
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		eventsCollection = mt.Coll
		deviceId := primitive.NewObjectID()
		cursor := mtest.CreateCursorResponse(1, "test.devices", mtest.FirstBatch, bson.D{
			{Key: "_id", Value: primitive.NewObjectID()},
			{Key: "uuid", Value: "test-uid"},
		})
		killCursors := mtest.CreateCursorResponse(0, "test.devices", mtest.NextBatch)
		success := mtest.CreateSuccessResponse()
		mt.AddMockResponses(cursor, killCursors, success)
		repo := NewEventRepository(&database.DataBaseMongo{
			MDB: mt.DB,
		})
		res, err := repo.Create(ctx, &domain.Event{
			DeviceId:   deviceId,
			Name:       "event name",
			CreatedAt:  time.Now(),
			Attributes: []interface{}{"text", 1, 0.99, true},
		})
		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.NotEmpty(t, res.ID)
	})
	mt.Run("error", func(mt *mtest.T) {
		eventsCollection = mt.Coll
		cursor := mtest.CreateCursorResponse(1, "test.devices", mtest.FirstBatch, bson.D{
			{Key: "_id", Value: primitive.NewObjectID()},
			{Key: "uuid", Value: "test-uid"},
		})
		killCursors := mtest.CreateCursorResponse(0, "test.devices", mtest.NextBatch)
		mt.AddMockResponses(cursor, killCursors, bson.D{{Key: "error", Value: 0}})
		repo := NewEventRepository(&database.DataBaseMongo{
			MDB: mt.DB,
		})
		res, err := repo.Create(ctx, &domain.Event{})
		assert.NotNil(t, err)
		assert.Nil(t, res)
	})
	mt.Run("not exist uuid", func(mt *mtest.T) {
		eventsCollection = mt.Coll
		mt.AddMockResponses(bson.D{{Key: "error", Value: 0}})
		repo := NewEventRepository(&database.DataBaseMongo{
			MDB: mt.DB,
		})
		res, err := repo.Create(ctx, &domain.Event{})
		assert.NotNil(t, err)
		assert.Nil(t, res)
	})
}

func TestGetEvents(t *testing.T) {
	ctx := context.Background()
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		eventsCollection = mt.Coll
		deviceId := primitive.NewObjectID()
		cursor := mtest.CreateCursorResponse(1, "test.devices", mtest.FirstBatch, bson.D{
			{Key: "_id", Value: primitive.NewObjectID()},
			{Key: "uuid", Value: "test-uid"},
		})
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
		mt.AddMockResponses(cursor, killCursors, find, getMore, killCursors)

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
		cursor := mtest.CreateCursorResponse(1, "test.devices", mtest.FirstBatch, bson.D{
			{Key: "_id", Value: primitive.NewObjectID()},
			{Key: "uuid", Value: "test-uid"},
		})
		killCursors := mtest.CreateCursorResponse(0, "test.events", mtest.NextBatch)
		mt.AddMockResponses(cursor, killCursors, bson.D{{Key: "error", Value: 0}})
		repo := NewEventRepository(&database.DataBaseMongo{
			MDB: mt.DB,
		})
		response, err := repo.Get(ctx, primitive.NewObjectID(), time.Now().Add(-2*time.Second), time.Now().Add(2*time.Second))
		assert.Nil(t, response)
		assert.NotNil(t, err)
	})
	mt.Run("not exist uuid", func(mt *mtest.T) {
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
