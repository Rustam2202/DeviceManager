package consumer

import (
	"context"
	"device-manager/internal/domain"
	"device-manager/internal/logger"
	"encoding/json"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

func (r *KafkaConsumer) deviceCreateServe(ctx context.Context, msg kafka.Message) error {
	req := domain.Device{}
	err := json.Unmarshal(msg.Value, &req)
	if err != nil {
		logger.Logger.Error("Failed to unmarshal kafka message.", zap.Error(err))
		return err
	}
	if err := r.deviceService.CreateDevice(ctx, req.UUID, req.Platform, req.Language, req.Geolocation, req.Email); err != nil {
		logger.Logger.Error("Failed to add Device to db: ", zap.Error(err))
		return err
	}
	return nil
}

func (r *KafkaConsumer) deviceUpdateServe(ctx context.Context, msg kafka.Message) error {
	req := domain.Device{}
	err := json.Unmarshal(msg.Value, &req)
	if err != nil {
		logger.Logger.Error("Failed to unmarshal kafka message.", zap.Error(err))
		return err
	}
	switch {
	case req.Language != "":
		if err := r.deviceService.UpdateLaguage(ctx, req.UUID, req.Language); err != nil {
			logger.Logger.Error("Failed to update Language in db: ", zap.Error(err))
			return err
		}
	case req.Geolocation != "":
		if err := r.deviceService.UpdateGeolocation(ctx, req.UUID, req.Geolocation); err != nil {
			logger.Logger.Error("Failed to update Geoposition in db: ", zap.Error(err))
			return err
		}
	case req.Email != "":
		if err := r.deviceService.UpdateEmail(ctx, req.UUID, req.Email); err != nil {
			logger.Logger.Error("Failed to update Email in db: ", zap.Error(err))
			return err
		}
	default:
		// logger.Logger.Error("Failed to add Peson to db: ", zap.Error(err))
		// return err
	}
	return nil
}

func (r *KafkaConsumer) deviceDeleteServe(ctx context.Context, msg kafka.Message) error {
	if err := r.deviceService.Delete(ctx, string(msg.Value)); err != nil {
		logger.Logger.Error("Failed to delete Peson from db: ", zap.Error(err))
		return err
	}
	return nil
}
