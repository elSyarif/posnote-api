package roles

import (
	"github.com/elSyarif/posnote-api.git/internal/core/services"
	"github.com/elSyarif/posnote-api.git/internal/repositories/mysql_db"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func NewRolesRoute(router *gin.RouterGroup, db *sqlx.DB) *gin.RouterGroup {
	repository := mysql_db.NewRolesRepository(db)
	service := services.NewRolesService(repository)
	handler := NewRolesHandler(service)

	router.POST("roles", handler.AddRole)
	router.GET("roles", handler.GetRoles)
	router.PUT("roles/:id", handler.EditRole)
	router.DELETE("roles/:id", handler.Delete)

	return router
}
