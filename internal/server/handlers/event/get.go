package event

import (
	"device-manager/internal/server/handlers"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type GetEventsRequest struct {
	UUID      string
	TimeBegin time.Time
	TimeEnd   time.Time
}

// @Summary		Get events
// @Description	Get events of device from database
// @Tags			Event
// @Accept			json
// @Produce		json
// @Param		request	body		GetEventsRequest	true	"Get Events Request"
// @Success		200		{object}	domain.Event
// @Failure		400		{object}	handlers.ErrorResponce
// @Failure		500		{object}	handlers.ErrorResponce
// @Router			/device [get]
func (h *EventHandler) GetEvents(ctx *gin.Context) {
	var req GetEventsRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			handlers.ErrorResponce{Message: "Failed to parse request", Error: err})
		return
	}
	events, err := h.service.GetDeviceEvents(ctx, req.UUID, req.TimeBegin, req.TimeEnd)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			handlers.ErrorResponce{Message: "Failed to add a new person to database", Error: err})
		return
	}
	ctx.JSON(http.StatusCreated, events)
}
