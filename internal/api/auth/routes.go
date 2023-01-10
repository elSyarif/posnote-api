package auth

import (
	"github.com/elSyarif/posnote-api.git/internal/core/services"
	"github.com/elSyarif/posnote-api.git/internal/repositories/mysql_db"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func NewAuthRoutes(router *gin.RouterGroup, db *sqlx.DB) *gin.RouterGroup {
	empRepository := mysql_db.NewEmployeeRepository(db)
	empService := services.NewEmployeeService(empRepository)

	authRepository := mysql_db.NewAuthRepository(db)
	authService := services.NewAuthService(authRepository)
	handler := NewAuthHandler(empService, authService)

	router.POST("/auth", handler.PostAuthentication)
	router.PUT("/auth", handler.PutAuthentication)
	router.DELETE("/auth", handler.DeleteAuthentication)

	return router
}
