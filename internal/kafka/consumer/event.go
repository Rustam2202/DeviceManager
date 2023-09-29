package consumer

import (
	"context"
	"device-manager/internal/logger"
	"encoding/json"
	"time"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

type addEventRequest struct {
	UUID       string        `json:"uuid"`
	Name       string        `json:"name"`
	Attributes []interface{} `json:"attributes"`
	CreatedAt  time.Time     `json:"created_at"`
}

func (r *KafkaConsumer) eventCreate(ctx context.Context, msg kafka.Message) error {
	var req addEventRequest
	err := json.Unmarshal(msg.Value, &req)
	if err != nil {
		logger.Logger.Error("Failed to unmarshal kafka message.", zap.Error(err))
		return err
	}
	err = r.eventService.Create(ctx, req.UUID, req.Name, req.Attributes, req.CreatedAt)
	if err != nil {
		logger.Logger.Error("Failed to add Event to db.", zap.Error(err))
		return err
	}
	return nil
}
