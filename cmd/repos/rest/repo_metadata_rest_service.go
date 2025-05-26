package rest

import (
	"net/http"

	"com.demo.poc/cmd/repos/dto/request"
	"com.demo.poc/cmd/repos/service"
	utils "com.demo.poc/commons/core/restserver/utils"
	"com.demo.poc/commons/core/validations"
	headers "com.demo.poc/commons/core/validations/headers"
	"github.com/gin-gonic/gin"
)

type RepoMetadataRestService struct {
	service        service.ParameterService
	paramValidator *validations.ParamValidator
	bodyValidator  *validations.BodyValidator
}

func NewRepoMetadataRestService(
	service service.ParameterService,
	paramValidator *validations.ParamValidator,
	bodyValidator *validations.BodyValidator,
) *RepoMetadataRestService {

	return &RepoMetadataRestService{
		service:        service,
		paramValidator: paramValidator,
		bodyValidator:  bodyValidator,
	}
}

func (rest *RepoMetadataRestService) InsertRepoMetadata(ctx *gin.Context) {
	var defaultHeaders headers.DefaultHeaders
	if !rest.paramValidator.ValidateParamAndBind(ctx, &defaultHeaders) {
		return
	}

	insertRequest, ok := validations.ValidateBodyAndGet[request.RepoMetadataInsertRequest](ctx, rest.bodyValidator)
	if !ok {
		return
	}

	err := rest.service.InsertRepoMetadata(ctx.Request.Context(), utils.ExtractHeadersAsMap(ctx.Request.Header), insertRequest)

	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}

	ctx.Status(http.StatusCreated)
}
