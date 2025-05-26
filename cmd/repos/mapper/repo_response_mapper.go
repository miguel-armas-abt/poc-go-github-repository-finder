package mapper

import (
	"com.demo.poc/cmd/repos/dto/response"
	wrapper "com.demo.poc/cmd/repos/repository/github/wrapper/response"
	"com.demo.poc/cmd/repos/utils"
	repoErrors "com.demo.poc/commons/core/errors/errors"
	"github.com/mitchellh/mapstructure"
)

func ToResponseDto(repoResponse wrapper.RepoResponseWrapper) (*response.RepoResponseDto, error) {
	var result response.RepoResponseDto

	if err := mapstructure.Decode(repoResponse, &result); err != nil {
		return nil, repoErrors.NewMappingError(err.Error())
	}

	result.PushedAt = utils.FormatDate(result.PushedAt)
	return &result, nil
}
