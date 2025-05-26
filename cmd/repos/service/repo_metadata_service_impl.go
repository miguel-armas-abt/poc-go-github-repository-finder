package service

import (
	"context"

	"poc/cmd/repos/dto/request"
	"poc/cmd/repos/mapper"
	repository "poc/cmd/repos/repository/metadata"
	"poc/commons/custom/properties"
)

type repoMetadataServiceImpl struct {
	parameterRepository repository.RepoMetadataRepository
}

func NewRepoMetadataServiceImpl(
	parameterRepository repository.RepoMetadataRepository) ParameterService {

	return &repoMetadataServiceImpl{
		parameterRepository: parameterRepository,
	}
}

func (service *repoMetadataServiceImpl) InsertRepoMetadata(
	ctx context.Context,
	headers map[string]string,
	insertRequest request.RepoMetadataInsertRequest) error {

	multimediaStorage := properties.Properties.MultimediaStorage
	request, mapperError := mapper.ToDocument(insertRequest, &multimediaStorage)
	if mapperError != nil {
		return mapperError
	}

	err := service.parameterRepository.Insert(ctx, request)
	if err != nil {
		return err
	}
	return nil
}
