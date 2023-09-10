package device

import (
	"device-manager/internal/server/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateLanguageRequest struct {
	UUID     string `json:"uuid"`
	Language string `json:"language"`
}

// @Summary		Update a device language
// @Tags			Device language
// @Accept			json
// @Produce		json
// @Param			request	body		UpdateLanguageRequest	true	"Update Device language Request"
// @Success		200		{object}	UpdateLanguageRequest
// @Failure		400		{object}	handlers.ErrorResponce
// @Failure		500		{object}	handlers.ErrorResponce
// @Router			/device [put]
func (h *DeviceHandler) UpdateLanguage(ctx *gin.Context) {
	var req UpdateLanguageRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			handlers.ErrorResponce{Message: "Failed to parse request", Error: err})
		return
	}
	err = h.service.UpdateLaguage(ctx, req.UUID, req.Language)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			handlers.ErrorResponce{Message: "Failed to add a new person to database", Error: err})
		return
	}
	ctx.JSON(http.StatusOK, nil)
}

type UpdateGeolocationRequest struct {
	UUID        string `json:"uuid"`
	Geolocation string `json:"geolocation"`
}

// @Summary		Update a device geolocation
// @Tags			Device geolocotaion
// @Accept			json
// @Produce		json
// @Param			request	body		UpdateGeolocationRequest	true	"Update Device geolocation Request"
// @Success		200		{object}	UpdateGeolocationRequest
// @Failure		400		{object}	handlers.ErrorResponce
// @Failure		500		{object}	handlers.ErrorResponce
// @Router			/device [put]
func (h *DeviceHandler) UpdateGeolocation(ctx *gin.Context) {
	var req UpdateGeolocationRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			handlers.ErrorResponce{Message: "Failed to parse request", Error: err})
		return
	}
	err = h.service.UpdateGeolocation(ctx, req.UUID, req.Geolocation)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			handlers.ErrorResponce{Message: "Failed to add a new person to database", Error: err})
		return
	}
	ctx.JSON(http.StatusOK, nil)
}

type UpdateEmailRequest struct {
	UUID  string `json:"uuid"`
	Email string `json:"email"`
}

// @Summary		Update a device E-mail
// @Tags			Device E-mail
// @Accept			json
// @Produce		json
// @Param			request	body		UpdateEmailRequest	true	"Update Device E-mail Request"
// @Success		200		{object}	UpdateEmailRequest
// @Failure		400		{object}	handlers.ErrorResponce
// @Failure		500		{object}	handlers.ErrorResponce
// @Router			/device [put]
func (h *DeviceHandler) UpdateEmail(ctx *gin.Context) {
	var req UpdateEmailRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			handlers.ErrorResponce{Message: "Failed to parse request", Error: err})
		return
	}
	err = h.service.UpdateEmail(ctx, req.UUID, req.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			handlers.ErrorResponce{Message: "Failed to add a new person to database", Error: err})
		return
	}
	ctx.JSON(http.StatusOK, nil)
}
