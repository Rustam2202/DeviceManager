package event

import (
	"device-manager/internal/kafka"
	"device-manager/internal/server/handlers"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type AddEventRequest struct {
	UUID       string
	Name       string
	CreatedAt  time.Time
	Attributes []interface{}
}

// @Summary		Add event
// @Description	Add a new event from device to database
// @Tags			Event
// @Accept			json
// @Produce		json
// @Param			request	body	AddEventRequest	true	"Add Event Request"
// @Success		201
// @Failure		400		{object}	handlers.ErrorResponce
// @Failure		500		{object}	handlers.ErrorResponce
// @Router			/event [post]
func (h *EventHandler) AddEventRequest(ctx *gin.Context) {
	var req AddEventRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			handlers.ErrorResponce{Message: "Failed to parse request", Error: err})
		return
	}
	req.CreatedAt = time.Now()
	bytes, err := json.Marshal(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			handlers.ErrorResponce{Message: "Failed to marshal request for kafka", Error: err})
		return
	}
	err = h.Producer.WriteMessage(ctx, kafka.EventCreate, bytes)
	// err = h.service.CreateEvent(ctx, req.UUID, req.Name, req.Attributes)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			handlers.ErrorResponce{Message: "Failed to add a new person to database", Error: err})
		return
	}
	ctx.JSON(http.StatusCreated, nil)
}
