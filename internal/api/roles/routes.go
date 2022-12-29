package api

import (
	"github.com/elSyarif/posnote-api.git/internal/core/services"
	"github.com/elSyarif/posnote-api.git/internal/repositories/mysql_db"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func NewRolesRoute(router *gin.RouterGroup, db *sqlx.DB) *gin.RouterGroup {
	repository := mysql_db.NewRolesRepository(db)
	services := services.NewRolesService(repository)
	handler := NewRolesHandler(services)

	router.POST("roles", handler.AddRole)

	return router
}
