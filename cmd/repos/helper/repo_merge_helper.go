package helper

import (
	"context"

	"golang.org/x/sync/errgroup"

	dto "poc/cmd/repos/dto/response"
	"poc/cmd/repos/mapper"
	params "poc/cmd/repos/params"
	repoRepository "poc/cmd/repos/repository/github"
	repoWrapper "poc/cmd/repos/repository/github/wrapper/response"
	paramRepository "poc/cmd/repos/repository/metadata"
	metadataDocument "poc/cmd/repos/repository/metadata/document"
)

type RepoMergeHelper struct {
	gitHubRepository    repoRepository.GitHubRepository
	parameterRepository paramRepository.RepoMetadataRepository
}

func NewRepoMergeHelper(
	gitHubRepository repoRepository.GitHubRepository,
	parameterRepository paramRepository.RepoMetadataRepository,
) *RepoMergeHelper {

	return &RepoMergeHelper{
		gitHubRepository:    gitHubRepository,
		parameterRepository: parameterRepository,
	}
}

func (helper *RepoMergeHelper) MergeRepositoriesByProfileAndLabel(
	ctx context.Context,
	headers map[string]string,
	params *params.RepoFinderParams) ([]dto.RepoResponseDto, error) {

	var (
		group        errgroup.Group
		repositories []repoWrapper.RepoResponseWrapper
		parameters   []*metadataDocument.RepoMetadataDocument
		err          error
	)

	group.Go(func() error {
		repositories, err = helper.gitHubRepository.FindRepositoriesByProfile(ctx, headers, params.Profile)
		return err
	})

	group.Go(func() error {
		parameters, err = helper.parameterRepository.FindByProfileAndLabel(ctx, params)
		return err
	})

	if err = group.Wait(); err != nil {
		return nil, err
	}

	parameterMap := make(map[string]*metadataDocument.RepoMetadataDocument, len(parameters))
	for _, param := range parameters {
		parameterMap[param.RepositoryName] = param
	}

	result := make([]dto.RepoResponseDto, len(parameters))

	var count int = 0
	for _, repo := range repositories {
		if document, exists := parameterMap[repo.Name]; exists {
			response, err := mapper.ToResponseDto(repo, *document)

			if err != nil {
				return nil, err
			}

			result[count] = *response
			count++
		}

	}

	return result, nil
}
