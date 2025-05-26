package errors

import (
	"net/http"

	errorDto "poc/commons/core/errors/dto"
)

// business and bad requests: 01.02.xx
func NewProfileNotFoundError(message string) errorDto.GenericError {
	return errorDto.GenericError{
		HttpStatus: http.StatusNotFound,
		Origin:     errorDto.ERROR_ORIGIN_OWN,
		Code:       "01.02.02",
		Message:    message,
	}
}
