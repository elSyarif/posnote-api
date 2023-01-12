package mysql_db

import (
	"context"

	"github.com/elSyarif/posnote-api.git/internal/core/domain"
	"github.com/elSyarif/posnote-api.git/internal/core/ports"
	"github.com/jmoiron/sqlx"
)

type employeePlantRepository struct {
	DB *sqlx.DB
}

func NewEmployeePlantRepostory(connect *sqlx.DB) ports.EmployeePlantRepostory {
	return &employeePlantRepository{
		DB: connect,
	}
}

func (repository *employeePlantRepository) Save(ctx context.Context, emplplant *domain.EmployeePlants) (*string, error) {
	return nil, nil
}

func (repository *employeePlantRepository) Find(ctx context.Context, params ...string) (*[]domain.EmployeePlants, error) {
	return nil, nil
}

func (repository *employeePlantRepository) FindByEmployeeId(ctx context.Context, employeeId string) (*domain.EmployeePlants, error) {
	return nil, nil
}

func (repository *employeePlantRepository) FindByPlantId(ctx context.Context, plantId string) (*[]domain.EmployeePlants, error) {
	return nil, nil
}

func (repository *employeePlantRepository) Update(ctx context.Context, id string, emplplant *domain.EmployeePlants) error {
	return nil
}

func (repository *employeePlantRepository) Delete(ctx context.Context, id string) error {
	return nil
}
