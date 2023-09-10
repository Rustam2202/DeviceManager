package event

import "device-manager/internal/service"

type EventHandler struct {
	service *service.DeviceService
}

func NewEventHandler(s *service.DeviceService) *EventHandler {
	return &EventHandler{service: s}
}
