package consumer

import (
	"context"
	"device-manager/internal/domain"
	"device-manager/internal/logger"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

func (r *KafkaConsumer) deviceCreateServe(ctx context.Context, msg kafka.Message) error {
	// req := proto.PersonCreateRequest{}
	// if err := pm.Unmarshal(msg.Value, &req); err != nil {
	// 	logger.Logger.Error("Failed to unmarshal message: ", zap.Error(err))
	// 	return err
	// }
	req := domain.Device{}
	kafka.Unmarshal(msg.Value, &req)

	if err := r.deviceService.CreateDevice(ctx, req.UUID, req.Platform, req.Language, req.Geolocation, req.Email); err != nil {
		logger.Logger.Error("Failed to add Peson to db: ", zap.Error(err))
		return err
	}
	return nil
}

func (r *KafkaConsumer) deviceUpdateServe(ctx context.Context, msg kafka.Message) error {
	// req := proto.PersonUpdateRequest{}
	// if err := pm.Unmarshal(msg.Value, &req); err != nil {
	// 	logger.Logger.Error("Failed to unmarshal message: ", zap.Error(err))
	// 	return err
	// }
	// if err := r.deviceService.PersonService.UpdatePerson(ctx, req.Id, req.Name); err != nil {
	// 	logger.Logger.Error("Failed to update Person in db: ", zap.Error(err))
	// 	return err
	// }
	 return nil
}

func (r *KafkaConsumer) deviceDeleteServe(ctx context.Context, msg kafka.Message) error {
	// req := proto.Id{}
	// if err := pm.Unmarshal(msg.Value, &req); err != nil {
	// 	logger.Logger.Error("Failed to unmarshal message: ", zap.Error(err))
	// 	return err
	// }
	// if err := r.deviceService.PersonService.DeletePersonById(ctx, req.Id); err != nil {
	// 	logger.Logger.Error("Failed to delete Person from db: ", zap.Error(err))
	// 	return err
	// }
	 return nil
}
