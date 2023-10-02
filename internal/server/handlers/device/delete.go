package device

import (
	"device-manager/internal/kafka"
	"device-manager/internal/server/handlers/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

//	@Summary		Delete a device
//	@Description	Delete a device from database
//	@Tags			Device
//	@Accept			json
//	@Produce		json
//	@Param			uuid	path	string	true	"Device Id"
//	@Success		200
//	@Failure		400	{object}	utils.ErrorResponce
//	@Failure		500	{object}	utils.ErrorResponce
//	@Router			/device/{uuid} [delete]
func (h *DeviceHandler) Delete(ctx *gin.Context) {
	req := ctx.Param("uuid")
	if errResp := utils.UuidValidation(req); errResp != nil {
		ctx.JSON(http.StatusBadRequest, errResp)
		return
	}
	bytes := []byte(req)
	err := h.Producer.WriteMessage(ctx, kafka.DeviceDelete, bytes)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			utils.ErrorResponce{Message: "Failed to write message to kafka", Error: err})
		return
	}
	ctx.JSON(http.StatusOK, nil)
}
