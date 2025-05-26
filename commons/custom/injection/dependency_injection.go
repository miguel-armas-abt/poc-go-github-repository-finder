package injection

import (
	"net/http"
	"time"

	gitHubErrorExtractor "poc/cmd/repos/repository/github/error"
	coreConfig "poc/commons/custom/config"

	errorSelector "poc/commons/core/errors/selector"
	errorInterceptor "poc/commons/core/interceptor/errors"
	"poc/commons/core/interceptor/restclient"
	"poc/commons/core/interceptor/restserver"
	"poc/commons/core/logging"
	restClientErrors "poc/commons/core/restclient/errors"
	"poc/commons/custom/properties"

	"poc/commons/core/validations"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func NewEngine(yamlBytes []byte) *gin.Engine {
	logging.InitLogger(logrus.InfoLevel)
	properties.Init(yamlBytes)
	http.DefaultClient.Transport = restclient.NewRestClientInterceptor(http.DefaultTransport, &properties.Properties)

	props := &properties.Properties

	responseErrorSelector := errorSelector.NewResponseErrorSelector()
	interceptor := errorInterceptor.NewErrorInterceptor(responseErrorSelector)

	corsOrigins := props.Server.CorsOrigins
	engine := gin.New()
	engine.Use(
		gin.Recovery(),
		gin.Logger(),
		interceptor.InterceptError(),
		restserver.InterceptRestServer(props),
		cors.New(cors.Config{
			AllowOrigins:     corsOrigins,
			AllowMethods:     []string{"GET", "POST", "OPTIONS"},
			AllowHeaders:     []string{"Content-Type", "traceParent", "channelId"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}),
	)

	restClientErrorSelector := errorSelector.NewRestClientErrorSelector()
	restClientErrorExtractors := []restClientErrors.RestClientErrorExtractor{
		restClientErrors.DefaultExtractor{},
		gitHubErrorExtractor.GitHubErrorExtractor{},
	}
	restClientErrorHandler := restClientErrors.NewRestCrestclientErrorHandler(restClientErrorSelector, restClientErrorExtractors)

	coreValidator := validations.GetValidator()
	paramValidator := validations.NewParamValidator(coreValidator, responseErrorSelector)
	bodyValidator := validations.NewBodyValidator(coreValidator)

	mongoClient := coreConfig.NewMongoConnection()
	mongoInstance := mongoClient.Database(props.Mongo.Database)

	InjectReposConfig(engine, paramValidator, bodyValidator, &restClientErrorHandler, mongoInstance)
	InjectProfileConfig(engine, paramValidator, bodyValidator, mongoInstance)

	return engine
}
