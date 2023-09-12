package repository

import (
	"context"
	"device-manager/internal/database"
	"device-manager/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestInsertOne(t *testing.T) {
	// var userCollection *mongo.Collection
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		// userCollection = mt.Coll
		ctx := context.Background()
		id := primitive.NewObjectID()
		mt.AddMockResponses(mtest.CreateSuccessResponse())
		db := mt.Client.Database("test")
		r := NewDeviceRepository(&database.DataBaseMongo{
			MDB: db,
		})

		dev, err := r.Create(ctx, &domain.Device{
			ID:          id,
		})
		assert.Nil(t, err)
		assert.Equal(t, domain.Device{
			ID:          id,
		}, dev)
	})
}
