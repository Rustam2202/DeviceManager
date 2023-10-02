package event

import (
	"device-manager/internal/kafka/producer"
	"device-manager/internal/service"
)

type EventHandler struct {
	Service  *service.EventService
	Producer *producer.KafkaProducer
}

func NewEventHandler(s *service.EventService, p *producer.KafkaProducer) *EventHandler {
	return &EventHandler{Service: s, Producer: p}
}
