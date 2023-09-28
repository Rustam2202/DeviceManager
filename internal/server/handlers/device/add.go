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
	Longitude float64 `json:"longitude" default:"55.646575"`
	Latitude  float64 `json:"latitude" default:"37.552375"`
}

type DeviceRequest struct {
	UUID        string      `json:"uuid" default:"f2366ac9-0663-4d0b-964f-98c388240d5c"`
	Platform    string      `json:"platform" default:"macOs"`
	Language    string      `json:"language" default:"en-US"`
	Coordinates Coordinates `json:"coordinates"`
	Email       string      `json:"email" default:"some@email.com"`
}

//	@Summary		Add device
//	@Description	Add a new device to database
//	@Tags			Device
//	@Accept			json
//	@Produce		json
//	@Param			request	body	DeviceRequest	true	"Add Device Request"
//	@Success		200
//	@Failure		400	{object}	handlers.ErrorResponce
//	@Failure		500	{object}	handlers.ErrorResponce
//	@Router			/device [post]
func (h *DeviceHandler) Add(ctx *gin.Context) {
	var req DeviceRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			handlers.ErrorResponce{Message: "Failed to parse request", Error: err})
		return
	}

	if _, errResp := utils.UuidValidationAndParse(req.UUID); errResp != nil {
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
