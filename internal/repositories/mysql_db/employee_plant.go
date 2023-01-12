package mysql_db

import (
	"context"
	"errors"
	"github.com/elSyarif/posnote-api.git/internal/helper"
	"time"

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
	query := "INSERT INTO employee_plants VALUES (?, ?, ?, ?, ?, ?, ?)"

	id := helper.GenerateUuid()
	emplplant.Id = id

	tx, err := repository.DB.Beginx()
	if err != nil {
		return nil, errors.New(err.Error())
	}

	result, err := tx.ExecContext(ctx, query, emplplant)
	if err != nil {
		tx.Rollback()
		return nil, errors.New(err.Error())
	}

	tx.Commit()
	if row, _ := result.RowsAffected(); row < 0 {
		return nil, errors.New(err.Error())
	}

	return &emplplant.Id, nil
}

func (repository *employeePlantRepository) Find(ctx context.Context, params string) (*[]domain.EmplPlantResponse, error) {
	query := `SELECT ep.id, e.fullname, ep.position, p.name as plant_name, ep.join_date
			  	FROM employee_plants ep
			  INNER JOIN employees e on e.id = ep.employee_id
			  INNER JOIN plants p on ep.plant_id = p.id`

	var emplPlanResponse []domain.EmplPlantResponse

	tx, err := repository.DB.Beginx()
	if err != nil {
		return nil, errors.New(err.Error())
	}

	err = tx.SelectContext(ctx, &emplPlanResponse, query)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &emplPlanResponse, nil
}

func (repository *employeePlantRepository) FindByEmployeeId(ctx context.Context, employeeId string) (*domain.EmplPlantResponse, error) {
	query := `SELECT ep.id, e.fullname, ep.position, p.name as plant_name, ep.join_date
			  	FROM employee_plants ep
			  INNER JOIN employees e on e.id = ep.employee_id
			  INNER JOIN plants p on ep.plant_id = p.id
			  WHERE ep.employee_id = ?`

	emplPlantResponse := domain.EmplPlantResponse{}

	tx, err := repository.DB.Beginx()
	if err != nil {
		return nil, errors.New(err.Error())
	}

	err = tx.GetContext(ctx, &emplPlantResponse, query, employeeId)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &emplPlantResponse, nil
}

func (repository *employeePlantRepository) FindByPlantId(ctx context.Context, plantId string) (*[]domain.EmplPlantResponse, error) {
	query := `SELECT ep.id, e.fullname, ep.position, p.name as plant_name, ep.join_date
			  	FROM employee_plants ep
			  INNER JOIN employees e on e.id = ep.employee_id
			  INNER JOIN plants p on ep.plant_id = p.id
			  WHERE ep.employee_id = ?`

	var emplPlantResponse []domain.EmplPlantResponse

	tx, err := repository.DB.Beginx()
	if err != nil {
		return nil, errors.New(err.Error())
	}

	err = tx.GetContext(ctx, &emplPlantResponse, query, plantId)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &emplPlantResponse, nil
}

func (repository *employeePlantRepository) Update(ctx context.Context, id string, emplplant *domain.EmployeePlants) error {
	query := `UPDATE employee_plants SET position = ?, join_date = ?, updated_at = ? WHERE id = ?`

	emplplant.UpdatedAt = time.Now().Local()

	tx, err := repository.DB.Beginx()
	if err != nil {
		return errors.New(err.Error())
	}

	result, err := tx.ExecContext(ctx, query, emplplant.Position, emplplant.JoinDate, emplplant.UpdatedAt, id)
	if err != nil {
		tx.Rollback()
		return errors.New(err.Error())
	}

	tx.Commit()
	if row, _ := result.RowsAffected(); row < 0 {
		return errors.New(err.Error())
	}

	return nil
}

func (repository *employeePlantRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM employee_plants WHERE id = ?`

	tx, err := repository.DB.Beginx()
	if err != nil {
		return errors.New(err.Error())
	}

	result, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		tx.Rollback()
		return errors.New(err.Error())
	}

	tx.Commit()
	if row, _ := result.RowsAffected(); row < 0 {
		return errors.New(err.Error())
	}

	return nil
}
