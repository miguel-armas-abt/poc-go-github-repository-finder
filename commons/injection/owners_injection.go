package injection

import (
	profileRepository "com.demo.poc/cmd/profile/repository/profile"
	profileRest "com.demo.poc/cmd/profile/rest"
	profileService "com.demo.poc/cmd/profile/service"

	properties "com.demo.poc/commons/properties"
	"com.demo.poc/commons/validations"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func InjectProfileConfig(
	engine *gin.Engine,
	props *properties.ApplicationProperties,
	paramValidator *validations.ParamValidator,
	bodyValidator *validations.BodyValidator,
	mongoInstance *mongo.Database,
) *gin.Engine {

	ownRepository := profileRepository.NewProfileRepositoryImpl(mongoInstance)
	ownService := profileService.NewProfileServiceImpl(ownRepository, *props)
	ownRestService := profileRest.NewProfileRestService(ownService, paramValidator, bodyValidator)
	return profileRest.NewRouter(engine, ownRestService)
}
