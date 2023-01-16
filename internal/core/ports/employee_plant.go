package ports

import (
	"context"

	"github.com/elSyarif/posnote-api.git/internal/core/domain"
)

type EmployeePlantRepostory interface {
	Save(ctx context.Context, emplplant *domain.EmployeePlants) (*string, error)
	Find(ctx context.Context, params string) (*[]domain.EmplPlantResponse, error)
	FindByEmployeeId(ctx context.Context, employeeId string) (*domain.EmplPlantResponse, error)
	FindByPlantId(ctx context.Context, plantId string) (*[]domain.EmplPlantResponse, error)
	Update(ctx context.Context, plantId string, emplId string, emplplant *domain.EmployeePlants) error
	Delete(ctx context.Context, plantId string, emplId string) error
}

type EmployeePlantService interface {
	AddEmployeePlant(ctx context.Context, emplplant *domain.EmployeePlants) (*string, error)
	GetEmloyeePlant(ctx context.Context, params string) (*[]domain.EmplPlantResponse, error)
	GetByEmployeeId(ctx context.Context, employeeId string) (*domain.EmplPlantResponse, error)
	GetByPlantId(ctx context.Context, plantId string) (*[]domain.EmplPlantResponse, error)
	Update(ctx context.Context, plantId string, emplId string, emplplant *domain.EmployeePlants) error
	Delete(ctx context.Context, plantId string, emplId string) error
}
