package ports

import (
	"context"

	"github.com/elSyarif/posnote-api.git/internal/core/domain"
)

type PlantsRepository interface {
	Save(ctx context.Context, plant *domain.Plants) (*string, error)
	Find(ctx context.Context, name string) (*[]domain.Plants, error)
	FindById(ctx context.Context, id string) (*domain.Plants, error)
	Update(ctx context.Context, id string, plant *domain.Plants) error
	Delete(ctx context.Context, id string) error
}

type PlantsService interface {
	AddPlants(ctx context.Context, plant *domain.Plants) (*string, error)
	GetPlants(ctx context.Context, name string) (*[]domain.Plants, error)
	GetById(ctx context.Context, id string) (*domain.Plants, error)
	Update(ctx context.Context, id string, plant *domain.Plants) error
	Delete(ctx context.Context, id string) error
}
