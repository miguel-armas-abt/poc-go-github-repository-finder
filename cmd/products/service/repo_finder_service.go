package service

import (
	"com.demo.poc/cmd/products/dto/response"
)

type RepoFinderService interface {
	FindRepositoriesByOwner(headers map[string]string, owner string) ([]response.RepoResponseDto, error)
}
