package mapper

import (
	"strings"

	"poc/cmd/profile/dto/response"
	"poc/cmd/profile/repository/profile/document"
	"poc/commons/core/constants"
	coreErrors "poc/commons/core/errors/errors"

	"github.com/mitchellh/mapstructure"
)

func ToResponse(document document.ProfileDocument, multimediaStorage *string) (*response.ProfileResponse, error) {
	var result response.ProfileResponse

	if err := mapstructure.Decode(document, &result); err != nil {
		return nil, coreErrors.NewMappingError(err.Error())
	}

	*multimediaStorage = strings.ReplaceAll(*multimediaStorage, "$USER", document.Username)
	*multimediaStorage = *multimediaStorage + constants.SLASH + document.Username + "/img/" + document.Username + ".webp"
	result.PhotoUrl = *multimediaStorage

	return &result, nil
}
