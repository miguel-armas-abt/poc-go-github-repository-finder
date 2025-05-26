package injection

import (
	profileRepository "poc/cmd/profile/repository/profile"
	profileRest "poc/cmd/profile/rest"
	profileService "poc/cmd/profile/service"

	"poc/commons/core/validations"
	"poc/commons/custom/properties"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func InjectProfileConfig(
	engine *gin.Engine,
	paramValidator *validations.ParamValidator,
	bodyValidator *validations.BodyValidator,
	mongoInstance *mongo.Database,
) *gin.Engine {

	ownRepository := profileRepository.NewProfileRepositoryImpl(mongoInstance)
	ownService := profileService.NewProfileServiceImpl(ownRepository, properties.Properties)
	ownRestService := profileRest.NewProfileRestService(ownService, paramValidator, bodyValidator)
	return profileRest.NewRouter(engine, ownRestService)
}
