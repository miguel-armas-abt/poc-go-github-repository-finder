package service

import (
	"context"

	"com.demo.poc/cmd/profile/dto/request"
	"com.demo.poc/cmd/profile/dto/response"
	"com.demo.poc/cmd/profile/mapper"
	repository "com.demo.poc/cmd/profile/repository/profile"
	"com.demo.poc/commons/properties"
)

type profileServiceImpl struct {
	profileRepository repository.ProfileRepository
	props             properties.ApplicationProperties
}

func NewProfileServiceImpl(
	profileRepository repository.ProfileRepository,
	props properties.ApplicationProperties) ProfileService {

	return &profileServiceImpl{
		profileRepository: profileRepository,
		props:             props,
	}
}

func (service *profileServiceImpl) InsertProfile(
	ctx context.Context,
	headers map[string]string,
	insertRequest request.ProfileInsertRequest) error {

	multimediaStorage := service.props.MultimediaStorage
	request, errorMapper := mapper.ToDocument(insertRequest, &multimediaStorage)
	if errorMapper != nil {
		return errorMapper
	}

	err := service.profileRepository.Insert(ctx, request)
	if err != nil {
		return err
	}
	return nil
}

func (service *profileServiceImpl) FindByUsername(
	ctx context.Context,
	headers map[string]string,
	username string) (*response.ProfileResponse, error) {

	profile, err := service.profileRepository.FindByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	response, err := mapper.ToResponse(*profile)
	if err != nil {
		return nil, err
	}

	return response, nil
}
