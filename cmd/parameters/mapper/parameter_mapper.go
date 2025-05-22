package mapper

import (
	requestDto "com.demo.poc/cmd/parameters/dto/request"
	"com.demo.poc/cmd/parameters/repository/parameters/document"
	"github.com/mitchellh/mapstructure"
)

func ToDocument(request requestDto.ParameterInsertRequest) document.ParameterDocument {
	var result document.ParameterDocument
	mapstructure.Decode(request, &result)
	return result
}
