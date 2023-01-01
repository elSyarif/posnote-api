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
	c, cancel := context.WithCancel(context.Background())
	defer cancel()

	var role *domain.Roles
	err := ctx.ShouldBindJSON(&role)
	if err != nil {
		ctx.Error(err)
		return
	}

	result, err := handler.service.AddRole(c, role)
	if err != nil {
		ctx.Error(err)
		return
	}
	helper.HTTPResponseSuccessWithData(ctx, http.StatusCreated, gin.H{"roleId": result.Id})
}

func (handler *Handler) GetRoles(ctx *gin.Context) {
	c, cancel := context.WithCancel(context.Background())
	defer cancel()

	name := ctx.Query("name")
	role, err := handler.service.GetRole(c, name)
	if err != nil {
		helper.HTTPResponseError(ctx, http.StatusBadRequest, "fail", "gagal ambil data role", err.Error())
		return
	}

	helper.HTTPResponseSuccessWithData(ctx, 200, gin.H{
		"roles": role,
	})
}

func (handler *Handler) EditRole(ctx *gin.Context) {
	c, cancel := context.WithCancel(context.Background())
	defer cancel()
	var role *domain.Roles
	roleId := ctx.Param("id")

	err := ctx.ShouldBindJSON(&role)
	if err != nil {
		ctx.Error(err)
		return
	}

	err = handler.service.EditRole(c, roleId, role)
	if err != nil {
		helper.HTTPResponseError(ctx, 400, "fail", "", err.Error())
		return
	}

	helper.HTTPResponseSuccess(ctx, 200, "Role berhasil diperbaharui")
}

func (handler *Handler) Delete(ctx *gin.Context) {
	c, cancel := context.WithCancel(context.Background())
	defer cancel()
	roleId := ctx.Param("id")

	err := handler.service.Delete(c, roleId)
	if err != nil {
		helper.HTTPResponseError(ctx, 400, "fail", "", err.Error())
		return
	}

	helper.HTTPResponseSuccess(ctx, 200, "Role berhasil dihapus")
}
