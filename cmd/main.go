package main

import (
	"context"
	"device-manager/internal/config"
	"device-manager/internal/database"
	"device-manager/internal/kafka/consumer"
	"device-manager/internal/kafka/producer"
	"device-manager/internal/logger"
	"device-manager/internal/repository"
	"device-manager/internal/server"
	"device-manager/internal/server/handlers/device"
	"device-manager/internal/server/handlers/event"
	"device-manager/internal/service"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	cfg := config.MustLoadConfig()
	logger.MustConfigLogger(*cfg.LoggerConfig)
	mdb := database.MustConnectMongo(ctx, cfg.Database)

	wg := &sync.WaitGroup{}

	kafkaProducer := producer.NewKafkaProducer(cfg.Kafka)

	deviceRepository := repository.NewDeviceRepository(mdb)
	eventRepository := repository.NewEventRepository(mdb)
	deviceService := service.NewDeviceService(deviceRepository)
	eventService := service.NewEventService(deviceRepository, eventRepository)
	deviceHandler := device.NewDeviceHandler(deviceService, kafkaProducer)
	eventHandler := event.NewEventHandler(eventService, kafkaProducer)

	kafkaConsumer := consumer.NewKafkaConsumer(cfg.Kafka, deviceService, eventService)
	kafkaConsumer.RunKafkaConsumer(ctx, wg)

	s := server.NewHTTP(cfg.Server, deviceHandler, eventHandler)
	wg.Add(1)
	go s.StartHTTP(ctx, wg)

	wg.Wait()
}
