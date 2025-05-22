package rest

import (
	errorInterceptor "com.demo.poc/commons/interceptor/errors"
	"com.demo.poc/commons/interceptor/restserver"
	props "com.demo.poc/commons/properties"
	"github.com/gin-gonic/gin"
)

func NewRouter(
	engine *gin.Engine,
	errorInterceptor *errorInterceptor.ErrorInterceptor,
	restService *RepoFinderRestService) *gin.Engine {

	engine.Use(gin.Recovery(), gin.Logger(), errorInterceptor.InterceptError(), restserver.InterceptRestServer(&props.Properties))

	productRouter := engine.Group("/poc/repositories/v1")
	{
		productRouter.GET("/users/:owner/repos", restService.FindRepositoriesByOwner)
	}

	return engine
}
