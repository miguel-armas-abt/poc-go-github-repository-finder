package mapper

import (
	"com.demo.poc/cmd/profile/dto/response"
	"com.demo.poc/cmd/profile/repository/profile/document"
	repoErrors "com.demo.poc/commons/errors/errors"
	"github.com/mitchellh/mapstructure"
)

func ToResponse(document document.ProfileDocument) (*response.ProfileResponse, error) {
	var result response.ProfileResponse

	if err := mapstructure.Decode(document, &result); err != nil {
		return nil, repoErrors.NewMappingError(err.Error())
	}

	return &result, nil
}
