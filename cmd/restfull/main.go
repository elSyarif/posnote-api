package main

import (
	"github.com/elSyarif/posnote-api.git/internal/api"
	"github.com/elSyarif/posnote-api.git/internal/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.New()
	app.Use(gin.Recovery())
	app.Use(middleware.ErrorHandler)

	serve := api.NewApiServer(app)

	serve.Run("localhost:5000")
}
