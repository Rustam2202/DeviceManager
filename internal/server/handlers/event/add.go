package event

import (
	"device-manager/internal/kafka"
	"device-manager/internal/server/handlers"
	"device-manager/internal/server/handlers/utils"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type AddEventRequest struct {
	UUID       string        `json:"uuid" example:"550e8400-e29b-41d4-a716-446655440000"`
	Name       string        `json:"name" example:"device event"`
	Attributes []interface{} `json:"attributes"`
}

type addEventRequestOutput struct {
	UUID       string        `json:"uuid"`
	Name       string        `json:"name"`
	Attributes []interface{} `json:"attributes"`
	Time       time.Time     `json:"created_at"`
}

//	@Summary		Add event
//	@Description	Add a new event from device to database
//	@Tags			Event
//	@Accept			json
//	@Produce		json
//	@Param			request	body	AddEventRequest	true	"Add Event Request"
//	@Success		200
//	@Failure		400	{object}	handlers.ErrorResponce
//	@Failure		500	{object}	handlers.ErrorResponce
//	@Router			/event [post]
func (h *EventHandler) Add(ctx *gin.Context) {
	createdAt := time.Now()
	var req AddEventRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			handlers.ErrorResponce{Message: "Failed to parse request", Error: err})
		return
	}
	if errResp := utils.UuidValidation(req.UUID); errResp != nil {
		ctx.JSON(http.StatusBadRequest, errResp)
		return
	}
	req.Attributes = utils.AttributesValidation(req.Attributes)
	outputReq := addEventRequestOutput{
		UUID:       req.UUID,
		Name:       req.Name,
		Attributes: req.Attributes,
		Time:       createdAt,
	}
	bytes, err := json.Marshal(outputReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			handlers.ErrorResponce{Message: "Failed to marshal request for kafka", Error: err})
		return
	}
	err = h.Producer.WriteMessage(ctx, kafka.EventCreate, bytes)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			handlers.ErrorResponce{Message: "Failed to write message to kafka", Error: err})
		return
	}
	ctx.JSON(http.StatusOK, nil)
}
