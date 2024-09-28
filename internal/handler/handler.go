package handler

import (
	"buildings/internal/group"

	"github.com/gin-gonic/gin"
)

// @title           Swagger API
// @version         1.0
// @description     This is an API server for working with buildings
func InitRoutes(router *gin.Engine, groups group.Groups) {
	api := router.Group("/api/v1")
	{
		buildings := api.Group("/buildings")
		{
			buildings.GET("/", groups.Building.Get)
			buildings.POST("/", groups.Save)
		}
	}
}
