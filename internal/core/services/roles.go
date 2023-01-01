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

func (service *roleService) GetRole(ctx context.Context, name string) (*[]domain.Roles, error) {
	return service.repository.FindRole(ctx, name)
}

func (service *roleService) EditRole(ctx context.Context, roleId string, role *domain.Roles) error {
	return service.repository.Update(ctx, roleId, role)
}

func (service *roleService) Delete(ctx context.Context, roleId string) error {
	return service.repository.Delete(ctx, roleId)
}
