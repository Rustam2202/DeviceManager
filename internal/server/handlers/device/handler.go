package device

import "device-manager/internal/service"

type DeviceHandler struct {
	service *service.DeviceService
}

func NewDeviceHandler(s *service.DeviceService) *DeviceHandler {
	return &DeviceHandler{service: s}
}
