package repository

import (
	"context"

	"com.demo.poc/cmd/repos/repository/github/wrapper/response"
)

type GitHubRepository interface {
	FindRepositoriesByProfile(
		ctx context.Context,
		headers map[string]string,
		profile string) ([]response.RepoResponseWrapper, error)
}
