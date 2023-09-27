package repository

import (
	"context"
	"device-manager/internal/database"
	"device-manager/internal/domain"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"golang.org/x/text/language"
)

var eventsCollection *mongo.Collection

func TestCreateDevice(t *testing.T) {
	ctx := context.Background()
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	mt.Run("success", func(mt *mtest.T) {
		eventsCollection = mt.Coll
		mt.AddMockResponses(mtest.CreateSuccessResponse())
		repo := NewDeviceRepository(&database.Mongo{DB: mt.DB})
		err := repo.Create(ctx, &domain.Device{UUID: uuid.New()})
		assert.Nil(t, err)
	})
	mt.Run("error", func(mt *mtest.T) {
		eventsCollection = mt.Coll
		mt.AddMockResponses(bson.D{{Key: "error", Value: 0}})
		repo := NewDeviceRepository(&database.Mongo{DB: mt.DB})
		err := repo.Create(ctx, &domain.Device{UUID: uuid.New()})
		assert.Error(t, err)
	})
	mt.Run("empty uuid", func(mt *mtest.T) {
		eventsCollection = mt.Coll
		mt.AddMockResponses(bson.D{{Key: "error", Value: 0}})
		repo := NewDeviceRepository(&database.Mongo{DB: mt.DB})
		err := repo.Create(ctx, &domain.Device{UUID: uuid.New()})
		assert.Error(t, err)
	})
	mt.Run("dublicate", func(mt *mtest.T) {
		eventsCollection = mt.Coll
		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Index:   1,
			Code:    11000,
			Message: "duplicate key error",
		}))
		repo := NewDeviceRepository(&database.Mongo{DB: mt.DB})
		err := repo.Create(ctx, &domain.Device{UUID: uuid.New()})
		assert.NotNil(t, err)
		assert.True(t, mongo.IsDuplicateKeyError(err))
	})
}

func TestGet(t *testing.T) {
	ctx := context.Background()
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		eventsCollection = mt.Coll
		uuid := uuid.New()
		expect := domain.Device{
			UUID:     uuid,
			Platform: "mac",
			Language: language.BritishEnglish,
			Location: domain.Location{},
			Email:    "test@email.com",
		}
		mt.AddMockResponses(mtest.CreateCursorResponse(1, "test.devices", mtest.FirstBatch, bson.D{
			{Key: "uuid", Value: expect.UUID},
			{Key: "platform", Value: expect.Platform},
			{Key: "language", Value: expect.Language},
			{Key: "geolocation", Value: expect.Location},
			{Key: "email", Value: expect.Email},
		}))
		repo := NewDeviceRepository(&database.Mongo{
			DB: mt.DB,
		})
		response, err := repo.Get(ctx, uuid)
		assert.Nil(t, err)
		// assert.Equal(t, expect.UUID, response.UUID)
		assert.Equal(t, expect.Platform, response.Platform)
		assert.Equal(t, expect.Language, response.Language)
		assert.Equal(t, expect.Email, response.Email)
	})
	mt.Run("error", func(mt *mtest.T) {
		eventsCollection = mt.Coll
		mt.AddMockResponses(bson.D{{Key: "error", Value: 0}})
		repo := NewDeviceRepository(&database.Mongo{
			DB: mt.DB,
		})
		response, err := repo.Get(ctx, uuid.New())
		assert.Nil(t, response)
		assert.NotNil(t, err)
	})
}

func TestUpdate(t *testing.T) {
	ctx := context.Background()
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success language update", func(mt *mtest.T) {
		eventsCollection = mt.Coll
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
		repo := NewDeviceRepository(&database.Mongo{
			DB: mt.DB,
		})
		err := repo.UpdateLanguage(ctx, uuid.New(), "eng")
		assert.Nil(t, err)
	})
	mt.Run("error", func(mt *mtest.T) {
		eventsCollection = mt.Coll
		mt.AddMockResponses(bson.D{{Key: "error", Value: 0}})
		repo := NewDeviceRepository(&database.Mongo{
			DB: mt.DB,
		})
		err := repo.UpdateLanguage(ctx, uuid.New(), "eng")
		assert.NotNil(t, err)
	})
	mt.Run("not modified", func(mt *mtest.T) {
		eventsCollection = mt.Coll
		mockResult := mongo.UpdateResult{
			MatchedCount:  1,
			ModifiedCount: 0,
		}
		mockResponse := bson.D{
			{Key: "ok", Value: 1},
			{Key: "n", Value: mockResult.MatchedCount},
			{Key: "nModified", Value: mockResult.ModifiedCount},
		}
		mt.AddMockResponses(mockResponse)
		repo := NewDeviceRepository(&database.Mongo{
			DB: mt.DB,
		})
		err := repo.UpdateLanguage(ctx, uuid.New(), "eng")
		assert.NotNil(t, err)
		assert.Error(t, mongo.ErrNoDocuments)
	})
}

func TestDelete(t *testing.T) {
	ctx := context.Background()
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		eventsCollection = mt.Coll
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "acknowledged", Value: true},
			{Key: "n", Value: 1},
		})
		repo := NewDeviceRepository(&database.Mongo{
			DB: mt.DB,
		})
		err := repo.Delete(ctx, uuid.New())
		assert.Nil(t, err)
	})
	mt.Run("error", func(mt *mtest.T) {
		eventsCollection = mt.Coll
		mt.AddMockResponses(bson.D{{Key: "error", Value: 0}})
		repo := NewDeviceRepository(&database.Mongo{
			DB: mt.DB,
		})
		err := repo.Delete(ctx, uuid.New())
		assert.NotNil(t, err)
	})
	mt.Run("no document deleted", func(mt *mtest.T) {
		eventsCollection = mt.Coll
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "acknowledged", Value: true},
			{Key: "n", Value: 0},
		})
		repo := NewDeviceRepository(&database.Mongo{
			DB: mt.DB,
		})
		err := repo.Delete(ctx, uuid.New())
		assert.NotNil(t, err)
	})
}
