package ports

import (
	"context"

	"github.com/elSyarif/posnote-api.git/internal/core/domain"
)

type RoleRepository interface {
	Save(ctx context.Context, role *domain.Roles) (*domain.Roles, error)
}

type RoleService interface {
	AddRole(ctx context.Context, role *domain.Roles) (*domain.Roles, error)
}
