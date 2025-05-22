package rest

import (
	"net/http"

	"com.demo.poc/cmd/parameters/dto/request"
	"com.demo.poc/cmd/parameters/service"
	utils "com.demo.poc/commons/restserver/utils"
	"com.demo.poc/commons/validations"
	headers "com.demo.poc/commons/validations/headers"
	"github.com/gin-gonic/gin"
)

type ParamtersCommandRestService struct {
	service        service.RepoParameterCommandService
	paramValidator *validations.ParamValidator
	bodyValidator  *validations.BodyValidator
}

func NewParametersCommandRestService(
	service service.RepoParameterCommandService,
	paramValidator *validations.ParamValidator,
	bodyValidator *validations.BodyValidator,
) *ParamtersCommandRestService {

	return &ParamtersCommandRestService{
		service:        service,
		paramValidator: paramValidator,
		bodyValidator:  bodyValidator,
	}
}

func (rest *ParamtersCommandRestService) InsertRepoParameter(ctx *gin.Context) {
	var defaultHeaders headers.DefaultHeaders
	if !rest.paramValidator.ValidateParamAndBind(ctx, &defaultHeaders) {
		return
	}

	insertRequest, ok := validations.ValidateBodyAndGet[request.RepoParameterInsertRequest](ctx, rest.bodyValidator)
	if !ok {
		return
	}

	err := rest.service.InsertRepoParameter(ctx.Request.Context(), utils.ExtractHeadersAsMap(ctx.Request.Header), insertRequest)

	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}

	ctx.Status(http.StatusAccepted)
}
