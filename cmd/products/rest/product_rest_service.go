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

func (rest *RepoFinderRestService) FindRepositoriesByOwner(context *gin.Context) {
	var defaultHeaders headers.DefaultHeaders
	if !rest.paramValidator.ValidateParamAndBind(context, &defaultHeaders) {
		return
	}
	headers := utils.ExtractHeadersAsMap(context.Request.Header)

	owner := context.Param("owner")

	if owner == constants.EMPTY || &owner == nil {
		context.Error(coreError.NewInvalidFieldError("owner must not be empty"))
		return
	}

	repoList, err := rest.service.FindRepositoriesByOwner(headers, owner)

	if err != nil {
		context.Error(err)
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, repoList)
}
