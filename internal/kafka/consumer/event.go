package consumer

import (
	"context"
	"device-manager/internal/logger"
	"encoding/json"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

type AddEventRequest struct {
	UUID       string
	Name       string
	Attributes []interface{}
}

func (r *KafkaConsumer) eventCreateServe(ctx context.Context, msg kafka.Message) error {
	var req AddEventRequest
	err := json.Unmarshal(msg.Value, &req)
	if err != nil {
		logger.Logger.Error("Failed to unmarshal kafka message.", zap.Error(err))
		return err
	}
	var validAttributes []interface{}
	for _, attr := range req.Attributes {
		switch v := attr.(type) {
		case string:
			validAttributes = append(validAttributes, v)
		case int:
			validAttributes = append(validAttributes, v)
		case float64:
			validAttributes = append(validAttributes, v)
		case bool:
			validAttributes = append(validAttributes, v)
		default:
			continue
		}
	}

	err = r.eventService.CreateEvent(ctx, req.UUID, req.Name, validAttributes)
	if err != nil {
		logger.Logger.Error("Failed to add Event to db.", zap.Error(err))
		return err
	}
	return nil
}
