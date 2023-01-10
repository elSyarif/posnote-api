package api

import (
	"net/http"

	"github.com/elSyarif/posnote-api.git/internal/api/auth"
	"github.com/elSyarif/posnote-api.git/internal/api/employees"
	"github.com/elSyarif/posnote-api.git/internal/api/plants"
	"github.com/elSyarif/posnote-api.git/internal/api/roles"
	"github.com/elSyarif/posnote-api.git/internal/config"
	"github.com/elSyarif/posnote-api.git/internal/helper"
	"github.com/gin-gonic/gin"
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
	auth.NewAuthRoutes(v1, db)
	plants.NewPlantRoutes(v1, db)

	app.NoRoute(func(ctx *gin.Context) {
		helper.HTTPResponseError(ctx, http.StatusNotFound, "NOT_FOUND", "page not found", nil)
	})
	return app
}
