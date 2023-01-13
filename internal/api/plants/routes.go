package plants

import (
	"github.com/elSyarif/posnote-api.git/internal/core/services"
	"github.com/elSyarif/posnote-api.git/internal/middleware"
	"github.com/elSyarif/posnote-api.git/internal/repositories/mysql_db"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func NewPlantRoutes(router *gin.RouterGroup, db *sqlx.DB) *gin.RouterGroup {
	repository := mysql_db.NewPlantRepository(db)
	service := services.NewPlantService(repository)
	handler := NewPlantHandler(service)

	plants := router.Group("/plants")
	plants.Use(middleware.Auth())
	{
		plants.POST("", handler.AddPlant)
		plants.GET("", handler.GetPlant)
		plants.GET("/:plantId", handler.GetPlantById)
		plants.PUT("/:plantId", handler.Update)
		plants.DELETE("/:plantId", handler.Delete)
	}

	return plants
}
