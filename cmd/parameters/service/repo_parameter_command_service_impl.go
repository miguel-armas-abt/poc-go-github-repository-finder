package service

import (
	"context"

	"com.demo.poc/cmd/parameters/dto/request"
	"com.demo.poc/cmd/parameters/mapper"
	parameterRepository "com.demo.poc/cmd/parameters/repository/parameters"
)

type repoParameterCommandServiceImpl struct {
	parameterRepository parameterRepository.RepoParameterRepository
}

func NewRepoParameterCommandServiceImpl(
	paramterRepository parameterRepository.RepoParameterRepository) RepoParameterCommandService {

	return &repoParameterCommandServiceImpl{
		parameterRepository: paramterRepository,
	}
}

func (service *repoParameterCommandServiceImpl) InsertRepoParameter(
	ctx context.Context,
	headers map[string]string,
	insertRequest request.RepoParameterInsertRequest) error {

	request := mapper.ToDocument(insertRequest)
	err := service.parameterRepository.Insert(ctx, &request)
	if err != nil {
		return err
	}
	return nil
}
