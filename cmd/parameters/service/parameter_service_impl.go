package service

import (
	"context"

	"com.demo.poc/cmd/parameters/dto/request"
	"com.demo.poc/cmd/parameters/mapper"
	repository "com.demo.poc/cmd/parameters/repository/parameters"
	"com.demo.poc/commons/properties"
)

type parameterServiceImpl struct {
	parameterRepository repository.ParameterRepository
	props               properties.ApplicationProperties
}

func NewParameterServiceImpl(
	parameterRepository repository.ParameterRepository,
	props properties.ApplicationProperties) ParameterService {

	return &parameterServiceImpl{
		parameterRepository: parameterRepository,
		props:               props,
	}
}

func (service *parameterServiceImpl) InsertRepoParameter(
	ctx context.Context,
	headers map[string]string,
	insertRequest request.ParameterInsertRequest) error {

	multimediaStorage := service.props.MultimediaStorage
	request := mapper.ToDocument(insertRequest, &multimediaStorage)
	err := service.parameterRepository.Insert(ctx, &request)
	if err != nil {
		return err
	}
	return nil
}
