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
	Get(context.Context, string) (*domain.Device, error)
	GetByLanguage(context.Context, string) ([]domain.Device, error)
	GetByGeolocation(context.Context, float64, float64, int) ([]domain.Device, error)
	GetByEmail(context.Context, string) ([]domain.Device, error)
	UpdateLanguage(context.Context, string, string) error
	UpdateGeolocation(context.Context, string, []float64) error
	UpdateEmail(context.Context, string, string) error
	Delete(context.Context, string) error
}

type EventRepository interface {
	Create(context.Context, *domain.Event) error
	Get(context.Context, string, time.Time, time.Time) ([]domain.Event, error)
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

func (s *DeviceService) Get(ctx context.Context, id string) (*domain.Device, error) {
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

func (s *DeviceService) UpdateLanguage(ctx context.Context, id string, lang string) error {
	if err := s.repoDevice.UpdateLanguage(ctx, id, lang); err != nil {
		return err
	}
	logger.Logger.Info("language updated in db", zap.String("uuid: ", id), zap.String("language", lang))
	return nil
}

func (s *DeviceService) UpdateGeolocation(ctx context.Context, id string, coordinates []float64) error {
	if err := s.repoDevice.UpdateGeolocation(ctx, id, coordinates); err != nil {
		return err
	}
	logger.Logger.Info("geoposition updated in db", zap.String("uuid: ", id))
	return nil
}

func (s *DeviceService) UpdateEmail(ctx context.Context, id, email string) error {
	if err := s.repoDevice.UpdateEmail(ctx, id, email); err != nil {
		return err
	}
	logger.Logger.Info("E-mail updated in db", zap.String("uuid: ", id), zap.String("e-mail:", email))
	return nil
}

func (s *DeviceService) Delete(ctx context.Context, id string) error {
	if err := s.repoDevice.Delete(ctx, id); err != nil {
		return err
	}
	logger.Logger.Info("device deleted from db", zap.String("uuid: ", id))
	return nil
}

func (s *EventService) Create(ctx context.Context,
	id string, name string, attributes []interface{}, createdAt time.Time) error {
	device, err := s.repoDevice.Get(ctx, id)
	if err != nil {
		return err
	}
	if device == nil {
		return fmt.Errorf("no device exist with '%s' uuid", id)
	}
	err = s.repoEvent.Create(ctx, &domain.Event{
		DeviceUUID: id,
		Name:       name,
		CreatedAt:  createdAt,
		Attributes: attributes,
	})
	if err != nil {
		return err
	}
	logger.Logger.Info("event added to db", zap.String("uuid", id), zap.String("name", name))
	return nil
}

func (s *EventService) Get(ctx context.Context,
	id string, begin, end time.Time, filter string) ([]domain.Event, error) {
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
