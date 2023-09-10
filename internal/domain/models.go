package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Device struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	UUID        string             `bson:"uuid"`
	Platform    string             `bson:"platform"`
	Language    string             `bson:"language"`
	Geolocation string             `bson:"geolocation"`
	Email       string             `bson:"email"`
}

type Event struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	DeviceId   primitive.ObjectID `bson:"device_id"`
	Name       string             `bson:"name"`
	CreatedAt  time.Time          `bson:"created_at"`
	Attributes []interface{}      `bson:"attributes"`
}
