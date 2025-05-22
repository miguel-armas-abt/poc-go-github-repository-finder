package service

import (
	"context"

	"com.demo.poc/cmd/parameters/dto/request"
)

type ParameterService interface {
	InsertRepoParameter(
		ctx context.Context,
		headers map[string]string,
		insertRequest request.ParameterInsertRequest) error
}
