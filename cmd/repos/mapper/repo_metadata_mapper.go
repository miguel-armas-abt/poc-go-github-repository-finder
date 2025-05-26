package mapper

import (
	"strings"

	requestDto "poc/cmd/repos/dto/request"
	"poc/cmd/repos/repository/metadata/document"
	"poc/commons/core/constants"
	coreErrors "poc/commons/core/errors/errors"

	"github.com/mitchellh/mapstructure"
)

func ToDocument(request requestDto.RepoMetadataInsertRequest, multimediaStorage *string) (*document.RepoMetadataDocument, error) {
	var result document.RepoMetadataDocument

	if err := mapstructure.Decode(request, &result); err != nil {
		return nil, coreErrors.NewMappingError(err.Error())
	}

	*multimediaStorage = strings.ReplaceAll(*multimediaStorage, "$USER", request.Profile)
	*multimediaStorage = *multimediaStorage + constants.SLASH + request.Profile + "/img/" + request.RepositoryName + ".webp"
	result.ImageUrl = *multimediaStorage

	return &result, nil
}
