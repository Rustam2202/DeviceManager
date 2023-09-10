package device

import (
	"device-manager/internal/server/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary		Get a device
// @Description	Get a device from database
// @Tags			Device
// @Accept			json
// @Produce		json
// @Param		uuid     path    string     true        "Device UUID"
// @Success		200		{object}	domain.Device
// @Failure		400		{object}	handlers.ErrorResponce
// @Failure		500		{object}	handlers.ErrorResponce
// @Router			/device/{uuid} [get]
func (h *DeviceHandler) GetDevice(ctx *gin.Context) {
	req := ctx.Param("uuid")
	device, err := h.service.GetDeviceInfo(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			handlers.ErrorResponce{Message: "Failed to get a person from database", Error: err})
		return
	}
	ctx.JSON(http.StatusOK, device)
}
