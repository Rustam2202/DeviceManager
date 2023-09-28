package device

import (
	"device-manager/internal/kafka"
	"device-manager/internal/server/handlers"
	"device-manager/internal/server/handlers/utils"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateLanguageRequest struct {
	UUID     string `json:"uuid"`
	Language string `json:"language"`
}

// @Summary	Update a device language
// @Tags		Device
// @Accept		json
// @Produce	json
// @Param		request	body		UpdateLanguageRequest	true	"Update Device language Request"
// @Success	200		{object}	UpdateLanguageRequest
// @Failure	400		{object}	handlers.ErrorResponce
// @Failure	500		{object}	handlers.ErrorResponce
// @Router		/device_lang [put]
func (h *DeviceHandler) UpdateLanguage(ctx *gin.Context) {
	var req UpdateLanguageRequest
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

	bytes, err := json.Marshal(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			handlers.ErrorResponce{Message: "Failed to marshal request for kafka", Error: err})
		return
	}
	err = h.Producer.WriteMessage(ctx, kafka.DeviceUpdateLanguage, bytes)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			handlers.ErrorResponce{Message: "Failed to add a new person to database", Error: err})
		return
	}
	ctx.JSON(http.StatusOK, nil)
}

type UpdateGeolocationRequest struct {
	UUID      string  `json:"uuid"`
	Longitude float64 `json:"longitude" default:"55.646575"`
	Latitude  float64 `json:"latitude" default:"37.552375"`
}

// @Summary	Update a device geolocation
// @Tags		Device
// @Accept		json
// @Produce	json
// @Param		request	body		UpdateGeolocationRequest	true	"Update Device geolocation Request"
// @Success	200		{object}	UpdateGeolocationRequest
// @Failure	400		{object}	handlers.ErrorResponce
// @Failure	500		{object}	handlers.ErrorResponce
// @Router		/device_geo [put]
func (h *DeviceHandler) UpdateGeolocation(ctx *gin.Context) {
	var req UpdateGeolocationRequest
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
	bytes, err := json.Marshal(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			handlers.ErrorResponce{Message: "Failed to marshal request for kafka", Error: err})
		return
	}
	err = h.Producer.WriteMessage(ctx, kafka.DeviceUpdateLanguage, bytes)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			handlers.ErrorResponce{Message: "Failed to add a new person to database", Error: err})
		return
	}
	ctx.JSON(http.StatusOK, nil)
}

type UpdateEmailRequest struct {
	UUID  string `json:"uuid"`
	Email string `json:"email"`
}

// @Summary	Update a device E-mail
// @Tags		Device
// @Accept		json
// @Produce	json
// @Param		request	body		UpdateEmailRequest	true	"Update Device E-mail Request"
// @Success	200		{object}	UpdateEmailRequest
// @Failure	400		{object}	handlers.ErrorResponce
// @Failure	500		{object}	handlers.ErrorResponce
// @Router		/device_email [put]
func (h *DeviceHandler) UpdateEmail(ctx *gin.Context) {
	var req UpdateEmailRequest
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
	err = h.Producer.WriteMessage(ctx, kafka.DeviceUpdateLanguage, bytes)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			handlers.ErrorResponce{Message: "Failed to add a new person to database", Error: err})
		return
	}
	ctx.JSON(http.StatusOK, nil)
}
