package ports

import (
	"context"

	"github.com/elSyarif/posnote-api.git/internal/core/domain"
)

type EmployeePlantRepostory interface {
	Save(ctx context.Context, emplplant *domain.EmployeePlants) (*string, error)
	Find(ctx context.Context, params ...string) (*[]domain.EmployeePlants, error)
	FindByEmployeeId(ctx context.Context, employeeId string) (*domain.EmployeePlants, error)
	FindByPlantId(ctx context.Context, plantId string) (*[]domain.EmployeePlants, error)
	Update(ctx context.Context, id string, emplplant *domain.EmployeePlants) error
	Delete(ctx context.Context, id string) error
}

type EmployeePlantService interface {
	AddEmployeePlant(ctx context.Context, emplplant *domain.EmployeePlants) (*string, error)
	GetEmloyeePlant(ctx context.Context, params ...string) (*[]domain.EmployeePlants, error)
	GetByEmployeeId(ctx context.Context, employeeId string) (*domain.EmployeePlants, error)
	GetByPlantId(ctx context.Context, plantId string) (*[]domain.EmployeePlants, error)
	Update(ctx context.Context, id string, emplplant *domain.EmployeePlants) error
	Delete(ctx context.Context, id string) error
}
