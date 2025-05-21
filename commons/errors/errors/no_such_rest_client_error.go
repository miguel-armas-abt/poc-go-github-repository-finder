package errors

import (
	"net/http"

	errorDto "com.demo.poc/commons/errors/dto"
)

func NoSuchRestClientError(message string) errorDto.GenericError {
	return errorDto.GenericError{
		HttpStatus: http.StatusInternalServerError,
		Origin:     errorDto.ERROR_ORIGIN_OWN,
		Code:       "01.00.02",
		Message:    message,
	}
}
