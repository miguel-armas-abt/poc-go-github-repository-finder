package rest

import (
	"github.com/gin-gonic/gin"
)

func NewRouter(
	engine *gin.Engine,
	repoFinderRestService *RepoFinderRestService,
	repoMetadataRestService *RepoMetadataRestService,
) *gin.Engine {

	router := engine.Group("/poc/repositories/v1")
	{
		router.GET("/users/:profile/repos", repoFinderRestService.FindRepositoriesByProfile)
		router.POST("/metadata", repoMetadataRestService.InsertRepoMetadata)
	}

	return engine
}
