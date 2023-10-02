package consumer

import (
	"context"
	"device-manager/internal/logger"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

type coordinates struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type deviceCreateMessage struct {
	UUID        uuid.UUID   `json:"uuid"`
	Platform    string      `json:"platform"`
	Language    string      `json:"language"`
	Coordinates coordinates `json:"coordinates"`
	Email       string      `json:"email"`
}

func (r *KafkaConsumer) deviceCreate(ctx context.Context, msg kafka.Message) error {
	var req deviceCreateMessage
	err := json.Unmarshal(msg.Value, &req)
	if err != nil {
		logger.Logger.Error("Failed to unmarshal kafka message.", zap.Error(err))
		return err
	}
	if err := r.deviceService.Create(ctx, req.UUID, req.Platform, req.Language, req.Email,
		[]float64{req.Coordinates.Longitude, req.Coordinates.Latitude}); err != nil {
		logger.Logger.Error("Failed to add Device to db: ", zap.Error(err))
		return err
	}
	return nil
}

type deviceUpdateLanguageMessage struct {
	UUID     string `json:"uuid"`
	Language string `json:"language"`
}

func (r *KafkaConsumer) deviceUpdateLanguage(ctx context.Context, msg kafka.Message) error {
	var req deviceUpdateLanguageMessage
	err := json.Unmarshal(msg.Value, &req)
	if err != nil {
		logger.Logger.Error("Failed to unmarshal kafka message.", zap.Error(err))
		return err
	}
	if err := r.deviceService.UpdateLanguage(ctx, req.UUID, req.Language); err != nil {
		logger.Logger.Error("Failed to update Device language in db: ", zap.Error(err))
		return err
	}
	return nil
}

type deviceUpdateGeolocationMessage struct {
	UUID      string  `json:"uuid"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

func (r *KafkaConsumer) deviceUpdateGeoposition(ctx context.Context, msg kafka.Message) error {
	var req deviceUpdateGeolocationMessage
	err := json.Unmarshal(msg.Value, &req)
	if err != nil {
		logger.Logger.Error("Failed to unmarshal kafka message.", zap.Error(err))
		return err
	}
	if err := r.deviceService.UpdateGeolocation(ctx, req.UUID, []float64{req.Longitude, req.Latitude}); err != nil {
		logger.Logger.Error("Failed to update Geoposition in db: ", zap.Error(err))
		return err
	}
	return nil
}

type deviceUpdateEmailMessage struct {
	UUID  string `json:"uuid"`
	Email string `json:"email"`
}

func (r *KafkaConsumer) deviceUpdateEmail(ctx context.Context, msg kafka.Message) error {
	var req deviceUpdateEmailMessage
	err := json.Unmarshal(msg.Value, &req)
	if err != nil {
		logger.Logger.Error("Failed to unmarshal kafka message.", zap.Error(err))
		return err
	}
	if err := r.deviceService.UpdateEmail(ctx, req.UUID, req.Email); err != nil {
		logger.Logger.Error("Failed to update Email in db: ", zap.Error(err))
		return err
	}
	return nil
}

func (r *KafkaConsumer) deviceDelete(ctx context.Context, msg kafka.Message) error {
	if err := r.deviceService.Delete(ctx, string(msg.Value)); err != nil {
		logger.Logger.Error("Failed to delete Peson from db: ", zap.Error(err))
		return err
	}
	return nil
}
