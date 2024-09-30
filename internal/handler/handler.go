package handler

import (
	"buildings/internal/group"

	_ "buildings/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Building API
// @version         1.0
// @description     This is an API server for working with buildings
//
//	@host			localhost:8080
func InitRoutes(router *gin.Engine, groups group.Groups) {
	api := router.Group("/api/v1")
	{
		buildings := api.Group("/buildings")
		{
			buildings.GET("/", groups.Building.Get)
			buildings.POST("/", groups.Building.Save)
		}
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
