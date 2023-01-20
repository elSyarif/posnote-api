package services

import (
	"context"

	"github.com/elSyarif/posnote-api.git/internal/core/domain"
	"github.com/elSyarif/posnote-api.git/internal/core/ports"
)

type employeePlantService struct {
	repository ports.EmployeePlantRepostory
}

func NewEmployeePlantService(repository ports.EmployeePlantRepostory) ports.EmployeePlantService {
	return &employeePlantService{
		repository: repository,
	}
}

func (service *employeePlantService) AddEmployeePlant(ctx context.Context, emplplant *domain.EmployeePlants) (*string, error) {
	return service.repository.Save(ctx, emplplant)
}
func (service *employeePlantService) GetEmloyeePlant(ctx context.Context, params string) (*[]domain.EmplPlantResponse, error) {
	return service.repository.Find(ctx, params)
}
func (service *employeePlantService) GetByEmployeeId(ctx context.Context, employeeId string) (*domain.EmplPlantResponse, error) {
	return service.repository.FindByEmployeeId(ctx, employeeId)
}
func (service *employeePlantService) GetByPlantId(ctx context.Context, plantId string) (*[]domain.EmplPlantResponse, error) {
	return service.repository.FindByPlantId(ctx, plantId)
}
func (service *employeePlantService) Update(ctx context.Context, plantId string, emplId string, emplplant *domain.EmployeePlants) error {
	return service.repository.Update(ctx, plantId, emplId, emplplant)
}
func (service *employeePlantService) Delete(ctx context.Context, plantId string, emplId string) error {
	return service.repository.Delete(ctx, plantId, emplId)
}
