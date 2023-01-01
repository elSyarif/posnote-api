package ports

import (
	"context"

	"github.com/elSyarif/posnote-api.git/internal/core/domain"
)

type RoleRepository interface {
	Save(ctx context.Context, role *domain.Roles) (*domain.Roles, error)
	FindRole(ctx context.Context, name string) (*[]domain.Roles, error)
	Update(ctx context.Context, roleId string, role *domain.Roles) error
	Delete(ctx context.Context, roleId string) error
}

type RoleService interface {
	AddRole(ctx context.Context, role *domain.Roles) (*domain.Roles, error)
	GetRole(ctx context.Context, name string) (*[]domain.Roles, error)
	EditRole(ctx context.Context, roleId string, role *domain.Roles) error
	Delete(ctx context.Context, roleId string) error
}
