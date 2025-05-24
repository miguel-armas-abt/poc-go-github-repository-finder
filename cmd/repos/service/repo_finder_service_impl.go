package service

import (
	"context"

	"com.demo.poc/cmd/repos/dto/response"
	mergeHelper "com.demo.poc/cmd/repos/helper"
	params "com.demo.poc/cmd/repos/params"
)

type repoFinderServiceImpl struct {
	repoMergeHelper *mergeHelper.RepoMergeHelper
}

func NewRepoFinderServiceImpl(
	repoMergeHelper *mergeHelper.RepoMergeHelper) RepoFinderService {

	return &repoFinderServiceImpl{
		repoMergeHelper: repoMergeHelper,
	}
}

func (service *repoFinderServiceImpl) FindRepositoriesByProfileAndLabel(
	ctx context.Context,
	headers map[string]string,
	params *params.RepoFinderParams) ([]response.RepoResponseDto, error) {

	repositories, err := service.repoMergeHelper.MergeRepositoriesByProfileAndLabel(ctx, headers, params)
	if err != nil {
		return nil, err
	}
	return repositories, nil
}
