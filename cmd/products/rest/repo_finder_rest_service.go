package rest

import (
	"net/http"

	"com.demo.poc/cmd/products/service"
	"com.demo.poc/commons/constants"
	coreError "com.demo.poc/commons/errors/errors"
	utils "com.demo.poc/commons/restserver/utils"
	"com.demo.poc/commons/validations"
	headers "com.demo.poc/commons/validations/headers"
	"github.com/gin-gonic/gin"
)

type RepoFinderRestService struct {
	service        service.RepoFinderService
	paramValidator *validations.ParamValidator
	bodyValidator  *validations.BodyValidator
}

func NewRepoFinderRestService(
	service service.RepoFinderService,
	paramValidator *validations.ParamValidator,
	bodyValidator *validations.BodyValidator,
) *RepoFinderRestService {

	return &RepoFinderRestService{
		service:        service,
		paramValidator: paramValidator,
		bodyValidator:  bodyValidator,
	}
}

func (rest *RepoFinderRestService) FindRepositoriesByOwner(ctx *gin.Context) {
	var defaultHeaders headers.DefaultHeaders
	if !rest.paramValidator.ValidateParamAndBind(ctx, &defaultHeaders) {
		return
	}

	owner := ctx.Param("owner")

	if owner == constants.EMPTY || &owner == nil {
		ctx.Error(coreError.NewInvalidFieldError("owner must not be empty"))
		return
	}

	repoList, err := rest.service.FindRepositoriesByOwner(ctx.Request.Context(), utils.ExtractHeadersAsMap(ctx.Request.Header), owner)

	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, repoList)
}
