package utils

import (
	"device-manager/internal/server/handlers"
	"net/http"
	"net/mail"

	"github.com/google/uuid"
	"golang.org/x/text/language"
)

func LanguageValidation(input string) *handlers.ErrorResponce {
	_, err := language.Parse(input)
	if err != nil {
		return &handlers.ErrorResponce{Code: http.StatusBadRequest, Message: "Failed to parse language", Error: err}
	}
	return nil
}

func UuidValidationAndParse(input string) (uuid.UUID, *handlers.ErrorResponce) {
	id, err := uuid.ParseBytes([]byte(input))
	if err != nil {
		return uuid.UUID{}, &handlers.ErrorResponce{Message: "Failed to parse UUID", Error: err}
	}
	return id, nil
}

func EmailValidation(input string) *handlers.ErrorResponce {
	_, err := mail.ParseAddress(input)
	if err != nil {
		return &handlers.ErrorResponce{Message: "Failed to parse E-ail", Error: err}
	}
	return nil
}

func CoordinatesValidation(coord []float64) *handlers.ErrorResponce {
	if len(coord) != 2 {
		return &handlers.ErrorResponce{Message: "Must be 2 coordinates values"}
	}
	return nil
}
