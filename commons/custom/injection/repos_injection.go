package injection

import (
	repoHelper "com.demo.poc/cmd/repos/helper"
	gitHubRepository "com.demo.poc/cmd/repos/repository/github"
	metadataRepository "com.demo.poc/cmd/repos/repository/metadata"
	repoRest "com.demo.poc/cmd/repos/rest"
	repoService "com.demo.poc/cmd/repos/service"
	"go.mongodb.org/mongo-driver/mongo"

	restClientErrors "com.demo.poc/commons/core/restclient/errors"
	validations "com.demo.poc/commons/core/validations"

	"com.demo.poc/commons/custom/properties"
	"github.com/gin-gonic/gin"
)

func InjectReposConfig(
	engine *gin.Engine,
	props *properties.ApplicationProperties,
	paramValidator *validations.ParamValidator,
	bodyValidator *validations.BodyValidator,
	restClientErrorHandler *restClientErrors.RestclientErrorHandler,
	mongoInstance *mongo.Database,
) *gin.Engine {

	repoMetadataRepository := metadataRepository.NewRepoMetadataRepositoryImpl(mongoInstance)
	repoMetadataService := repoService.NewRepoMetadataServiceImpl(repoMetadataRepository, *props)
	repoMetadataRestService := repoRest.NewRepoMetadataRestService(repoMetadataService, paramValidator, bodyValidator)

	githubRepository := gitHubRepository.NewGitHubRepositoryImpl(props, restClientErrorHandler)
	repoMergeHelper := repoHelper.NewRepoMergeHelper(githubRepository, repoMetadataRepository)
	repoFinderService := repoService.NewRepoFinderServiceImpl(repoMergeHelper)
	repoRestService := repoRest.NewRepoFinderRestService(repoFinderService, paramValidator, bodyValidator)
	return repoRest.NewRouter(engine, repoRestService, repoMetadataRestService)
}
