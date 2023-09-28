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
// @Success		200
// @Failure		400	{object}	handlers.ErrorResponce
// @Failure		500	{object}	handlers.ErrorResponce
// @Router			/event [post]
func (h *EventHandler) Add(ctx *gin.Context) {
	var req AddEventRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			handlers.ErrorResponce{Message: "Failed to parse request", Error: err})
		return
	}
	if _, errResp := utils.UuidValidationAndParse(req.UUID); errResp != nil {
		ctx.JSON(http.StatusBadRequest, errResp)
		return
	}
	req.Attributes = attributesValidation(req.Attributes)
	req.CreatedAt = time.Now()
	bytes, err := json.Marshal(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			handlers.ErrorResponce{Message: "Failed to marshal request for kafka", Error: err})
		return
	}
	err = h.Producer.WriteMessage(ctx, kafka.EventCreate, bytes)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			handlers.ErrorResponce{Message: "Failed to add a new person to database", Error: err})
		return
	}
	ctx.JSON(http.StatusOK, nil)
}

func attributesValidation(attr []interface{}) []interface{} {
	var validAttributes []interface{}
	for _, attr := range attr {
		switch v := attr.(type) {
		case string:
			validAttributes = append(validAttributes, v)
		case int:
			validAttributes = append(validAttributes, v)
		case int8:
			validAttributes = append(validAttributes, v)
		case int16:
			validAttributes = append(validAttributes, v)
		case int32:
			validAttributes = append(validAttributes, v)
		case int64:
			validAttributes = append(validAttributes, v)
		case uint:
			validAttributes = append(validAttributes, v)
		case uint8:
			validAttributes = append(validAttributes, v)
		case uint16:
			validAttributes = append(validAttributes, v)
		case uint32:
			validAttributes = append(validAttributes, v)
		case uint64:
			validAttributes = append(validAttributes, v)
		case float32:
			validAttributes = append(validAttributes, v)
		case float64:
			validAttributes = append(validAttributes, v)
		case bool:
			validAttributes = append(validAttributes, v)
		default:
			continue
		}
	}
	return validAttributes
}
