package roles

import (
	"context"
	"net/http"

	"github.com/elSyarif/posnote-api.git/internal/core/domain"
	"github.com/elSyarif/posnote-api.git/internal/core/ports"
	"github.com/gin-gonic/gin"
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

	ctx.JSON(http.StatusCreated, gin.H{
		"data": role,
	})
}
