package injection

import (
	//parameters command
	parameterRepository "com.demo.poc/cmd/parameters/repository/parameters"
	paramtersRest "com.demo.poc/cmd/parameters/rest"
	parametersService "com.demo.poc/cmd/parameters/service"

	//repo finder
	gitHubRepository "com.demo.poc/cmd/products/repository/github"
	gitHubErrorExtractor "com.demo.poc/cmd/products/repository/github/error"
	repoRest "com.demo.poc/cmd/products/rest"
	repoService "com.demo.poc/cmd/products/service"

	//commons
	coreConfig "com.demo.poc/commons/config"
	errorSelector "com.demo.poc/commons/errors/selector"
	errorInterceptor "com.demo.poc/commons/interceptor/errors"
	properties "com.demo.poc/commons/properties"
	restClientErrors "com.demo.poc/commons/restclient/errors"

	"com.demo.poc/commons/validations"

	"github.com/gin-gonic/gin"
)

func NewEngine() *gin.Engine {
	//commons
	engine := gin.New()

	props := &properties.Properties

	restClientErrorSelector := errorSelector.NewRestClientErrorSelector(&properties.Properties)
	restClientErrorExtractors := []restClientErrors.RestClientErrorExtractor{
		restClientErrors.DefaultExtractor{},
		gitHubErrorExtractor.GitHubErrorExtractor{},
	}
	restClientErrorHandler := restClientErrors.NewRestCrestclientErrorHandler(restClientErrorSelector, restClientErrorExtractors)

	coreValidator := validations.GetValidator()
	paramValidator := validations.NewParamValidator(coreValidator)
	bodyValidator := validations.NewBodyValidator(coreValidator)

	responseErrorSelector := errorSelector.NewResponseErrorSelector(props)
	interceptor := errorInterceptor.NewErrorInterceptor(responseErrorSelector)

	//parameters command
	mongoClient := coreConfig.NewMongoConnection(props)
	mongoInstance := mongoClient.Database(props.Mongo.Database)
	parameterRepository := parameterRepository.NewRepoParameterRepositoryImpl(mongoInstance)
	parameterCommandService := parametersService.NewRepoParameterCommandServiceImpl(parameterRepository)
	paramtersRestService := paramtersRest.NewParametersCommandRestService(parameterCommandService, paramValidator, bodyValidator)
	paramtersRest.NewRouter(engine, interceptor, paramtersRestService)

	//repo finder
	githubRepository := gitHubRepository.NewGitHubRepositoryImpl(props, &restClientErrorHandler)
	repoFinderService := repoService.NewRepoFinderServiceImpl(githubRepository)
	repoRestService := repoRest.NewRepoFinderRestService(repoFinderService, paramValidator, bodyValidator)
	repoRest.NewRouter(engine, interceptor, repoRestService)

	return engine
}
