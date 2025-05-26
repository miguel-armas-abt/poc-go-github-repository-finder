package service

import (
	"context"

	"poc/cmd/repos/dto/response"
	params "poc/cmd/repos/params"
)

type RepoFinderService interface {
	FindRepositoriesByProfileAndLabel(
		ctx context.Context,
		headers map[string]string,
		params *params.RepoFinderParams) ([]response.RepoResponseDto, error)
}
