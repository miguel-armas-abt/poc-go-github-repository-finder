package repository

import "com.demo.poc/cmd/products/repository/github/wrapper/response"

type GitHubRepository interface {
	FindRepositoriesByOwner(headers map[string]string, owner string) ([]response.RepoResponseWrapper, error)
}
