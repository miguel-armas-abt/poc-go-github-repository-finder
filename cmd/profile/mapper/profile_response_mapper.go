package mapper

import (
	"strings"

	"com.demo.poc/cmd/profile/dto/response"
	"com.demo.poc/cmd/profile/repository/profile/document"
	"com.demo.poc/commons/constants"
	repoErrors "com.demo.poc/commons/errors/errors"
	"github.com/mitchellh/mapstructure"
)

func ToResponse(document document.ProfileDocument, multimediaStorage *string) (*response.ProfileResponse, error) {
	var result response.ProfileResponse

	if err := mapstructure.Decode(document, &result); err != nil {
		return nil, repoErrors.NewMappingError(err.Error())
	}

	*multimediaStorage = strings.ReplaceAll(*multimediaStorage, "$USER", document.Username)
	*multimediaStorage = *multimediaStorage + constants.SLASH + document.Username + "/img/" + document.Username + ".webp"
	result.PhotoUrl = *multimediaStorage

	return &result, nil
}
