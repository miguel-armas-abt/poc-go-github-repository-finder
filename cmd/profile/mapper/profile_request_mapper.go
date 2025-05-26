package mapper

import (
	"strings"

	requestDto "com.demo.poc/cmd/profile/dto/request"
	document "com.demo.poc/cmd/profile/repository/profile/document"
	"com.demo.poc/commons/core/constants"
	repoErrors "com.demo.poc/commons/core/errors/errors"
	"github.com/mitchellh/mapstructure"
)

func ToDocument(
	request requestDto.ProfileInsertRequest,
	multimediaStorage *string,
	gitHubDomain *string,
) (*document.ProfileDocument, error) {

	var result document.ProfileDocument

	if err := mapstructure.Decode(request, &result); err != nil {
		return nil, repoErrors.NewMappingError(err.Error())
	}

	*multimediaStorage = strings.ReplaceAll(*multimediaStorage, "$USER", request.Username)
	*multimediaStorage = *multimediaStorage + constants.SLASH + request.Username + "/pdf/" + request.CvName + ".pdf"
	result.CvUrl = *multimediaStorage

	result.GitHubUrl = *gitHubDomain + constants.SLASH + request.Username

	return &result, nil
}
