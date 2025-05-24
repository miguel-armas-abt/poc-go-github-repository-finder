package rest

import (
	"github.com/gin-gonic/gin"
)

func NewRouter(
	engine *gin.Engine,
	profileRestService *ProfileRestService,
) *gin.Engine {

	router := engine.Group("/poc/repositories/v1")
	{
		router.POST("/profiles", profileRestService.InsertProfile)
		router.GET("/profiles/:username", profileRestService.FindByUsername)
	}

	return engine
}
