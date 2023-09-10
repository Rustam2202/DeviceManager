package repository

import (
	"context"
	"device-manager/internal/database"
	"device-manager/internal/domain"

	"go.mongodb.org/mongo-driver/bson"
)

type DeviceRepository struct {
	MongoDB *database.DataBaseMongo
}

func NewDeviceRepository(mdb *database.DataBaseMongo) *DeviceRepository {
	return &DeviceRepository{MongoDB: mdb}
}

func (r *DeviceRepository) Create(ctx context.Context, device *domain.Device) error {
	devicesCollection := r.MongoDB.MDB.Collection("devices")
	_, err := devicesCollection.InsertOne(ctx, device)
	if err != nil {
		return err
	}
	return nil
}

func (r *DeviceRepository) Get(ctx context.Context, uuid string) (*domain.Device, error) {
	var result domain.Device
	err := r.MongoDB.MDB.Collection("devices").FindOne(ctx, bson.M{"uuid": uuid}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *DeviceRepository) Update(ctx context.Context, device *domain.Device) error {
	// filter := bson.M{"_id": device.ID}
	update := bson.M{
		"$set": bson.M{
			"$or": []bson.M{
				{"language": device.Language},
				{"geolocation": device.Geolocation},
				{"email": device.Email},
			},
		},
	}

	// filter := bson.D{
	// 	{Key: "language", Value: device.Language},
	// 	{Key: "geolocation", Value: device.Geolocation},
	// 	{Key: "email", Value: device.Email},
	// }
	_, err := r.MongoDB.MDB.Collection("devices").UpdateByID(ctx, device.ID, update)
	if err != nil {
		return err
	}
	return nil
}

func (r *DeviceRepository) Delete(ctx context.Context, uuid string) error {
	_, err := r.MongoDB.MDB.Collection("devices").DeleteOne(ctx, bson.M{"uuid": uuid})
	if err != nil {
		return err
	}
	return nil
}
