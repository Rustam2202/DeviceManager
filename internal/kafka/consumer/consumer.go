package consumer

import (
	"context"
	k "device-manager/internal/kafka"
	"device-manager/internal/service"
	"sync"

	"github.com/segmentio/kafka-go"
)

type KafkaConsumer struct {
	deviceService *service.DeviceService
	eventService  *service.EventService
	cfg           *kafka.ReaderConfig
}

func NewKafkaConsumer(cfg *k.KafkaConfig, ds *service.DeviceService, es *service.EventService) *KafkaConsumer {
	config := kafka.ReaderConfig{}
	config.Brokers = cfg.Brokers
	config.GroupID = cfg.Group
	return &KafkaConsumer{cfg: &config, deviceService: ds, eventService: es}
}

func (r *KafkaConsumer) RunKafkaConsumer(ctx context.Context, wg *sync.WaitGroup) {
	TopicsServe := map[string]func(context.Context, kafka.Message) error{
		k.DeviceCreate: r.deviceCreateServe,
		k.DeviceUpdate: r.deviceUpdateServe,
		k.DeviceDelete: r.deviceDeleteServe,
		k.EventCreate:  r.eventCreateServe,
	}
	for k, v := range TopicsServe {
		wg.Add(1)
		go r.RunReader(ctx, wg, k, v)
	}
}
