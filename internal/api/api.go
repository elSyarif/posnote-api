package api

import (
	"github.com/elSyarif/posnote-api.git/internal/api/employees"
	"github.com/elSyarif/posnote-api.git/internal/api/roles"
	"github.com/elSyarif/posnote-api.git/internal/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewApiServer(app *gin.Engine) *gin.Engine {
	env := config.New()
	db, err := config.NewDB(env)
	if err != nil {
		panic(err)
	}

	app.Group("v1")
	v1 := app.Group("v1")
	roles.NewRolesRoute(v1, db)
	employees.NewEmployeeRoutes(v1, db)

	app.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  "NOT_FOUND",
			"message": "Page not found",
		})
	})
	return app
}
