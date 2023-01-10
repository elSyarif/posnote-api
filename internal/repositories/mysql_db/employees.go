package mysql_db

import (
	"context"
	"errors"
	"time"

	"github.com/elSyarif/posnote-api.git/internal/core/domain"
	"github.com/elSyarif/posnote-api.git/internal/core/ports"
	"github.com/elSyarif/posnote-api.git/internal/helper"
	"github.com/jmoiron/sqlx"
)

type employeeRepository struct {
	DB *sqlx.DB
}

func NewEmployeeRepository(connect *sqlx.DB) ports.EmployeeRepository {
	return &employeeRepository{
		DB: connect,
	}
}

func (repsitory *employeeRepository) Save(ctx context.Context, employee *domain.Employees) (*domain.Employees, error) {
	query := "INSERT INTO employees VALUES (?, ?, ?, ?, ?, ?, ?)"
	employee.Id = helper.GenerateUuid()
	employee.CreatedAt = time.Now()
	employee.UpdatedAt = time.Now()
	hashPassword := helper.HashPassword(employee.Password)

	tx, err := repsitory.DB.Beginx()
	if err != nil {
		return nil, errors.New(err.Error())
	}
	result, err := tx.ExecContext(ctx, query, employee.Id, employee.Fullname, employee.Username, hashPassword, employee.RoleId, employee.CreatedAt, employee.UpdatedAt)
	if err != nil {
		tx.Rollback()
		return nil, errors.New(err.Error())
	}

	tx.Commit()
	if row, _ := result.RowsAffected(); row < 0 {
		return nil, errors.New(err.Error())
	}

	return employee, nil
}

func (repsitory *employeeRepository) FindById(ctx context.Context, id string) (*domain.Employees, error) {
	query := "SELECT id, fullname, username, role_id, created_at, updated_at FROM employees WHERE id = ?"

	employee := domain.Employees{}

	tx, err := repsitory.DB.Beginx()
	if err != nil {
		return nil, errors.New(err.Error())
	}

	err = tx.GetContext(ctx, &employee, query, id)
	if err != nil {
		return nil, errors.New("NOT_FOUND")
	}

	return &employee, nil
}

func (repsitory *employeeRepository) VerifyUsername(ctx context.Context, username string) error {
	query := "SELECT username FROM employees WHERE username = ?"
	var employee domain.Employees

	tx, err := repsitory.DB.Beginx()
	if err != nil {
		return errors.New(err.Error())
	}

	tx.GetContext(ctx, &employee, query, username)

	if employee.Username != "" {
		return errors.New("gagal, username suadah digunakan")
	}

	return nil
}

func (repsitory *employeeRepository) VerifyCredential(ctx context.Context, username string, password string) (string, error) {
	query := "select id, username, password FROM employees WHERE username = ?"
	var employee domain.Employees

	tx, err := repsitory.DB.Beginx()
	if err != nil {
		return "", errors.New(err.Error())
	}

	err = tx.GetContext(ctx, &employee, query, username)
	if err != nil {
		return "", errors.New(err.Error())
	}

	isMatch := helper.CheckPasswordHash(password, employee.Password)
	if isMatch {
		return employee.Id, nil
	}

	return "", errors.New("username/password salah")
}
