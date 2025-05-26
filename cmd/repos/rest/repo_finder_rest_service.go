package rest

import (
	"net/http"

	params "poc/cmd/repos/params"
	"poc/cmd/repos/service"
	utils "poc/commons/core/restserver/utils"
	"poc/commons/core/validations"
	headers "poc/commons/core/validations/headers"

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

func (rest *RepoFinderRestService) FindRepositoriesByProfile(ctx *gin.Context) {
	var defaultHeaders headers.DefaultHeaders
	if !rest.paramValidator.ValidateParamAndBind(ctx, &defaultHeaders) {
		return
	}

	profile := ctx.Param("profile")
	label := ctx.Query("label")

	var inputParams params.RepoFinderParams
	inputParams.Profile = profile
	inputParams.Label = label

	if !rest.paramValidator.ValidateParamAndBind(ctx, &inputParams) {
		return
	}

	repoList, err := rest.service.FindRepositoriesByProfileAndLabel(ctx.Request.Context(), utils.ExtractHeadersAsMap(ctx.Request.Header), &inputParams)

	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, repoList)
}
