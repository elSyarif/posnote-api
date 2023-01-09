package employees

import (
	"github.com/elSyarif/posnote-api.git/internal/core/services"
	"github.com/elSyarif/posnote-api.git/internal/middleware"
	"github.com/elSyarif/posnote-api.git/internal/repositories/mysql_db"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func NewEmployeeRoutes(router *gin.RouterGroup, db *sqlx.DB) *gin.RouterGroup {
	reposiroty := mysql_db.NewEmployeeRepository(db)
	service := services.NewEmployeeService(reposiroty)
	handler := NewEmployeeHandler(service)

	router.POST("/employees", handler.AddEmployee)
	router.GET("/employees/:id", middleware.Protected(), handler.GetEmployeeById)

	return router
}
