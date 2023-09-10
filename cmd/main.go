package main

import (
	"device-manager/internal/config"
	"device-manager/internal/database"
	"device-manager/internal/logger"
	"device-manager/internal/repository"
	"device-manager/internal/server"
	"device-manager/internal/server/handlers/device"
	"device-manager/internal/service"

	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	cfg := config.LoadConfig()
	logger.IntializeLogger(*cfg.LoggerConfig)
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	wg := &sync.WaitGroup{}

	mdb := database.NewMongo(cfg.DatabaseConfig)
	devicerRepository := repository.NewDeviceRepository(mdb)
	eventRepository := repository.NewEventRepository(mdb)
	deviceService := service.NewDeviceService(devicerRepository, eventRepository)
	deviceHandler := device.NewDeviceHandler(deviceService)
	s := server.NewHTTPServer(cfg.ServerHTTPConfig, deviceHandler)
	s.StartHTTPServer(ctx, wg)
}
