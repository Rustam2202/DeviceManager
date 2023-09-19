package device

import (
	"device-manager/internal/kafka"
	"device-manager/internal/server/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary		Delete a device
// @Description	Delete a device from database
// @Tags			Device
// @Accept			json
// @Produce		json
// @Param			uuid	path	string	true	"Device Id"
// @Success		200
// @Failure		400	{object}	handlers.ErrorResponce
// @Failure		500	{object}	handlers.ErrorResponce
// @Router			/device/{uuid} [delete]
func (h *DeviceHandler) Delete(ctx *gin.Context) {
	req := ctx.Param("uuid")
	// err := h.service.Delete(ctx, req)
	bytes := []byte(req)
	err := h.Producer.WriteMessage(ctx, kafka.DeviceDelete, bytes)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			handlers.ErrorResponce{Message: "Failed to get a person from database", Error: err})
		return
	}
	ctx.JSON(http.StatusOK, nil)
}
