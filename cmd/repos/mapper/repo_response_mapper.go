package mapper

import (
	"com.demo.poc/cmd/repos/dto/response"
	repoErrors "com.demo.poc/cmd/repos/errors"
	wrapper "com.demo.poc/cmd/repos/repository/github/wrapper/response"
	"com.demo.poc/cmd/repos/utils"
	"github.com/mitchellh/mapstructure"
)

func ToResponseDto(w wrapper.RepoResponseWrapper) (*response.RepoResponseDto, error) {
	var dto response.RepoResponseDto

	if err := mapstructure.Decode(w, &dto); err != nil {
		return nil, repoErrors.NewMappingError(err.Error())
	}

	dto.PushedAt = utils.FormatDate(dto.PushedAt)
	return &dto, nil
}
