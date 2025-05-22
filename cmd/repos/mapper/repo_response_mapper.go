package mapper

import (
	"fmt"

	"com.demo.poc/cmd/repos/dto/response"
	wrapper "com.demo.poc/cmd/repos/repository/github/wrapper/response"
	"com.demo.poc/cmd/repos/utils"
	"github.com/mitchellh/mapstructure"
)

func ToResponseDto(w wrapper.RepoResponseWrapper) response.RepoResponseDto {
	var dto response.RepoResponseDto

	if err := mapstructure.Decode(w, &dto); err != nil {
		panic(fmt.Sprintf("Mapping could not be performed: %v", err))
	}

	dto.PushedAt = utils.FormatDate(dto.PushedAt)
	return dto
}
