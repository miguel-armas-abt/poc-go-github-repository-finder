package errors

import (
	"net/http"

	errorDto "com.demo.poc/commons/errors/dto"
)

// business and bad requests: 01.02.xx
func NewInvalidFieldError(message string) errorDto.GenericError {
	return errorDto.GenericError{
		HttpStatus: http.StatusBadRequest,
		Origin:     errorDto.ERROR_ORIGIN_OWN,
		Code:       "01.02.01",
		Message:    message,
	}
}
