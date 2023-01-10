package ports

import (
	"context"
	"github.com/elSyarif/posnote-api.git/internal/core/domain"
)

type EmployeeRepository interface {
	Save(ctx context.Context, employee *domain.Employees) (*domain.Employees, error)
	FindById(ctx context.Context, id string) (*domain.Employees, error)
	VerifyUsername(ctx context.Context, username string) error
	VerifyCredential(ctx context.Context, username string, password string) (string, error)
}

type EmployeeService interface {
	AddEmployee(ctx context.Context, employee *domain.Employees) (*domain.Employees, error)
	GetById(ctx context.Context, id string) (*domain.Employees, error)
	VerifyUsername(ctx context.Context, username string) error
	VerifyCredential(ctx context.Context, username string, password string) (string, error)
}
