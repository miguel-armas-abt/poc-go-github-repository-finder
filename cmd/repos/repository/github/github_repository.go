package repository

import (
	"context"

	"com.demo.poc/cmd/repos/repository/github/wrapper/response"
)

type GitHubRepository interface {
	FindRepositoriesByOwner(
		ctx context.Context,
		headers map[string]string,
		owner string) ([]response.RepoResponseWrapper, error)
}
