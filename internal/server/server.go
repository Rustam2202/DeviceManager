package server

import (
	"device-manager/docs"
	"device-manager/internal/logger"
	"device-manager/internal/server/handlers/device"
	"device-manager/internal/server/handlers/event"
	"fmt"

	"context"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

type HTTPServer struct {
	cfg           *ServerHTTPConfig
	HTTPServer    *http.Server
	deviceHandler *device.DeviceHandler
	eventHandler  *event.EventHandler
	
}

func NewHTTPServer(cfg *ServerHTTPConfig, dh *device.DeviceHandler, eh *event.EventHandler) *HTTPServer {
	return &HTTPServer{cfg: cfg, deviceHandler: dh, eventHandler: eh}
}

//	@title		Device Manager API
//	@version	1.0
//	@description
//	@BasePath
func (s *HTTPServer) StartHTTPServer(ctx context.Context, wg *sync.WaitGroup) {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%d", s.cfg.Host, s.cfg.Port)

	{
		r.GET("/", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "Hello world from Party Calc http server")
		})
		r.POST("/device", s.deviceHandler.AddDevice)
		r.GET("/device/:uuid", s.deviceHandler.GetDevice)
		r.PUT("/device_lang", s.deviceHandler.UpdateLanguage)
		r.PUT("/device_geo", s.deviceHandler.UpdateGeolocation)
		r.PUT("/device_email", s.deviceHandler.UpdateEmail)
		r.DELETE("/device/:uuid", s.deviceHandler.Delete)
		
		r.POST("/event", s.eventHandler.AddEventRequest)
		r.GET("/event", s.eventHandler.GetEvents)

		r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	s.HTTPServer = &http.Server{
		Addr:    fmt.Sprintf("%s:%d", s.cfg.Host, s.cfg.Port),
		Handler: r,
	}

	go func() {
		defer wg.Done()
		logger.Logger.Info("Starting HTTP server ...")
		err := s.HTTPServer.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			logger.Logger.Error("Failed to start HTTP server", zap.Error(err))
		}
	}()
	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	logger.Logger.Info("Shutting down HTTP server ...")
	if err := s.HTTPServer.Shutdown(shutdownCtx); err != nil {
		logger.Logger.Error("Failed to shutdown HTTP server", zap.Error(err))
	}
}
