package service

import (
	"context"
	"device-manager/internal/domain"
	"device-manager/internal/logger"
	"fmt"
	"time"

	"github.com/google/uuid"
	"golang.org/x/text/language"
)

type DeviceRepository interface {
	Create(context.Context, *domain.Device) error
	Get(context.Context, uuid.UUID) (*domain.Device, error)
	GetByLanguage(context.Context, language.Tag) ([]domain.Device, error)
	GetByGeolocation(context.Context, float64, float64, float64) ([]domain.Device, error)
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
	uuid uuid.UUID, platform string, lang language.Tag, email string, coordinates []float64) error {
	// if u.String() == "" { //
	// 	return fmt.Errorf("uuid mustn't be empty")
	// }
	dev, err := s.Get(ctx, uuid)
	if err != nil {
		return err
	}
	if dev != nil {
		return fmt.Errorf("devise %s already exist", dev.UUID)
	}
	err = s.repoDevice.Create(ctx, &domain.Device{
		UUID:     uuid,
		Platform: platform,
		Language: lang.String(),
		Location: domain.Location{
			Type:        "Point",
			Coordinates: coordinates,
		},
		Email: email,
	})
	if err != nil {
		return err
	}
	// logger.Logger.Info(fmt.Sprintf("device %s added to db with id:%s", device.UUID, device.ID.String()))
	return nil
}

func (s *DeviceService) Get(ctx context.Context, uuid uuid.UUID) (*domain.Device, error) {
	device, err := s.repoDevice.Get(ctx, uuid)
	if err != nil {
		return nil, err
	}
	return device, nil
}

func (s *DeviceService) GetByLanguage(ctx context.Context, lang language.Tag) ([]domain.Device, error) {
	device, err := s.repoDevice.GetByLanguage(ctx, lang)
	if err != nil {
		return nil, err
	}
	return device, nil
}

func (s *DeviceService) GetByGeolocation(ctx context.Context,
	long, lat, radius float64) ([]domain.Device, error) {
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

func (s *DeviceService) UpdateLanguage(ctx context.Context, uuid uuid.UUID, lang string) error {
	device, err := s.repoDevice.Get(ctx, uuid)
	if err != nil {
		return err
	}
	if device == nil {
		return fmt.Errorf("no device exist with '%s' uuid", uuid)
	}
	err = s.repoDevice.UpdateLanguage(ctx, uuid, lang)
	if err != nil {
		return err
	}
	logger.Logger.Info(fmt.Sprintf("device %s: language was updated on %s", device.UUID, device.Language))
	return nil
}

func (s *DeviceService) UpdateGeolocation(ctx context.Context, uuid uuid.UUID, coordinates []float64) error {
	device, err := s.repoDevice.Get(ctx, uuid)
	if err != nil {
		return err
	}
	if device == nil {
		return fmt.Errorf("no device exist with '%s' uuid", uuid)
	}
	err = s.repoDevice.UpdateGeolocation(ctx, uuid, coordinates)
	if err != nil {
		return err
	}
	logger.Logger.Info(fmt.Sprintf("device %s: geoposition was updated on %s", device.UUID, device.Location))
	return nil
}

func (s *DeviceService) UpdateEmail(ctx context.Context, uuid uuid.UUID, email string) error {
	device, err := s.repoDevice.Get(ctx, uuid)
	if err != nil {
		return err
	}
	if device == nil {
		return fmt.Errorf("no device exist with '%s' uuid", uuid)
	}
	device.Email = email
	err = s.repoDevice.UpdateEmail(ctx, uuid, email)
	if err != nil {
		return err
	}
	logger.Logger.Info(fmt.Sprintf("device %s: e-mail was updated on %s", device.UUID, device.Email))
	return nil
}

func (s *DeviceService) Delete(ctx context.Context, uuid uuid.UUID) error {
	err := s.repoDevice.Delete(ctx, uuid)
	if err != nil {
		return err
	}
	logger.Logger.Info(fmt.Sprintf("device %s was deleted", uuid))
	return nil
}

func (s *EventService) Create(ctx context.Context,
	uuid uuid.UUID, name string, attributes []interface{}) error {
	device, err := s.repoDevice.Get(ctx, uuid)
	if err != nil {
		return err
	}
	if device == nil {
		return fmt.Errorf("no device exist with '%s' uuid", uuid)
	}
	event := domain.Event{
		DeviceUUID: uuid,
		Name:       name,
		CreatedAt:  time.Now(),
		Attributes: attributes,
	}
	err = s.repoEvent.Create(ctx, &event)
	if err != nil {
		return nil
	}
	logger.Logger.Info(fmt.Sprintf("event '%s' added to db with id:%s", event.Name, event.ID.String()))
	return nil
}

func (s *EventService) Get(ctx context.Context, uuid uuid.UUID, begin, end time.Time, filter string) ([]domain.Event, error) {
	device, err := s.repoDevice.Get(ctx, uuid)
	if err != nil {
		return nil, err
	}
	if device == nil {
		return nil, fmt.Errorf("no device exist with '%s' uuid", uuid)
	}
	events, err := s.repoEvent.Get(ctx, uuid, begin, end)
	if err != nil {
		return nil, err
	}
	return events, nil
}
