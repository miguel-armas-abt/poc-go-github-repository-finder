package injection

import (
	gitHubRepository "com.demo.poc/cmd/products/repository/github"
	gitHubErrorExtractor "com.demo.poc/cmd/products/repository/github/error"
	"com.demo.poc/cmd/products/rest"
	"com.demo.poc/cmd/products/service"
	errorSelector "com.demo.poc/commons/errors/selector"
	errorInterceptor "com.demo.poc/commons/interceptor/errors"
	properties "com.demo.poc/commons/properties"
	restClientErrors "com.demo.poc/commons/restclient/errors"
	"com.demo.poc/commons/validations"

	"github.com/gin-gonic/gin"
)

func NewEngine() *gin.Engine {
	engine := gin.New()

	props := &properties.Properties

	restClientErrorSelector := errorSelector.NewRestClientErrorSelector(&properties.Properties)
	restClientErrorExtractors := []restClientErrors.RestClientErrorExtractor{
		restClientErrors.DefaultExtractor{},
		gitHubErrorExtractor.GitHubErrorExtractor{},
	}

	restClientErrorHandler := restClientErrors.NewRestCrestclientErrorHandler(restClientErrorSelector, restClientErrorExtractors)
	// dbConnection := coreConfig.NewDatabaseConnection()

	githubRepository := gitHubRepository.NewGitHubRepositoryImpl(props, &restClientErrorHandler)
	svc := service.NewRepoFinderServiceImpl(githubRepository)

	coreValidator := validations.GetValidator()
	paramValidator := validations.NewParamValidator(coreValidator)
	bodyValidator := validations.NewBodyValidator(coreValidator)

	responseErrorSelector := errorSelector.NewResponseErrorSelector(props)
	interceptor := errorInterceptor.NewErrorInterceptor(responseErrorSelector)

	restService := rest.NewRepoFinderRestService(svc, paramValidator, bodyValidator)
	rest.NewRouter(engine, interceptor, restService)

	return engine
}
