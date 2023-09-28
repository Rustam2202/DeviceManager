package domain

import (
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Location struct {
	Type        string    `bson:"type"`
	Coordinates []float64 `bson:"coordinates"`
}

type Device struct {
	// ID          primitive.ObjectID `bson:"_id,omitempty"`
	UUID     string `bson:"_id"`
	Platform string    `bson:"platform"`
	Language string    `json:"language"`
	Location Location  `bson:"location"`
	Email    string    `bson:"email"`
}

type Event struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	DeviceUUID uuid.UUID          `bson:"device_id"`
	Name       string             `bson:"name"`
	CreatedAt  time.Time          `bson:"created_at"`
	Attributes []interface{}      `bson:"attributes"`
}
