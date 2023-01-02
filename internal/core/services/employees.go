package services

import (
	"context"
	"errors"

	"github.com/elSyarif/posnote-api.git/internal/core/domain"
	"github.com/elSyarif/posnote-api.git/internal/core/ports"
)

type employeeService struct {
	repositroy ports.EmployeeRepository
}

func NewEmployeeService(repositoy ports.EmployeeRepository) ports.EmployeeService {
	return &employeeService{
		repositroy: repositoy,
	}
}

func (service *employeeService) AddEmployee(ctx context.Context, employee *domain.Employees) (*domain.Employees, error) {
	err := service.repositroy.VerifyUsername(ctx, employee.Username)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return service.repositroy.Save(ctx, employee)
}

func (service *employeeService) GetById(ctx context.Context, id string) (*domain.Employees, error) {
	return service.repositroy.FindById(ctx, id)
}

func (service *employeeService) VerifyUsername(ctx context.Context, username string) error {
	return service.repositroy.VerifyUsername(ctx, username)
}

func (service *employeeService) VerifyCredential(ctx context.Context, username string, password string) error {
	return service.repositroy.VerifyCredential(ctx, username, password)
}
