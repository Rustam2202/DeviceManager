package device

import (
	"device-manager/internal/kafka"
	"device-manager/internal/server/handlers"
	"device-manager/internal/server/handlers/utils"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Coordinates struct {
	Longitude float64 `json:"longitude" example:"55.646574"`
	Latitude  float64 `json:"latitude" example:"37.552375"`
}

type DeviceRequest struct {
	UUID        string      `json:"uuid" example:"550e8400-e29b-41d4-a716-446655440000"`
	Platform    string      `json:"platform" example:"Mozilla/5.0 (Linux; U; Android 2.3.7; en-us; Nexus One Build/FRF91) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1"`
	Language    string      `json:"language" example:"en-US"`
	Coordinates Coordinates `json:"coordinates"`
	Email       string      `json:"email" example:"example@email.com"`
}

// @Summary		Add device
// @Description	Add a new device to database
// @Tags			Device
// @Accept			json
// @Produce		json
// @Param			request	body	DeviceRequest	true	"Add Device Request"
// @Success		200
// @Failure		400	{object}	handlers.ErrorResponce
// @Failure		500	{object}	handlers.ErrorResponce
// @Router			/device [post]
func (h *DeviceHandler) Add(ctx *gin.Context) {
	var req DeviceRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			handlers.ErrorResponce{Message: "Failed to parse request", Error: err})
		return
	}

	if errResp := utils.UuidValidation(req.UUID); errResp != nil {
		ctx.JSON(http.StatusBadRequest, errResp)
		return
	}
	if errResp := utils.LanguageValidation(req.Language); errResp != nil {
		ctx.JSON(http.StatusBadRequest, errResp)
		return
	}
	if errResp := utils.EmailValidation(req.Email); errResp != nil {
		ctx.JSON(http.StatusBadRequest, errResp)
		return
	}
	if errResp := utils.PlatformValidation(req.Platform); errResp != nil {
		ctx.JSON(http.StatusBadRequest, errResp)
		return
	}

	bytes, err := json.Marshal(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			handlers.ErrorResponce{Message: "Failed to marshal request for kafka", Error: err})
		return
	}
	err = h.Producer.WriteMessage(ctx, kafka.DeviceCreate, bytes)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			handlers.ErrorResponce{Message: "Failed to write message to kafka", Error: err})
		return
	}
	ctx.JSON(http.StatusOK, nil)
}
