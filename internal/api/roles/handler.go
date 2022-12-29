package api

import (
	"context"
	"net/http"

	"github.com/elSyarif/posnote-api.git/internal/core/domain"
	"github.com/elSyarif/posnote-api.git/internal/core/ports"
	"github.com/gin-gonic/gin"
)

type RolesHandler struct {
	service ports.RoleService
}

func NewRolesHandler(rolesService ports.RoleService) *RolesHandler {
	return &RolesHandler{
		service: rolesService,
	}
}

func (handler *RolesHandler) AddRole(ctx *gin.Context) {
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	var role *domain.Roles
	ctx.ShouldBindJSON(&role)

	ctx.JSON(http.StatusCreated, gin.H{
		"data": role,
	})
}
