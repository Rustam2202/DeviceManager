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
		userCollection = mt.Coll
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
		userCollection = mt.Coll
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
		userCollection = mt.Coll
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
		userCollection = mt.Coll
		id := primitive.NewObjectID()
		deviceId := primitive.NewObjectID()
		loc, _ := time.LoadLocation("Europe/Moscow")
		createdTime := time.Now().Truncate(time.Second)
		expect := domain.Event{
			ID:         id,
			DeviceId:   deviceId,
			Name:       "event name",
			CreatedAt:  createdTime,
			Attributes: []interface{}{"text", 1, 0.99, true},
		}
		mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{Key: "_id", Value: expect.ID},
			{Key: "device_id", Value: expect.DeviceId},
			{Key: "name", Value: expect.Name},
			{Key: "created_at", Value: createdTime},
			{Key: "attributes", Value: expect.Attributes},
		}))
		repo := NewEventRepository(&database.DataBaseMongo{
			MDB: mt.DB,
		})
		response, err := repo.Get(ctx, expect.DeviceId, time.Now().Add(-3*time.Minute), time.Now().Add(3*time.Minute))
		assert.Nil(t, err)
		// assert.ElementsMatch(t,expect,response[0])
		// if !expect.CreatedAt.In(loc).Equal(response[0].CreatedAt.In(loc)) {
		// 	t.Error(expect.CreatedAt, response[0].CreatedAt)
		// }
		assert.EqualValues(t, expect.CreatedAt.In(loc), response[0].CreatedAt.In(loc))
		// assert.EqualValues(t, expect.Attributes, response[0].Attributes)
	})
	mt.Run("error", func(mt *mtest.T) {
		userCollection = mt.Coll
		mt.AddMockResponses(bson.D{{Key: "error", Value: 0}})
		repo := NewEventRepository(&database.DataBaseMongo{
			MDB: mt.DB,
		})
		response, err := repo.Get(ctx, primitive.NewObjectID(), time.Now().Add(-2*time.Second), time.Now().Add(2*time.Second))
		assert.Nil(t, response)
		assert.NotNil(t, err)
	})
}
