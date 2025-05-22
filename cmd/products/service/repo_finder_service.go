package service

import (
	"context"

	"com.demo.poc/cmd/products/dto/response"
)

type RepoFinderService interface {
	FindRepositoriesByOwner(
		ctx context.Context,
		headers map[string]string,
		owner string) ([]response.RepoResponseDto, error)
}
