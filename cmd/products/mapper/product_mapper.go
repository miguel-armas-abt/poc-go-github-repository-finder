package mapper

import (
	"com.demo.poc/cmd/products/dto/response"
	responseWrapper "com.demo.poc/cmd/products/repository/github/wrapper/response"
	"github.com/mitchellh/mapstructure"
)

func ToResponseDto(repoWrapperList []responseWrapper.RepoResponseWrapper) []response.RepoResponseDto {
	productResponses := make([]response.RepoResponseDto, len(repoWrapperList))
	for i, wrapper := range repoWrapperList {
		productResponses[i] = toResponseDto(wrapper)
	}
	return productResponses
}

func toResponseDto(repoWrapper responseWrapper.RepoResponseWrapper) response.RepoResponseDto {
	var result response.RepoResponseDto
	mapstructure.Decode(repoWrapper, &result)
	return result
}
