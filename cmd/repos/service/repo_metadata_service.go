package service

import (
	"context"

	"com.demo.poc/cmd/repos/dto/request"
)

type ParameterService interface {
	InsertRepoMetadata(
		ctx context.Context,
		headers map[string]string,
		insertRequest request.RepoMetadataInsertRequest) error
}
