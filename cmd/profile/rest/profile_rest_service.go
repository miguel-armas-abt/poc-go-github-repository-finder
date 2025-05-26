package rest

import (
	"net/http"

	"com.demo.poc/cmd/profile/dto/request"
	"com.demo.poc/cmd/profile/service"
	"com.demo.poc/commons/core/constants"
	coreErrors "com.demo.poc/commons/core/errors/errors"
	utils "com.demo.poc/commons/core/restserver/utils"
	"com.demo.poc/commons/core/validations"
	headers "com.demo.poc/commons/core/validations/headers"
	"github.com/gin-gonic/gin"
)

type ProfileRestService struct {
	service        service.ProfileService
	paramValidator *validations.ParamValidator
	bodyValidator  *validations.BodyValidator
}

func NewProfileRestService(
	service service.ProfileService,
	paramValidator *validations.ParamValidator,
	bodyValidator *validations.BodyValidator,
) *ProfileRestService {

	return &ProfileRestService{
		service:        service,
		paramValidator: paramValidator,
		bodyValidator:  bodyValidator,
	}
}

func (rest *ProfileRestService) InsertProfile(ctx *gin.Context) {
	var defaultHeaders headers.DefaultHeaders
	if !rest.paramValidator.ValidateParamAndBind(ctx, &defaultHeaders) {
		return
	}

	insertRequest, ok := validations.ValidateBodyAndGet[request.ProfileInsertRequest](ctx, rest.bodyValidator)
	if !ok {
		return
	}

	err := rest.service.InsertProfile(ctx.Request.Context(), utils.ExtractHeadersAsMap(ctx.Request.Header), insertRequest)

	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}

	ctx.Status(http.StatusCreated)
}

func (rest *ProfileRestService) FindByUsername(ctx *gin.Context) {
	var defaultHeaders headers.DefaultHeaders
	if !rest.paramValidator.ValidateParamAndBind(ctx, &defaultHeaders) {
		return
	}

	username := ctx.Param("username")
	if username == constants.EMPTY {
		ctx.Error(coreErrors.NewInvalidFieldError("profile must not be empty"))
		ctx.Abort()
		return
	}

	profile, err := rest.service.FindByUsername(ctx.Request.Context(), utils.ExtractHeadersAsMap(ctx.Request.Header), username)

	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, profile)
}
