package service

import (
	"context"

	"com.demo.poc/cmd/parameters/dto/request"
	"com.demo.poc/cmd/parameters/mapper"
	repository "com.demo.poc/cmd/parameters/repository/parameters"
)

type parameterServiceImpl struct {
	parameterRepository repository.ParameterRepository
}

func NewParameterServiceImpl(
	parameterRepository repository.ParameterRepository) ParameterService {

	return &parameterServiceImpl{
		parameterRepository: parameterRepository,
	}
}

func (service *parameterServiceImpl) InsertRepoParameter(
	ctx context.Context,
	headers map[string]string,
	insertRequest request.ParameterInsertRequest) error {

	request := mapper.ToDocument(insertRequest)
	err := service.parameterRepository.Insert(ctx, &request)
	if err != nil {
		return err
	}
	return nil
}
