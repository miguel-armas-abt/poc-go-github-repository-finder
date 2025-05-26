package injection

import (
	repoHelper "poc/cmd/repos/helper"
	gitHubRepository "poc/cmd/repos/repository/github"
	metadataRepository "poc/cmd/repos/repository/metadata"
	repoRest "poc/cmd/repos/rest"
	repoService "poc/cmd/repos/service"

	"go.mongodb.org/mongo-driver/mongo"

	restClientErrors "poc/commons/core/restclient/errors"
	validations "poc/commons/core/validations"

	"github.com/gin-gonic/gin"
)

func InjectReposConfig(
	engine *gin.Engine,
	paramValidator *validations.ParamValidator,
	bodyValidator *validations.BodyValidator,
	restClientErrorHandler *restClientErrors.RestclientErrorHandler,
	mongoInstance *mongo.Database,
) *gin.Engine {

	repoMetadataRepository := metadataRepository.NewRepoMetadataRepositoryImpl(mongoInstance)
	repoMetadataService := repoService.NewRepoMetadataServiceImpl(repoMetadataRepository)
	repoMetadataRestService := repoRest.NewRepoMetadataRestService(repoMetadataService, paramValidator, bodyValidator)

	githubRepository := gitHubRepository.NewGitHubRepositoryImpl(restClientErrorHandler)
	repoMergeHelper := repoHelper.NewRepoMergeHelper(githubRepository, repoMetadataRepository)
	repoFinderService := repoService.NewRepoFinderServiceImpl(repoMergeHelper)
	repoRestService := repoRest.NewRepoFinderRestService(repoFinderService, paramValidator, bodyValidator)
	return repoRest.NewRouter(engine, repoRestService, repoMetadataRestService)
}
