package service

import (
	"context"

	"com.demo.poc/cmd/products/dto/response"
	"com.demo.poc/cmd/products/mapper"
	gitHubRepository "com.demo.poc/cmd/products/repository/github"
)

type repoFinderServiceImpl struct {
	githubRepository gitHubRepository.GitHubRepository
}

func NewRepoFinderServiceImpl(
	githubRepository gitHubRepository.GitHubRepository) RepoFinderService {

	return &repoFinderServiceImpl{
		githubRepository: githubRepository,
	}
}

func (service *repoFinderServiceImpl) FindRepositoriesByOwner(
	ctx context.Context,
	headers map[string]string,
	owner string) ([]response.RepoResponseDto, error) {

	repoList, err := service.githubRepository.FindRepositoriesByOwner(ctx, headers, owner)
	if err != nil {
		return nil, err
	}
	product := mapper.ToResponseDto(repoList)
	return product, nil
}
