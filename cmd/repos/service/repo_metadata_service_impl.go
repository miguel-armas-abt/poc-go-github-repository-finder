package service

import (
	"context"

	"com.demo.poc/cmd/repos/dto/request"
	"com.demo.poc/cmd/repos/mapper"
	repository "com.demo.poc/cmd/repos/repository/metadata"
	"com.demo.poc/commons/properties"
)

type repoMetadataServiceImpl struct {
	parameterRepository repository.RepoMetadataRepository
	props               properties.ApplicationProperties
}

func NewRepoMetadataServiceImpl(
	parameterRepository repository.RepoMetadataRepository,
	props properties.ApplicationProperties) ParameterService {

	return &repoMetadataServiceImpl{
		parameterRepository: parameterRepository,
		props:               props,
	}
}

func (service *repoMetadataServiceImpl) InsertRepoMetadata(
	ctx context.Context,
	headers map[string]string,
	insertRequest request.RepoMetadataInsertRequest) error {

	multimediaStorage := service.props.MultimediaStorage
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
