package mapper

import (
	"poc/cmd/repos/dto/response"
	wrapper "poc/cmd/repos/repository/github/wrapper/response"
	"poc/cmd/repos/utils"
	coreErrors "poc/commons/core/errors/errors"

	"github.com/mitchellh/mapstructure"
)

func ToResponseDto(repoResponse wrapper.RepoResponseWrapper) (*response.RepoResponseDto, error) {
	var result response.RepoResponseDto

	if err := mapstructure.Decode(repoResponse, &result); err != nil {
		return nil, coreErrors.NewMappingError(err.Error())
	}

	result.PushedAt = utils.FormatDate(result.PushedAt)
	return &result, nil
}
