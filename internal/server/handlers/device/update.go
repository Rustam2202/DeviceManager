package device

import (
	"device-manager/internal/kafka"
	"device-manager/internal/server/handlers/utils"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LanguageRequest struct {
	UUID     string `json:"uuid" example:"550e8400-e29b-41d4-a716-446655440000"`
	Language string `json:"language" example:"ru"`
}

//	@Summary	Update a device language
//	@Tags		Device UPDATE
//	@Accept		json
//	@Produce	json
//	@Param		request	body	LanguageRequest	true	"Update Device language Request"
//	@Success	200
//	@Failure	400	{object}	utils.ErrorResponce
//	@Failure	500	{object}	utils.ErrorResponce
//	@Router		/device_lang [put]
func (h *DeviceHandler) UpdateLanguage(ctx *gin.Context) {
	var req LanguageRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			utils.ErrorResponce{Message: "Failed to parse request", Error: err})
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

	bytes, err := json.Marshal(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			utils.ErrorResponce{Message: "Failed to marshal request for kafka", Error: err})
		return
	}
	err = h.Producer.WriteMessage(ctx, kafka.DeviceUpdateLanguage, bytes)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			utils.ErrorResponce{Message: "Failed to write message to kafka", Error: err})
		return
	}
	ctx.JSON(http.StatusOK, nil)
}

type GeolocationRequest struct {
	UUID      string  `json:"uuid" example:"550e8400-e29b-41d4-a716-446655440000"`
	Longitude float64 `json:"longitude" example:"55.643881"`
	Latitude  float64 `json:"latitude" example:"37.551595"`
}

//	@Summary	Update a device geolocation
//	@Tags		Device UPDATE
//	@Accept		json
//	@Produce	json
//	@Param		request	body	GeolocationRequest	true	"Update Device geolocation Request"
//	@Success	200
//	@Failure	400	{object}	utils.ErrorResponce
//	@Failure	500	{object}	utils.ErrorResponce
//	@Router		/device_geo [put]
func (h *DeviceHandler) UpdateGeolocation(ctx *gin.Context) {
	var req GeolocationRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			utils.ErrorResponce{Message: "Failed to parse request", Error: err})
		return
	}
	if errResp := utils.UuidValidation(req.UUID); errResp != nil {
		ctx.JSON(http.StatusBadRequest, errResp)
		return
	}
	bytes, err := json.Marshal(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			utils.ErrorResponce{Message: "Failed to marshal request for kafka", Error: err})
		return
	}
	err = h.Producer.WriteMessage(ctx, kafka.DeviceUpdateGeoposition, bytes)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			utils.ErrorResponce{Message: "Failed to write message to kafka", Error: err})
		return
	}
	ctx.JSON(http.StatusOK, nil)
}

type EmailRequest struct {
	UUID  string `json:"uuid" example:"550e8400-e29b-41d4-a716-446655440000"`
	Email string `json:"email" example:"another@email.com"`
}

//	@Summary	Update a device E-mail
//	@Tags		Device UPDATE
//	@Accept		json
//	@Produce	json
//	@Param		request	body	EmailRequest	true	"Update Device E-mail Request"
//	@Success	200
//	@Failure	400	{object}	utils.ErrorResponce
//	@Failure	500	{object}	utils.ErrorResponce
//	@Router		/device_email [put]
func (h *DeviceHandler) UpdateEmail(ctx *gin.Context) {
	var req EmailRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			utils.ErrorResponce{Message: "Failed to parse request", Error: err})
		return
	}
	if errResp := utils.UuidValidation(req.UUID); errResp != nil {
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
			utils.ErrorResponce{Message: "Failed to marshal request for kafka", Error: err})
		return
	}
	err = h.Producer.WriteMessage(ctx, kafka.DeviceUpdateEmail, bytes)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			utils.ErrorResponce{Message: "Failed to write message to kafka", Error: err})
		return
	}
	ctx.JSON(http.StatusOK, nil)
}
