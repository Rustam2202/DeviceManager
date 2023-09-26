package repository

import (
	"context"
	"device-manager/internal/database"
	"device-manager/internal/domain"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DeviceRepository struct {
	CollectionName string
	MongoDB        *database.DataBaseMongo
}

func NewDeviceRepository(mdb *database.DataBaseMongo) *DeviceRepository {
	return &DeviceRepository{CollectionName: "devices", MongoDB: mdb}
}

func (r *DeviceRepository) Create(ctx context.Context, device *domain.Device) error {
	devicesCollection := r.MongoDB.MDB.Collection(r.CollectionName)
	_, err := devicesCollection.InsertOne(ctx, device)
	if err != nil {
		return err
	}
	return nil
}

func (r *DeviceRepository) Get(ctx context.Context, uuid string) (*domain.Device, error) {
	var result domain.Device
	err := r.MongoDB.MDB.Collection(r.CollectionName).
		FindOne(ctx, bson.M{"_id": uuid}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *DeviceRepository) UpdateLanguage(ctx context.Context, uuid uuid.UUID, lang string) error {
	result, err := r.MongoDB.MDB.Collection(r.CollectionName).
		UpdateOne(ctx, bson.M{"_id": uuid}, bson.M{"$set": bson.M{"language": lang}})
	if err != nil {
		return err
	}
	if result.ModifiedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}

func (r *DeviceRepository) UpdateGeolocation(ctx context.Context, uuid uuid.UUID, geo []float64) error {
	result, err := r.MongoDB.MDB.Collection(r.CollectionName).
		UpdateOne(ctx, bson.M{"_id": uuid}, bson.M{"$set": bson.M{"geolocation": geo}})
	if err != nil {
		return err
	}
	if result.ModifiedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}

func (r *DeviceRepository) UpdateEmail(ctx context.Context, uuid uuid.UUID, email string) error {
	result, err := r.MongoDB.MDB.Collection(r.CollectionName).
		UpdateOne(ctx, bson.M{"_id": uuid}, bson.M{"$set": bson.M{"email": email}})
	if err != nil {
		return err
	}
	if result.ModifiedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}

func (r *DeviceRepository) Delete(ctx context.Context, uuid string) error {
	result, err := r.MongoDB.MDB.Collection(r.CollectionName).
		DeleteOne(ctx, bson.M{"_id": uuid})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
		// return errors.New("no document deleted")
	}
	return nil
}
