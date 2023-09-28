package main

import (
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
	"fmt"

	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/google/uuid"
)

func main() {
	u, _ := uuid.NewRandom()
	fmt.Println(u)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	cfg := config.LoadConfig()
	logger.MustLogger(*cfg.LoggerConfig)

	mdb := database.MustConnectMongo(ctx, cfg.DatabaseConfig)

	wg := &sync.WaitGroup{}

	kafkaProducer := producer.NewKafkaProducer(cfg.KafkaConfig)

	deviceRepository := repository.NewDeviceRepository(mdb)
	eventRepository := repository.NewEventRepository(mdb)
	deviceService := service.NewDeviceService(deviceRepository)
	eventService := service.NewEventService(deviceRepository, eventRepository)
	deviceHandler := device.NewDeviceHandler(deviceService, kafkaProducer)
	eventHandler := event.NewEventHandler(eventService, kafkaProducer)

	kafkaConsumer := consumer.NewKafkaConsumer(cfg.KafkaConfig, deviceService, eventService)
	kafkaConsumer.RunKafkaConsumer(ctx, wg)

	s := server.NewHTTP(cfg.ServerHTTPConfig, deviceHandler, eventHandler)
	wg.Add(1)
	go s.StartHTTP(ctx, wg)

	wg.Wait()
}
