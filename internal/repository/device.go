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
	MongoDB        *database.Mongo
}

func NewDeviceRepository(mdb *database.Mongo) *DeviceRepository {
	return &DeviceRepository{CollectionName: "devices", MongoDB: mdb}
}

func (r *DeviceRepository) Create(ctx context.Context, device *domain.Device) error {
	devicesCollection := r.MongoDB.DB.Collection(r.CollectionName)
	_, err := devicesCollection.InsertOne(ctx, device)
	if err != nil {
		return err
	}
	return nil
}

func (r *DeviceRepository) Get(ctx context.Context, uuid uuid.UUID) (*domain.Device, error) {
	var result domain.Device
	err := r.MongoDB.DB.Collection(r.CollectionName).
		FindOne(ctx, bson.M{"_id": uuid}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *DeviceRepository) GetByLanguage(ctx context.Context, lang string) ([]domain.Device, error) {
	filter := bson.M{
		"language": lang,
	}
	cursor, err := r.MongoDB.DB.Collection(r.CollectionName).Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var devices []domain.Device
	for cursor.Next(ctx) {
		var device domain.Device
		err := cursor.Decode(&device)
		if err != nil {
			return nil, err
		}
		devices = append(devices, device)
	}
	return devices, nil
}

func (r *DeviceRepository) GetByGeolocation(ctx context.Context, long, lat, distance float64) ([]domain.Device, error) {
	location := bson.D{
		{Key: "type", Value: "Point"},
		{Key: "coordinates", Value: []float64{long, lat}}}
	filter := bson.D{
		{Key: "location", Value: bson.D{
			{Key: "$nearSphere", Value: bson.D{
				{Key: "$geometry", Value: location},
				{Key: "maxDistance", Value: distance},
			}},
		}},
	}
	cursor, err := r.MongoDB.DB.Collection(r.CollectionName).Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var devices []domain.Device
	err = cursor.All(ctx, devices)
	if err != nil {
		return nil, err
	}

	// for cursor.Next(ctx) {
	// 	var device domain.Device
	// 	err := cursor.Decode(&device)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	devices = append(devices, device)
	// }
	return devices, nil
}

func (r *DeviceRepository) GetByEmail(ctx context.Context, email string) ([]domain.Device, error) {
	filter := bson.M{
		"email": email,
	}
	cursor, err := r.MongoDB.DB.Collection(r.CollectionName).Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var devices []domain.Device
	for cursor.Next(ctx) {
		var device domain.Device
		err := cursor.Decode(&device)
		if err != nil {
			return nil, err
		}
		devices = append(devices, device)
	}
	return devices, nil
}

func (r *DeviceRepository) UpdateLanguage(ctx context.Context, uuid uuid.UUID, lang string) error {
	result, err := r.MongoDB.DB.Collection(r.CollectionName).
		UpdateOne(ctx, bson.M{"_id": uuid}, bson.M{"$set": bson.M{"language": lang}})
	if err != nil {
		return err
	}
	if result.ModifiedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}

func (r *DeviceRepository) UpdateGeolocation(ctx context.Context, uuid uuid.UUID, coordinates []float64) error {
	result, err := r.MongoDB.DB.Collection(r.CollectionName).
		UpdateOne(ctx, bson.M{"_id": uuid}, bson.M{"$set": bson.M{"location": domain.Location{
			Type:        "Point",
			Coordinates: coordinates,
		}}})
	if err != nil {
		return err
	}
	if result.ModifiedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}

func (r *DeviceRepository) UpdateEmail(ctx context.Context, uuid uuid.UUID, email string) error {
	result, err := r.MongoDB.DB.Collection(r.CollectionName).
		UpdateOne(ctx, bson.M{"_id": uuid}, bson.M{"$set": bson.M{"email": email}})
	if err != nil {
		return err
	}
	if result.ModifiedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}

func (r *DeviceRepository) Delete(ctx context.Context, uuid uuid.UUID) error {
	result, err := r.MongoDB.DB.Collection(r.CollectionName).
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
