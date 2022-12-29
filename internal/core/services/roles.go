package services

import (
	"context"

	"github.com/elSyarif/posnote-api.git/internal/core/domain"
	"github.com/elSyarif/posnote-api.git/internal/core/ports"
)

type roleService struct {
	repository ports.RoleRepository
}

func NewRolesService(repository ports.RoleRepository) ports.RoleService {
	return &roleService{
		repository: repository,
	}
}

func (service *roleService) AddRole(ctx context.Context, role *domain.Roles) (*domain.Roles, error) {
	return service.repository.Save(ctx, role)
}
