package mapper

import (
	requestDto "com.demo.poc/cmd/parameters/dto/request"
	"com.demo.poc/cmd/parameters/repository/parameters/document"
	"github.com/mitchellh/mapstructure"
)

func ToDocument(request requestDto.RepoParameterInsertRequest) document.RepoParameterDocument {
	var result document.RepoParameterDocument
	mapstructure.Decode(request, &result)
	return result
}
