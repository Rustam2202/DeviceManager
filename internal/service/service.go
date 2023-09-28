package service

import (
	"context"
	"device-manager/internal/domain"
	"device-manager/internal/logger"
	"fmt"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type DeviceRepository interface {
	Create(context.Context, *domain.Device) error
	Get(context.Context, uuid.UUID) (*domain.Device, error)
	GetByLanguage(context.Context, string) ([]domain.Device, error)
	GetByGeolocation(context.Context, float64, float64, int) ([]domain.Device, error)
	GetByEmail(context.Context, string) ([]domain.Device, error)
	UpdateLanguage(context.Context, uuid.UUID, string) error
	UpdateGeolocation(context.Context, uuid.UUID, []float64) error
	UpdateEmail(context.Context, uuid.UUID, string) error
	Delete(context.Context, uuid.UUID) error
}

type EventRepository interface {
	Create(context.Context, *domain.Event) error
	Get(context.Context, uuid.UUID, time.Time, time.Time) ([]domain.Event, error)
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

func (s *DeviceService) Create(ctx context.Context,
	id uuid.UUID, platform string, lang string, email string, coordinates []float64) error {
	err := s.repoDevice.Create(ctx, &domain.Device{
		UUID:     id.String(),
		Platform: platform,
		Language: lang,
		Location: domain.Location{
			Type:        "Point",
			Coordinates: coordinates,
		},
		Email: email,
	})
	if err != nil {
		return err
	}
	logger.Logger.Info("device added to db", zap.String("uuid: ", id.String()))
	return nil
}

func (s *DeviceService) Get(ctx context.Context, id uuid.UUID) (*domain.Device, error) {
	device, err := s.repoDevice.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return device, nil
}

func (s *DeviceService) GetByLanguage(ctx context.Context, lang string) ([]domain.Device, error) {
	device, err := s.repoDevice.GetByLanguage(ctx, lang)
	if err != nil {
		return nil, err
	}
	return device, nil
}

func (s *DeviceService) GetByGeolocation(ctx context.Context,
	long, lat float64, radius int) ([]domain.Device, error) {
	device, err := s.repoDevice.GetByGeolocation(ctx, long, lat, radius)
	if err != nil {
		return nil, err
	}
	return device, nil
}

func (s *DeviceService) GetByEmail(ctx context.Context, email string) ([]domain.Device, error) {
	device, err := s.repoDevice.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return device, nil
}

func (s *DeviceService) UpdateLanguage(ctx context.Context, id uuid.UUID, lang string) error {
	device, err := s.repoDevice.Get(ctx, id)
	if err != nil {
		return err
	}
	if device == nil {
		return fmt.Errorf("no device exist with '%s' uuid", id)
	}
	err = s.repoDevice.UpdateLanguage(ctx, id, lang)
	if err != nil {
		return err
	}
	logger.Logger.Info("language updated in db", zap.String("uuid: ", id.String()), zap.String("language", lang))
	return nil
}

func (s *DeviceService) UpdateGeolocation(ctx context.Context, id uuid.UUID, coordinates []float64) error {
	device, err := s.repoDevice.Get(ctx, id)
	if err != nil {
		return err
	}
	if device == nil {
		return fmt.Errorf("no device exist with '%s' uuid", id)
	}
	err = s.repoDevice.UpdateGeolocation(ctx, id, coordinates)
	if err != nil {
		return err
	}
	logger.Logger.Info("geoposition updated in db", zap.String("uuid: ", id.String()))
	return nil
}

func (s *DeviceService) UpdateEmail(ctx context.Context, id uuid.UUID, email string) error {
	device, err := s.repoDevice.Get(ctx, id)
	if err != nil {
		return err
	}
	if device == nil {
		return fmt.Errorf("no device exist with '%s' uuid", id)
	}
	device.Email = email
	err = s.repoDevice.UpdateEmail(ctx, id, email)
	if err != nil {
		return err
	}
	logger.Logger.Info("E-mail updated in db", zap.String("uuid: ", id.String()), zap.String("e-mail:", email))
	return nil
}

func (s *DeviceService) Delete(ctx context.Context, id uuid.UUID) error {
	err := s.repoDevice.Delete(ctx, id)
	if err != nil {
		return err
	}
	logger.Logger.Info("device deleted from db", zap.String("uuid: ", id.String()))
	return nil
}

func (s *EventService) Create(ctx context.Context,
	id uuid.UUID, name string, attributes []interface{}) error {
	device, err := s.repoDevice.Get(ctx, id)
	if err != nil {
		return err
	}
	if device == nil {
		return fmt.Errorf("no device exist with '%s' uuid", id)
	}
	event := domain.Event{
		DeviceUUID: id,
		Name:       name,
		CreatedAt:  time.Now(),
		Attributes: attributes,
	}
	err = s.repoEvent.Create(ctx, &event)
	if err != nil {
		return nil
	}
	logger.Logger.Info("event added to db", zap.String("uuid", id.String()), zap.String("name", event.Name))
	return nil
}

func (s *EventService) Get(ctx context.Context,
	id uuid.UUID, begin, end time.Time, filter string) ([]domain.Event, error) {
	device, err := s.repoDevice.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	if device == nil {
		return nil, fmt.Errorf("no device exist with '%s' uuid", id)
	}
	events, err := s.repoEvent.Get(ctx, id, begin, end)
	if err != nil {
		return nil, err
	}
	return events, nil
}
