package device

import (
	"device-manager/internal/logger"
	"device-manager/internal/server/handlers/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//	@Summary		Get a device
//	@Description	Get device info
//	@Tags			Device GET
//	@Accept			json
//	@Produce		json
//	@Param			uuid	path		string	true	"Device UUID"
//	@Success		200		{object}	domain.Device
//	@Failure		400		{object}	utils.ErrorResponce
//	@Failure		404
//	@Router			/device/{uuid} [get]
func (h *DeviceHandler) Get(ctx *gin.Context) {
	req := ctx.Param("uuid")
	errResp := utils.UuidValidation(req)
	if errResp != nil {
		ctx.JSON(http.StatusBadRequest, errResp)
		return
	}
	device, err := h.service.Get(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusNotFound, nil)
		logger.Logger.Error("Device not found", zap.Error(err))
		return
	}
	ctx.JSON(http.StatusOK, device)
}

//	@Summary	Get devices by Language filter
//	@Description
//	@Tags		Device GET
//	@Accept		json
//	@Produce	json
//	@Param		language	path		string	true	"Devices Language"
//	@Success	200			{object}	[]domain.Device
//	@Failure	404
//	@Failure	500	{object}	utils.ErrorResponce
//	@Router		/device_lang/{language} [get]
func (h *DeviceHandler) GetByLanguage(ctx *gin.Context) {
	req := ctx.Param("language")
	if errResp := utils.LanguageValidation(req); errResp != nil {
		ctx.JSON(http.StatusBadRequest, errResp)
	}
	devices, err := h.service.GetByLanguage(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			utils.ErrorResponce{Message: "Failed to get devices by language"})
		logger.Logger.Error("Failed to get devices by language", zap.Error(err))
		return
	}
	if devices == nil {
		ctx.JSON(http.StatusNotFound, nil)
		return
	}
	ctx.JSON(http.StatusOK, devices)
}

//	@Summary	Get devices by Geoposition
//	@Description
//	@Tags		Device GET
//	@Accept		json
//	@Produce	json
//	@Param		longitude	query		number	true	"longitude"
//	@Param		latitude	query		number	true	"latitude"
//	@Param		distance	query		integer	true	"distance"
//	@Success	200			{object}	domain.Device
//	@Failure	400			{object}	utils.ErrorResponce
//	@Failure	404
//	@Failure	500	{object}	utils.ErrorResponce
//	@Router		/device_geo [get]
func (h *DeviceHandler) GetByGeolocation(ctx *gin.Context) {
	queryLong := ctx.Query("longitude")
	queryLat := ctx.Query("latitude")
	queryDist := ctx.Query("distance")

	longitude, err := strconv.ParseFloat(queryLong, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			utils.ErrorResponce{Message: "Failed to parse longitude", Error: err})
		return
	}
	latitude, err := strconv.ParseFloat(queryLat, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			utils.ErrorResponce{Message: "Failed to parse latitude", Error: err})
		return
	}
	distance, err := strconv.ParseInt(queryDist, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			utils.ErrorResponce{Message: "Failed to parse distance", Error: err})
		return
	}

	// longitude := ctx.GetFloat64("longitude")
	// latitude := ctx.GetFloat64("latitude")
	// distance := ctx.GetInt("distance")

	devices, err := h.service.GetByGeolocation(ctx, []float64{longitude, latitude}, int(distance))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			utils.ErrorResponce{Message: "Failed to get devices by geoposition", Error: err})
		logger.Logger.Error("Failed to get devices by geoposition", zap.Error(err))
		return
	}
	if devices == nil {
		ctx.JSON(http.StatusNotFound, nil)
		return
	}
	ctx.JSON(http.StatusOK, devices)
}

//	@Summary	Get devices by Email filter
//	@Description
//	@Tags		Device GET
//	@Accept		json
//	@Produce	json
//	@Param		email	path		string	true	"Devices Email"
//	@Success	200		{object}	domain.Device
//	@Failure	404
//	@Failure	500	{object}	utils.ErrorResponce
//	@Router		/device_email/{email} [get]
func (h *DeviceHandler) GetByEmail(ctx *gin.Context) {
	req := ctx.Param("email")
	if errResp := utils.EmailValidation(req); errResp != nil {
		ctx.JSON(http.StatusBadRequest, errResp)
	}
	devices, err := h.service.GetByEmail(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			utils.ErrorResponce{Message: "Failed to get devices by E-mail", Error: err})
		logger.Logger.Error("Failed to get devices by E-mail", zap.Error(err))
		return
	}
	if devices == nil {
		ctx.JSON(http.StatusNotFound, nil)
		return
	}
	ctx.JSON(http.StatusOK, devices)
}
