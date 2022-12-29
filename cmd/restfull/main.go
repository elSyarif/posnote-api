package main

import (
	api "github.com/elSyarif/posnote-api.git/internal/api/roles"
	"github.com/elSyarif/posnote-api.git/internal/config"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.New()
	env := config.New()
	db, err := config.NewDB(env)
	if err != nil {
		panic(err)
	}

	v1 := app.Group("v1")
	api.NewRolesRoute(v1, db)

	app.Run("localhost:5000")
}
