package auth

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/elSyarif/posnote-api.git/internal/core/domain"
	"github.com/elSyarif/posnote-api.git/internal/core/ports"
	"github.com/elSyarif/posnote-api.git/internal/helper"
	"github.com/gin-gonic/gin"
)

type handler struct {
	employeeService ports.EmployeeService
	authService     ports.AuthService
}

func NewAuthHandler(employeeService ports.EmployeeService, authService ports.AuthService) *handler {
	return &handler{
		employeeService: employeeService,
		authService:     authService,
	}
}

func (handler *handler) PostAuthentication(ctx *gin.Context) {
	c, cancel := context.WithCancel(context.Background())
	defer cancel()

	var auth *domain.Auth
	err := ctx.ShouldBindJSON(&auth)
	if err != nil {
		ctx.Error(err)
		return
	}

	id, err := handler.employeeService.VerifyCredential(c, auth.Username, auth.Password)
	if err != nil {
		ctx.Error(err)
		return
	}

	accessToken, err := helper.GenerateAccessToken(id)
	if err != nil {
		ctx.Error(err)
		return
	}
	refreshToken, err := helper.GenerateRefreshToken(id)
	if err != nil {
		ctx.Error(err)
		return
	}

	err = handler.authService.AddRefreshToken(c, fmt.Sprintf("%s", refreshToken))
	if err != nil {
		ctx.Error(err)
		return
	}

	helper.HTTPResponseSuccessWithData(ctx, 200, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

func (handler *handler) PutAuthentication(ctx *gin.Context) {
	c, cancel := context.WithCancel(context.Background())
	defer cancel()

	var token *domain.RefreshToken

	err := ctx.ShouldBindJSON(&token)
	if err != nil {
		ctx.Error(err)
		return
	}

	err = handler.authService.VerifyRefreshToken(c, token.RefreshToken)
	if err != nil {
		ctx.Error(err)
		return
	}
	data, err := helper.GetJWTData(token.RefreshToken, os.Getenv("REFRESH_TOKEN_KEY"))
	if err != nil {
		ctx.Error(err)
		return
	}

	now := time.Now().Unix()
	if now > data.ExpiresAt {
		ctx.Error(errors.New("refresh token sudah kadaluarsa"))
		return
	}

	accessToken, err := helper.GenerateAccessToken(data.EmpId.String())
	if err != nil {
		ctx.Error(err)
		return
	}

	helper.HTTPResponseSuccessWithData(ctx, 200, gin.H{
		"accessToken": accessToken,
	})
}

func (handler *handler) DeleteAuthentication(ctx *gin.Context) {
	c, cancel := context.WithCancel(context.Background())
	defer cancel()

	var token *domain.RefreshToken

	err := ctx.ShouldBindJSON(&token)
	if err != nil {
		ctx.Error(err)
		return
	}

	err = handler.authService.VerifyRefreshToken(c, token.RefreshToken)
	if err != nil {
		ctx.Error(err)
		return
	}
	_, err = helper.GetJWTData(token.RefreshToken, os.Getenv("REFRESH_TOKEN_KEY"))
	if err != nil {
		ctx.Error(err)
		return
	}

	err = handler.authService.DeleteRefreshToken(ctx, token.RefreshToken)
	if err != nil {
		ctx.Error(err)
		return
	}

	helper.HTTPResponseSuccess(ctx, 200, "refresh token berhasil dihapus")
}
