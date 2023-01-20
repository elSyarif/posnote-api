package employeeplant

import (
	"github.com/elSyarif/posnote-api.git/internal/core/services"
	"github.com/elSyarif/posnote-api.git/internal/repositories/mysql_db"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func NewRoutesEmployeePlant(router *gin.RouterGroup, db *sqlx.DB) *gin.RouterGroup {
	repository := mysql_db.NewEmployeePlantRepostory(db)
	service := services.NewEmployeePlantService(repository)
	handler := NewHandlerEmployeePlant(service)

	plant := router.Group("/plants")

	plant.POST("/:plantId/employee", handler.PostEmployeePlant)
	plant.GET("/:plantId/employee", handler.GetEmployeePlantByPlantId)
	plant.GET("/:plantId/employee/:emplId", handler.GetEmployeePlantByEmployeeId)
	plant.PUT("/:plantId/employee/:emplId", handler.PutEmployeePlant)
	plant.DELETE("/:plantId/employee/:emplId", handler.DeleteEmployeePlant)

	return plant
}
