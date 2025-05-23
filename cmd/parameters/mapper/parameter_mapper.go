package mapper

import (
	"strings"

	requestDto "com.demo.poc/cmd/parameters/dto/request"
	"com.demo.poc/cmd/parameters/repository/parameters/document"
	"github.com/mitchellh/mapstructure"
)

func ToDocument(request requestDto.ParameterInsertRequest, multimediaStorage *string) document.ParameterDocument {
	var result document.ParameterDocument
	mapstructure.Decode(request, &result)

	*multimediaStorage = strings.ReplaceAll(*multimediaStorage, "$USER", request.Owner)
	*multimediaStorage = strings.ReplaceAll(*multimediaStorage, "$REPOSITORY_NAME", request.RepositoryName)

	result.ImageUrl = *multimediaStorage
	return result
}
