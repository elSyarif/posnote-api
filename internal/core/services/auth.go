package services

import (
	"context"

	"github.com/elSyarif/posnote-api.git/internal/core/ports"
)

type authService struct {
	repository ports.AuthRepository
}

func NewAuthService(repository ports.AuthRepository) ports.AuthService {
	return &authService{
		repository: repository,
	}
}

func (service *authService) AddRefreshToken(ctx context.Context, token string) error {
	return service.repository.Save(ctx, token)
}

func (service *authService) VerifyRefreshToken(ctx context.Context, token string) error {
	return service.repository.VerifyRefreshToken(ctx, token)
}

func (service *authService) DeleteRefreshToken(ctx context.Context, token string) error {
	return service.repository.DeleteRefreshToken(ctx, token)
}
