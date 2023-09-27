package consumer

import (
	"context"
	"device-manager/internal/logger"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
	"golang.org/x/text/language"
)

type deviceCreateMessage struct {
	UUID        uuid.UUID    `json:"uuid"`
	Platform    string       `json:"platform"`
	Language    language.Tag `json:"language"`
	Coordinates []float64    `json:"coordinates"`
	Email       string       `json:"email"`
}

func (r *KafkaConsumer) deviceCreate(ctx context.Context, msg kafka.Message) error {
	var req deviceCreateMessage
	err := json.Unmarshal(msg.Value, &req)
	if err != nil {
		logger.Logger.Error("Failed to unmarshal kafka message.", zap.Error(err))
		return err
	}
	if err := r.deviceService.Create(ctx, req.UUID, req.Platform, req.Language, req.Email, req.Coordinates); err != nil {
		logger.Logger.Error("Failed to add Device to db: ", zap.Error(err))
		return err
	}
	return nil
}

type deviceUpdateLanguageMessage struct {
	UUID     uuid.UUID `json:"uuid"`
	Language string    `json:"language"`
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
	UUID        uuid.UUID `json:"uuid"`
	Coordinates []float64 `json:"coordinates"`
}

func (r *KafkaConsumer) deviceUpdateGeoposition(ctx context.Context, msg kafka.Message) error {
	var req deviceUpdateGeolocationMessage
	err := json.Unmarshal(msg.Value, &req)
	if err != nil {
		logger.Logger.Error("Failed to unmarshal kafka message.", zap.Error(err))
		return err
	}
	if err := r.deviceService.UpdateGeolocation(ctx, req.UUID, req.Coordinates); err != nil {
		logger.Logger.Error("Failed to update Geoposition in db: ", zap.Error(err))
		return err
	}
	return nil
}

type deviceUpdateEmailMessage struct {
	UUID  uuid.UUID `json:"uuid"`
	Email string    `json:"email"`
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
	u, err := uuid.Parse(string(msg.Value))
	if err != nil {
		return err
	}
	if err := r.deviceService.Delete(ctx, u); err != nil {
		logger.Logger.Error("Failed to delete Peson from db: ", zap.Error(err))
		return err
	}
	return nil
}
