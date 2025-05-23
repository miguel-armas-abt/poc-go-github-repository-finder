package rest

import (
	"time"

	errorInterceptor "com.demo.poc/commons/interceptor/errors"
	"com.demo.poc/commons/interceptor/restserver"
	props "com.demo.poc/commons/properties"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter(
	engine *gin.Engine,
	errorInterceptor *errorInterceptor.ErrorInterceptor,
	restService *ParameterRestService) *gin.Engine {

	engine.Use(
		gin.Recovery(),
		gin.Logger(),
		errorInterceptor.InterceptError(),
		restserver.InterceptRestServer(&props.Properties),
		cors.New(cors.Config{
			AllowOrigins:     []string{"http://localhost:63342", "https://repository-finder.vercel.app"},
			AllowMethods:     []string{"GET", "POST", "OPTIONS"},
			AllowHeaders:     []string{"Content-Type", "traceParent", "channelId"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}),
	)

	productRouter := engine.Group("/poc/repositories/v1")
	{
		productRouter.POST("/parameters", restService.InsertRepoParameter)
	}

	return engine
}
