package roles

import (
	"github.com/elSyarif/posnote-api.git/internal/core/services"
	"github.com/elSyarif/posnote-api.git/internal/middleware"
	"github.com/elSyarif/posnote-api.git/internal/repositories/mysql_db"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func NewRolesRoute(router *gin.RouterGroup, db *sqlx.DB) *gin.RouterGroup {
	repository := mysql_db.NewRolesRepository(db)
	service := services.NewRolesService(repository)
	handler := NewRolesHandler(service)

	roles := router.Group("roles")
	roles.POST("", handler.AddRole)

	roles.Use(middleware.Auth())
	{
		roles.GET("", handler.GetRoles)
		roles.PUT("/:id", handler.EditRole)
		roles.DELETE("/:id", handler.Delete)
	}

	return roles
}
