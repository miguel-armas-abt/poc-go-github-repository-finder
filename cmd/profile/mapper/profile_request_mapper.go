package mapper

import (
	"strings"

	requestDto "poc/cmd/profile/dto/request"
	document "poc/cmd/profile/repository/profile/document"
	"poc/commons/core/constants"
	coreErrors "poc/commons/core/errors/errors"

	"github.com/mitchellh/mapstructure"
)

func ToDocument(
	request requestDto.ProfileInsertRequest,
	multimediaStorage *string,
	gitHubDomain *string,
) (*document.ProfileDocument, error) {

	var result document.ProfileDocument

	if err := mapstructure.Decode(request, &result); err != nil {
		return nil, coreErrors.NewMappingError(err.Error())
	}

	*multimediaStorage = strings.ReplaceAll(*multimediaStorage, "$USER", request.Username)
	*multimediaStorage = *multimediaStorage + constants.SLASH + request.Username + "/pdf/" + request.CvName + ".pdf"
	result.CvUrl = *multimediaStorage

	result.GitHubUrl = *gitHubDomain + constants.SLASH + request.Username

	return &result, nil
}
