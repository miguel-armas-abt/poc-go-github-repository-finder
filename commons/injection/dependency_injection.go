package injection

import (
	//parameters command
	"net/http"

	parameterRepository "com.demo.poc/cmd/parameters/repository/parameters"
	paramtersRest "com.demo.poc/cmd/parameters/rest"
	parametersService "com.demo.poc/cmd/parameters/service"

	//repo finder
	repoHelper "com.demo.poc/cmd/repos/helper"
	gitHubRepository "com.demo.poc/cmd/repos/repository/github"
	gitHubErrorExtractor "com.demo.poc/cmd/repos/repository/github/error"
	repoRest "com.demo.poc/cmd/repos/rest"
	repoService "com.demo.poc/cmd/repos/service"

	//commons
	coreConfig "com.demo.poc/commons/config"
	errorSelector "com.demo.poc/commons/errors/selector"
	errorInterceptor "com.demo.poc/commons/interceptor/errors"
	"com.demo.poc/commons/interceptor/restclient"
	"com.demo.poc/commons/logging"
	properties "com.demo.poc/commons/properties"
	restClientErrors "com.demo.poc/commons/restclient/errors"

	"com.demo.poc/commons/validations"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func NewEngine(yamlBytes []byte) *gin.Engine {
	//commons
	logging.InitLogger(logrus.InfoLevel)
	properties.Init(yamlBytes)
	http.DefaultClient.Transport = restclient.NewRestClientInterceptor(http.DefaultTransport, &properties.Properties)

	engine := gin.New()

	props := &properties.Properties

	responseErrorSelector := errorSelector.NewResponseErrorSelector(props)
	interceptor := errorInterceptor.NewErrorInterceptor(responseErrorSelector)

	restClientErrorSelector := errorSelector.NewRestClientErrorSelector(&properties.Properties)
	restClientErrorExtractors := []restClientErrors.RestClientErrorExtractor{
		restClientErrors.DefaultExtractor{},
		gitHubErrorExtractor.GitHubErrorExtractor{},
	}
	restClientErrorHandler := restClientErrors.NewRestCrestclientErrorHandler(restClientErrorSelector, restClientErrorExtractors)

	coreValidator := validations.GetValidator()
	paramValidator := validations.NewParamValidator(coreValidator, responseErrorSelector)
	bodyValidator := validations.NewBodyValidator(coreValidator)

	//parameters command
	mongoClient := coreConfig.NewMongoConnection(props)
	mongoInstance := mongoClient.Database(props.Mongo.Database)
	parameterRepository := parameterRepository.NewParameterRepositoryImpl(mongoInstance)
	parameterCommandService := parametersService.NewParameterServiceImpl(parameterRepository, *props)
	paramtersRestService := paramtersRest.NewParameterRestService(parameterCommandService, paramValidator, bodyValidator)
	paramtersRest.NewRouter(engine, interceptor, paramtersRestService)

	//repo finder
	githubRepository := gitHubRepository.NewGitHubRepositoryImpl(props, &restClientErrorHandler)
	repoMergeHelper := repoHelper.NewRepoMergeHelper(githubRepository, parameterRepository)
	repoFinderService := repoService.NewRepoFinderServiceImpl(repoMergeHelper)
	repoRestService := repoRest.NewRepoFinderRestService(repoFinderService, paramValidator, bodyValidator)
	repoRest.NewRouter(engine, interceptor, repoRestService)

	return engine
}
