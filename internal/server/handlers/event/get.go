package event

import (
	"device-manager/internal/server/handlers"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summary		Get events
// @Description	Get events of device from database
// @Tags			Event
// @Accept			json
// @Produce		json
// @Param			uuid		query		string	true	"UUID"
// @Param			timeBegin	query		string	true	"Start time"
// @Param			timeEnd		query		string	true	"End time"
// @Success		200			{object}	[]domain.Event
// @Failure		400			{object}	handlers.ErrorResponce
// @Failure		500			{object}	handlers.ErrorResponce
// @Router			/event [get]
func (h *EventHandler) GetEvents(ctx *gin.Context) {
	layout := "2006-01-02T15:04:05.999-07:00"
	uuid := ctx.Query("uuid")
	timeBegin := ctx.Query("timeBegin")
	timeEnd := ctx.Query("timeEnd")

	tb, err := time.Parse(layout, timeBegin)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			handlers.ErrorResponce{Message: "Failed to parse begin date", Error: err})
		return
	}
	te, err := time.Parse(layout, timeEnd)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			handlers.ErrorResponce{Message: "Failed to parse end date", Error: err})
		return
	}
	
	events, err := h.Service.GetDeviceEvents(ctx, uuid, tb, te)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			handlers.ErrorResponce{Message: "Failed to add a new person to database", Error: err})
		return
	}
	ctx.JSON(http.StatusCreated, events)
}
