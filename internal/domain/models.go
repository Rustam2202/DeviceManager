package domain

import (
	"time"
)

type Location struct {
	Type        string    `bson:"type"`
	Coordinates []float64 `bson:"coordinates"`
}

type Device struct {
	UUID     string   `bson:"_id"`
	Platform string   `bson:"platform"`
	Language string   `json:"language"`
	Location Location `bson:"location"`
	Email    string   `bson:"email"`
}

type Event struct {
	DeviceUUID string        `bson:"device_id"`
	Name       string        `bson:"name"`
	CreatedAt  time.Time     `bson:"created_at"`
	Attributes []interface{} `bson:"attributes"`
}
