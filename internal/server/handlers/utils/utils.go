package utils

import (
	"net/http"
	"net/mail"

	"github.com/google/uuid"
	"github.com/mssola/useragent"
	"golang.org/x/text/language"
)

func LanguageValidation(input string) *ErrorResponce {
	_, err := language.Parse(input)
	if err != nil {
		return &ErrorResponce{Code: http.StatusBadRequest, Message: "Failed to parse language", Error: err}
	}
	return nil
}

func UuidValidation(input string) *ErrorResponce {
	_, err := uuid.ParseBytes([]byte(input))
	if err != nil {
		return &ErrorResponce{Message: "Failed to parse UUID", Error: err}
	}
	return nil
}

func EmailValidation(input string) *ErrorResponce {
	_, err := mail.ParseAddress(input)
	if err != nil {
		return &ErrorResponce{Message: "Failed to parse E-mail", Error: err}
	}
	return nil
}

func PlatformValidation(input string) *ErrorResponce {
	platform := useragent.New(input)
	if platform == nil {
		return &ErrorResponce{Message: "Empty Platform request"}
	}
	if platform.OS() == "" {
		return &ErrorResponce{Message: "Empty OS value"}
	}
	if platform.Platform() == "" {
		return &ErrorResponce{Message: "Empty platform value"}
	}
	return nil
}

func AttributesValidation(attr []interface{}) []interface{} {
	var validAttributes []interface{}
	for _, attr := range attr {
		switch v := attr.(type) {
		case string:
			validAttributes = append(validAttributes, v)
		case int, int8, int16, int32, int64:
			validAttributes = append(validAttributes, v)
		case uint, uint8, uint16, uint32, uint64:
			validAttributes = append(validAttributes, v)
		case float32, float64:
			validAttributes = append(validAttributes, v)
		case bool:
			validAttributes = append(validAttributes, v)
		default:
			continue
		}
	}
	return validAttributes
}
