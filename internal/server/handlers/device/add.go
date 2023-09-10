package device

import (
	"device-manager/internal/domain"
	"device-manager/internal/server/handlers"

	"net/http"

	"github.com/gin-gonic/gin"
)

type AddDeviceRequest struct {
	UUID        string `json:"uuid"`
	Platform    string `json:"platform"`
	Language    string `json:"language"`
	Geolocation string `json:"geolocation"`
	Email       string `json:"email"`
}

// @Summary		Add device
// @Description	Add a new device to database
// @Tags			Device
// @Accept			json
// @Produce		json
// @Param			request	body		AddDeviceRequest	true	"Add Device Request"
// @Success		201 
// @Failure		400		{object}	handlers.ErrorResponce
// @Failure		500		{object}	handlers.ErrorResponce
// @Router			/device [post]
func (h *DeviceHandler) AddDevice(ctx *gin.Context) {
	var req AddDeviceRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			handlers.ErrorResponce{Message: "Failed to parse request", Error: err})
		return
	}
	err = h.service.CreateDevice(ctx, req.UUID, req.Platform, req.Language, req.Geolocation, req.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			handlers.ErrorResponce{Message: "Failed to add a new person to database", Error: err})
		return
	}
	ctx.JSON(http.StatusCreated, domain.Device{
		UUID:        req.UUID,
		Platform:    req.Platform,
		Language:    req.Language,
		Geolocation: req.Geolocation,
		Email:       req.Email,
	})
}

