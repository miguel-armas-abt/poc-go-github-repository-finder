package service

import (
	"context"

	"com.demo.poc/cmd/repos/dto/response"
	params "com.demo.poc/cmd/repos/params"
)

type RepoFinderService interface {
	FindRepositoriesByProfileAndLabel(
		ctx context.Context,
		headers map[string]string,
		params *params.RepoFinderParams) ([]response.RepoResponseDto, error)
}
