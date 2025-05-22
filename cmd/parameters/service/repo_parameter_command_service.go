package service

import (
	"context"

	"com.demo.poc/cmd/parameters/dto/request"
)

type RepoParameterCommandService interface {
	InsertRepoParameter(
		ctx context.Context,
		headers map[string]string,
		insertRequest request.RepoParameterInsertRequest) error
}
