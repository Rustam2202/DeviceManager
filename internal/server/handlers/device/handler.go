package device

import (
	"device-manager/internal/kafka/producer"
	"device-manager/internal/service"
)

type DeviceHandler struct {
	service  *service.DeviceService
	Producer *producer.KafkaProducer
}

func NewDeviceHandler(s *service.DeviceService, p *producer.KafkaProducer) *DeviceHandler {
	return &DeviceHandler{service: s, Producer: p}
}
