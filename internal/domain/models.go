package domain

import (
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Device struct {
	// ID          primitive.ObjectID `bson:"_id,omitempty"`
	UUID        uuid.UUID `bson:"_id"`
	Platform    string    `bson:"platform"`
	Language    string    `bson:"language"`
	Geolocation []float64 `bson:"gocation"`
	Email       string    `bson:"email"`
}

type Event struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	DeviceId   primitive.ObjectID `bson:"device_id"`
	Name       string             `bson:"name"`
	CreatedAt  time.Time          `bson:"created_at"`
	Attributes []interface{}      `bson:"attributes"`
}
