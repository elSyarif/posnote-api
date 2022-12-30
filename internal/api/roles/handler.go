package roles

import (
	"context"
	"github.com/elSyarif/posnote-api.git/internal/core/domain"
	"github.com/elSyarif/posnote-api.git/internal/core/ports"
	"github.com/elSyarif/posnote-api.git/internal/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	service ports.RoleService
}

func NewRolesHandler(rolesService ports.RoleService) *Handler {
	return &Handler{
		service: rolesService,
	}
}

func (handler *Handler) AddRole(ctx *gin.Context) {
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	var role *domain.Roles
	err := ctx.ShouldBindJSON(&role)
	if err != nil {
		ctx.Error(err)
		return
	}

	result, err := handler.service.AddRole(ctx, role)
	if err != nil {
		ctx.Error(err)
		return
	}
	helper.HTTPResponseSuccessWithData(ctx, http.StatusCreated, gin.H{"roleId": result.Id})
}
