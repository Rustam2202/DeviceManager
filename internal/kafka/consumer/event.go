package consumer

import (
	"context"
	"device-manager/internal/logger"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

type addEventRequest struct {
	UUID       uuid.UUID
	Name       string
	Attributes []interface{}
}

func (r *KafkaConsumer) eventCreate(ctx context.Context, msg kafka.Message) error {
	var req addEventRequest
	err := json.Unmarshal(msg.Value, &req)
	if err != nil {
		logger.Logger.Error("Failed to unmarshal kafka message.", zap.Error(err))
		return err
	}
	err = r.eventService.Create(ctx, req.UUID, req.Name, req.Attributes)
	if err != nil {
		logger.Logger.Error("Failed to add Event to db.", zap.Error(err))
		return err
	}
	return nil
}
