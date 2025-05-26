package mapper

import (
	"strings"

	requestDto "com.demo.poc/cmd/repos/dto/request"
	"com.demo.poc/cmd/repos/repository/metadata/document"
	"com.demo.poc/commons/core/constants"
	repoErrors "com.demo.poc/commons/core/errors/errors"
	"github.com/mitchellh/mapstructure"
)

func ToDocument(request requestDto.RepoMetadataInsertRequest, multimediaStorage *string) (*document.RepoMetadataDocument, error) {
	var result document.RepoMetadataDocument

	if err := mapstructure.Decode(request, &result); err != nil {
		return nil, repoErrors.NewMappingError(err.Error())
	}

	*multimediaStorage = strings.ReplaceAll(*multimediaStorage, "$USER", request.Profile)
	*multimediaStorage = *multimediaStorage + constants.SLASH + request.Profile + "/img/" + request.RepositoryName + ".webp"
	result.ImageUrl = *multimediaStorage

	return &result, nil
}
