package event

import (
	"device-manager/internal/logger"
	"device-manager/internal/server/handlers"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

//	@Summary		Get events
//	@Description	Get events of device from database
//	@Tags			Event
//	@Accept			json
//	@Produce		json
//	@Param			uuid		query		string	true	"UUID"
//	@Param			timeBegin	query		string	true	"Begin time range"
//	@Param			timeEnd		query		string	true	"End time range"
//	@Success		200			{object}	[]domain.Event
//	@Failure		400			{object}	handlers.ErrorResponce
//	@Failure		500			{object}	handlers.ErrorResponce
//	@Router			/event [get]
func (h *EventHandler) Get(ctx *gin.Context) {
	layout := "2006-01-02T15:04:05.999-07:00"

	uuidReq := ctx.Query("uuid")
	uuid, err := uuid.Parse(uuidReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			handlers.ErrorResponce{Message: "Failed to parse request", Error: err})
		logger.Logger.Error("Failed to parse request", zap.Error(err))
		return
	}
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

	events, err := h.Service.Get(ctx, uuid, tb, te, "")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			handlers.ErrorResponce{Message: "Failed to get event from database", Error: err})
		return
	}
	ctx.JSON(http.StatusOK, events)
}
