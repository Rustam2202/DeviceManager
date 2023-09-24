package repository

import (
	"device-manager/internal/database"
	"device-manager/internal/domain"
	"errors"

	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type DeviceRepository struct {
	CollectionName string
	MongoDB        *database.DataBaseMongo
}

func NewDeviceRepository(mdb *database.DataBaseMongo) *DeviceRepository {
	return &DeviceRepository{CollectionName: "devices", MongoDB: mdb}
}

func (r *DeviceRepository) Create(ctx context.Context, device *domain.Device) (*domain.Device, error) {
	devicesCollection := r.MongoDB.MDB.Collection(r.CollectionName)
	result, err := devicesCollection.InsertOne(ctx, device)
	if err != nil {
		return nil, err
	}
	device.ID = result.InsertedID.(primitive.ObjectID)
	return device, nil
}

// func (r *DeviceRepository) GetByID(ctx context.Context, id primitive.ObjectID) (*domain.Device, error) {
// 	var result domain.Device
// 	err := r.MongoDB.MDB.
// 		Collection(r.CollectionName).
// 		FindOne(ctx, bson.M{"_id": id}).Decode(&result)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &result, nil
// }

func (r *DeviceRepository) GetByUUID(ctx context.Context, uuid string) (*domain.Device, error) {
	var result domain.Device
	err := r.MongoDB.MDB.
		Collection(r.CollectionName).
		FindOne(ctx, bson.M{"uuid": uuid}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *DeviceRepository) Update(ctx context.Context, device *domain.Device) error {
	update := bson.M{
		"language":    device.Language,
		"geolocation": device.Geolocation,
		"email":       device.Email,
	}
	result, err := r.MongoDB.MDB.
		Collection(r.CollectionName).
		UpdateOne(ctx, bson.M{"_id": device.ID}, bson.M{"$set": update})
	if err != nil {
		return err
	}
	if result.ModifiedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}

func (r *DeviceRepository) Delete(ctx context.Context, uuid string) error {
	result, err := r.MongoDB.MDB.
		Collection(r.CollectionName).
		DeleteOne(ctx, bson.M{"uuid": uuid})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("no document deleted")
	}
	return nil
}
