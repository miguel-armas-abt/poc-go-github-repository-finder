package injection

import (
	"net/http"
	"time"

	gitHubErrorExtractor "com.demo.poc/cmd/repos/repository/github/error"
	coreConfig "com.demo.poc/commons/config"

	errorSelector "com.demo.poc/commons/errors/selector"
	errorInterceptor "com.demo.poc/commons/interceptor/errors"
	"com.demo.poc/commons/interceptor/restclient"
	"com.demo.poc/commons/interceptor/restserver"
	"com.demo.poc/commons/logging"
	properties "com.demo.poc/commons/properties"
	restClientErrors "com.demo.poc/commons/restclient/errors"

	"com.demo.poc/commons/validations"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func NewEngine(yamlBytes []byte) *gin.Engine {
	logging.InitLogger(logrus.InfoLevel)
	properties.Init(yamlBytes)
	http.DefaultClient.Transport = restclient.NewRestClientInterceptor(http.DefaultTransport, &properties.Properties)

	props := &properties.Properties

	responseErrorSelector := errorSelector.NewResponseErrorSelector(props)
	interceptor := errorInterceptor.NewErrorInterceptor(responseErrorSelector)

	engine := gin.New()
	engine.Use(
		gin.Recovery(),
		gin.Logger(),
		interceptor.InterceptError(),
		restserver.InterceptRestServer(props),
		cors.New(cors.Config{
			AllowOrigins:     []string{"http://localhost:63342", "https://repository-finder.vercel.app"},
			AllowMethods:     []string{"GET", "POST", "OPTIONS"},
			AllowHeaders:     []string{"Content-Type", "traceParent", "channelId"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}),
	)

	restClientErrorSelector := errorSelector.NewRestClientErrorSelector(&properties.Properties)
	restClientErrorExtractors := []restClientErrors.RestClientErrorExtractor{
		restClientErrors.DefaultExtractor{},
		gitHubErrorExtractor.GitHubErrorExtractor{},
	}
	restClientErrorHandler := restClientErrors.NewRestCrestclientErrorHandler(restClientErrorSelector, restClientErrorExtractors)

	coreValidator := validations.GetValidator()
	paramValidator := validations.NewParamValidator(coreValidator, responseErrorSelector)
	bodyValidator := validations.NewBodyValidator(coreValidator)

	mongoClient := coreConfig.NewMongoConnection(props)
	mongoInstance := mongoClient.Database(props.Mongo.Database)

	InjectReposConfig(engine, props, paramValidator, bodyValidator, &restClientErrorHandler, mongoInstance)
	InjectProfileConfig(engine, props, paramValidator, bodyValidator, mongoInstance)

	return engine
}
