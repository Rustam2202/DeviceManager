package device

import (
	"device-manager/internal/logger"
	"device-manager/internal/server/handlers"
	"device-manager/internal/server/handlers/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//	@Summary		Get a device
//	@Description	Get device info
//	@Tags			Device
//	@Accept			json
//	@Produce		json
//	@Param			uuid	path		string	true	"Device UUID"
//	@Success		200		{object}	domain.Device
//	@Failure		400		{object}	handlers.ErrorResponce
//	@Failure		404
//	@Router			/device/{uuid} [get]
func (h *DeviceHandler) Get(ctx *gin.Context) {
	req := ctx.Param("uuid")
	id, errResp := utils.UuidValidationAndParse(req)
	if errResp != nil {
		ctx.JSON(http.StatusBadRequest, errResp)
		return
	}
	device, err := h.service.Get(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, nil)
		logger.Logger.Error("Device not found", zap.Error(err))
		return
	}
	ctx.JSON(http.StatusOK, device)
}

//	@Summary	Get devices by Language filter
//	@Description
//	@Tags		Device
//	@Accept		json
//	@Produce	json
//	@Param		language	path		string	true	"Devices Language"
//	@Success	200			{object}	[]domain.Device
//	@Failure	404
//	@Failure	500	{object}	handlers.ErrorResponce
//	@Router		/device_lang/{language} [get]
func (h *DeviceHandler) GetByLanguage(ctx *gin.Context) {
	req := ctx.Param("language")
	if errResp := utils.LanguageValidation(req); errResp != nil {
		ctx.JSON(http.StatusBadRequest, errResp)
	}
	devices, err := h.service.GetByLanguage(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			handlers.ErrorResponce{Message: "Failed to get a devices from database", Error: err})
		logger.Logger.Error("Failed to get a person from database", zap.Error(err))
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
//	@Tags		Device
//	@Accept		json
//	@Produce	json
//	@Param		longitude	query		number	true	"longitude"
//	@Param		latitude	query		number	true	"latitude"
//	@Param		distance	query		integer	true	"distance"
//	@Success	200			{object}	domain.Device
//	@Failure	400			{object}	handlers.ErrorResponce
//	@Failure	404
//	@Failure	500	{object}	handlers.ErrorResponce
//	@Router		/device_geo{} [get]
func (h *DeviceHandler) GetByGeolocation(ctx *gin.Context) {
	query:=ctx.Query("longitude")
	fmt.Printf("query: %v\n", query)
	query1:=ctx.Query("latitude")
	fmt.Printf("query: %v\n", query1)
	query2:=ctx.Query("distance")
	fmt.Printf("query: %v\n", query2)

	longitude := ctx.GetFloat64("longitude")
	latitude := ctx.GetFloat64("latitude")
	distance := ctx.GetInt("distance")

	devices, err := h.service.GetByGeolocation(ctx, longitude, latitude, distance)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			handlers.ErrorResponce{Message: "Failed to get a person from database", Error: err})
		logger.Logger.Error("Failed to get devices from database", zap.Error(err))
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
//	@Tags		Device
//	@Accept		json
//	@Produce	json
//	@Param		email	path		string	true	"Devices Email"
//	@Success	200		{object}	domain.Device
//	@Failure	404
//	@Failure	500	{object}	handlers.ErrorResponce
//	@Router		/device_email/{email} [get]
func (h *DeviceHandler) GetByEmail(ctx *gin.Context) {
	req := ctx.Param("email")
	if errResp := utils.EmailValidation(req); errResp != nil {
		ctx.JSON(http.StatusBadRequest, errResp)
	}
	devices, err := h.service.GetByEmail(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			handlers.ErrorResponce{Message: "Failed to get devices from database", Error: err})
		logger.Logger.Error("Failed to get devices from database", zap.Error(err))
		return
	}
	if devices == nil {
		ctx.JSON(http.StatusNotFound, nil)
		return
	}
	ctx.JSON(http.StatusOK, devices)
}
