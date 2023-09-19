package repository

import (
	"context"
	"device-manager/internal/database"
	"device-manager/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

var userCollection *mongo.Collection

func TestCreateDevice(t *testing.T) {
	ctx := context.Background()
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	mt.Run("success", func(mt *mtest.T) {
		userCollection = mt.Coll
		mt.AddMockResponses(mtest.CreateSuccessResponse())
		repo := NewDeviceRepository(&database.DataBaseMongo{
			MDB: mt.DB,
		})
		device, err := repo.Create(ctx, &domain.Device{
			UUID:        "test-uuid",
			Platform:    "mac",
			Language:    "en",
			Geolocation: "here",
			Email:       "test@email.com",
		})
		assert.Nil(t, err)
		assert.NotNil(t, device)
	})
	mt.Run("empty uuid", func(mt *mtest.T) {
		userCollection = mt.Coll
		mt.AddMockResponses(bson.D{{Key: "error", Value: 0}})
		repo := NewDeviceRepository(&database.DataBaseMongo{
			MDB: mt.DB,
		})
		device, err := repo.Create(ctx, &domain.Device{
			UUID: "",
		})
		assert.Nil(t, device)
		assert.NotNil(t, err)
	})
	mt.Run("dublicate", func(mt *mtest.T) {
		userCollection = mt.Coll
		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Index:   1,
			Code:    11000,
			Message: "duplicate key error",
		}))
		repo := NewDeviceRepository(&database.DataBaseMongo{
			MDB: mt.DB,
		})
		device, err := repo.Create(ctx, &domain.Device{})
		assert.Nil(t, device)
		assert.NotNil(t, err)
		assert.True(t, mongo.IsDuplicateKeyError(err))
	})
}

func TestGet(t *testing.T) {
	ctx := context.Background()
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		userCollection = mt.Coll
		id := primitive.NewObjectID()
		expect := domain.Device{
			ID:          id,
			UUID:        "test-uuid",
			Platform:    "mac",
			Language:    "en",
			Geolocation: "here",
			Email:       "test@email.com",
		}
		mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{Key: "_id", Value: expect.ID},
			{Key: "uuid", Value: expect.UUID},
			{Key: "platform", Value: expect.Platform},
			{Key: "language", Value: expect.Language},
			{Key: "geolocation", Value: expect.Geolocation},
			{Key: "email", Value: expect.Email},
		}))
		repo := NewDeviceRepository(&database.DataBaseMongo{
			MDB: mt.DB,
		})
		response, err := repo.GetByUUID(ctx, "test-uuid")
		assert.Nil(t, err)
		assert.Equal(t, expect, *response)
	})
	mt.Run("error", func(mt *mtest.T) {
		userCollection = mt.Coll
		mt.AddMockResponses(bson.D{{Key: "error", Value: 0}})
		repo := NewDeviceRepository(&database.DataBaseMongo{
			MDB: mt.DB,
		})
		response, err := repo.GetByUUID(ctx, "test-uuid")
		assert.Nil(t, response)
		assert.NotNil(t, err)
	})
}

func TestUpdate(t *testing.T) {
	ctx := context.Background()
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		userCollection = mt.Coll
		id := primitive.NewObjectID()
		device := domain.Device{
			ID:          id,
			UUID:        "test-uuid",
			Platform:    "mac",
			Language:    "en",
			Geolocation: "here",
			Email:       "test@email.com",
		}
		mockResult := mongo.UpdateResult{
			MatchedCount:  1,
			ModifiedCount: 1,
		}
		mockResponse := bson.D{
			{Key: "ok", Value: 1},
			{Key: "n", Value: mockResult.MatchedCount},
			{Key: "nModified", Value: mockResult.ModifiedCount},
		}
		mt.AddMockResponses(mockResponse)
		repo := NewDeviceRepository(&database.DataBaseMongo{
			MDB: mt.DB,
		})
		err := repo.Update(ctx, &device)
		assert.Nil(t, err)
	})
}

func TestDelete(t *testing.T) {
	ctx := context.Background()
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		userCollection = mt.Coll
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "acknowledged", Value: true},
			{Key: "n", Value: 1},
		})
		repo := NewDeviceRepository(&database.DataBaseMongo{
			MDB: mt.DB,
		})
		err := repo.Delete(ctx, "")
		assert.Nil(t, err)
	})

	mt.Run("no document deleted", func(mt *mtest.T) {
		userCollection = mt.Coll
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "acknowledged", Value: true},
			{Key: "n", Value: 0},
		})
		repo := NewDeviceRepository(&database.DataBaseMongo{
			MDB: mt.DB,
		})
		err := repo.Delete(ctx, "")
		assert.NotNil(t, err)
	})
}
