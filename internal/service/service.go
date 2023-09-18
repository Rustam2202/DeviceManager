package service

import (
	"context"
	"device-manager/internal/domain"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DeviceRepository interface {
	Create(context.Context, *domain.Device) (*domain.Device, error)
	Get(context.Context, string) (*domain.Device, error)
	Update(context.Context, *domain.Device) error
	Delete(context.Context, string) error
}

type EventRepository interface {
	Create(context.Context, *domain.Event) error
	Get(context.Context, primitive.ObjectID, time.Time, time.Time) ([]domain.Event, error)
}

type DeviceService struct {
	repoDevice DeviceRepository
}

type EventService struct {
	repoDevice DeviceRepository
	repoEvent  EventRepository
}

func NewDeviceService(rd DeviceRepository) *DeviceService {
	return &DeviceService{repoDevice: rd}
}

func NewEventService(rd DeviceRepository, re EventRepository) *EventService {
	return &EventService{repoDevice: rd, repoEvent: re}
}

func (s *DeviceService) CreateDevice(ctx context.Context, uuid, platform, lang, geo, email string) error {
	_, err := s.repoDevice.Create(ctx, &domain.Device{
		UUID:        uuid,
		Platform:    platform,
		Language:    lang,
		Geolocation: geo,
		Email:       email,
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *EventService) CreateEvent(ctx context.Context, uuid, name string, attributes []interface{}) error {
	device, err := s.repoDevice.Get(ctx, uuid)
	if err != nil {
		return fmt.Errorf("no device exist with '%s' uuid", uuid)
	}
	event := domain.Event{
		DeviceId:   device.ID,
		Name:       name,
		CreatedAt:  time.Now(),
		Attributes: attributes,
	}
	err = s.repoEvent.Create(ctx, &event)
	if err != nil {
		return nil
	}
	return nil
}

func (s *DeviceService) GetDeviceInfo(ctx context.Context, uuid string) (*domain.Device, error) {
	device, err := s.repoDevice.Get(ctx, uuid)
	if err != nil {
		return nil, err
	}
	return device, nil
}

func (s *EventService) GetDeviceEvents(ctx context.Context, uuid string, begin, end time.Time, filter string) ([]domain.Event, error) {
	device, err := s.repoDevice.Get(ctx, uuid)
	if err != nil {
		return nil, err
	}
	events, err := s.repoEvent.Get(ctx, device.ID, begin, end)
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (s *DeviceService) UpdateLaguage(ctx context.Context, uuid, lang string) error {
	device, err := s.repoDevice.Get(ctx, uuid)
	if err != nil {
		return err
	}
	device.Language = lang
	err = s.repoDevice.Update(ctx, device)
	if err != nil {
		return err
	}
	return nil
}

func (s *DeviceService) UpdateGeolocation(ctx context.Context, uuid, geo string) error {
	device, err := s.repoDevice.Get(ctx, uuid)
	if err != nil {
		return err
	}
	device.Geolocation = geo
	err = s.repoDevice.Update(ctx, device)
	if err != nil {
		return err
	}
	return nil
}

func (s *DeviceService) UpdateEmail(ctx context.Context, uuid, email string) error {
	device, err := s.repoDevice.Get(ctx, uuid)
	if err != nil {
		return err
	}
	device.Email = email
	err = s.repoDevice.Update(ctx, device)
	if err != nil {
		return err
	}
	return nil
}

func (s *DeviceService) Delete(ctx context.Context, uuid string) error {
	err := s.repoDevice.Delete(ctx, uuid)
	if err != nil {
		return err
	}
	return nil
}
