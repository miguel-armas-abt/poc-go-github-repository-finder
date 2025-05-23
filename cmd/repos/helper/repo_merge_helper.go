package helper

import (
	"context"

	"golang.org/x/sync/errgroup"

	paramRepository "com.demo.poc/cmd/parameters/repository/parameters"
	paramDocument "com.demo.poc/cmd/parameters/repository/parameters/document"
	dto "com.demo.poc/cmd/repos/dto/response"
	"com.demo.poc/cmd/repos/mapper"
	params "com.demo.poc/cmd/repos/params"
	repoRepository "com.demo.poc/cmd/repos/repository/github"
	repoWrapper "com.demo.poc/cmd/repos/repository/github/wrapper/response"
)

type RepoMergeHelper struct {
	gitHubRepository    repoRepository.GitHubRepository
	parameterRepository paramRepository.ParameterRepository
}

func NewRepoMergeHelper(
	gitHubRepository repoRepository.GitHubRepository,
	parameterRepository paramRepository.ParameterRepository,
) *RepoMergeHelper {

	return &RepoMergeHelper{
		gitHubRepository:    gitHubRepository,
		parameterRepository: parameterRepository,
	}
}

func (helper *RepoMergeHelper) MergeRepositoriesByOwnerAndLabel(
	ctx context.Context,
	headers map[string]string,
	params *params.RepoFinderParams) ([]dto.RepoResponseDto, error) {

	var (
		group        errgroup.Group
		repositories []repoWrapper.RepoResponseWrapper
		parameters   []*paramDocument.ParameterDocument
		err          error
	)

	group.Go(func() error {
		repositories, err = helper.gitHubRepository.FindRepositoriesByOwner(ctx, headers, params.Owner)
		return err
	})

	group.Go(func() error {
		parameters, err = helper.parameterRepository.FindByOwnerAndLabel(ctx, params)
		return err
	})

	if err = group.Wait(); err != nil {
		return nil, err
	}

	parameterMap := make(map[string]*paramDocument.ParameterDocument, len(parameters))
	for _, param := range parameters {
		parameterMap[param.RepositoryName] = param
	}

	result := make([]dto.RepoResponseDto, len(parameters))

	var count int = 0
	for _, repo := range repositories {
		if document, exists := parameterMap[repo.Name]; exists {
			response, err := mapper.ToResponseDto(repo)

			if err != nil {
				return nil, err
			}

			response.ImageUrl = document.ImageUrl
			result[count] = *response
			count++
		}

	}

	return result, nil
}
