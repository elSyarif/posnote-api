package services

import (
	"context"

	"github.com/elSyarif/posnote-api.git/internal/core/domain"
	"github.com/elSyarif/posnote-api.git/internal/core/ports"
)

type plantService struct {
	repository ports.PlantsRepository
}

func NewPlantService(repo ports.PlantsRepository) ports.PlantsService {
	return &plantService{
		repository: repo,
	}
}

func (service *plantService) AddPlants(ctx context.Context, plant *domain.Plants) (*string, error) {
	return service.repository.Save(ctx, plant)
}
func (service *plantService) GetPlants(ctx context.Context, name string) (*[]domain.Plants, error) {
	return service.repository.Find(ctx, name)
}
func (service *plantService) GetById(ctx context.Context, id string) (*domain.Plants, error) {
	return service.repository.FindById(ctx, id)
}
func (service *plantService) Update(ctx context.Context, id string, plant *domain.Plants) error {
	return service.repository.Update(ctx, id, plant)
}
func (service *plantService) Delete(ctx context.Context, id string) error {
	return service.repository.Delete(ctx, id)
}
