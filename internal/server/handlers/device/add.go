package device

import (
	"device-manager/internal/kafka"
	"device-manager/internal/server/handlers"
	"encoding/json"
	"net/http"
	"net/mail"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/text/language"
)

type AddDeviceRequest struct {
	UUID        string    `json:"uuid"`
	Platform    string    `json:"platform"`
	Language    string    `json:"language"`
	Coordinates []float64 `json:"coordinates"`
	Email       string    `json:"email"`
}

// @Summary		Add device
// @Description	Add a new device to database
// @Tags			Device
// @Accept			json
// @Produce		json
// @Param			request	body	AddDeviceRequest	true	"Add Device Request"
// @Success		200
// @Failure		400	{object}	handlers.ErrorResponce
// @Failure		500	{object}	handlers.ErrorResponce
// @Router			/device [post]
func (h *DeviceHandler) Add(ctx *gin.Context) {
	var req AddDeviceRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			handlers.ErrorResponce{Message: "Failed to parse request", Error: err})
		return
	}

	_, err = uuid.ParseBytes([]byte(req.UUID))
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			handlers.ErrorResponce{Message: "Failed to parse UUID", Error: err})
		return
	}
	_, err = language.Parse(req.Language)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			handlers.ErrorResponce{Message: "Failed to parse language", Error: err})
		return
	}
	_, err = mail.ParseAddress(req.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			handlers.ErrorResponce{Message: "Failed to parse E-mail", Error: err})
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
