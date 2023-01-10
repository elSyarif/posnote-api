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

type plantRepository struct {
	DB *sqlx.DB
}

func NewPlantRepository(connect *sqlx.DB) ports.PlantsRepository {
	return &plantRepository{
		DB: connect,
	}
}

func (repository *plantRepository) Save(ctx context.Context, plant *domain.Plants) (*string, error) {
	query := "INSERT INTO plants VALUES (?, ?, ?, ?, ?, ?)"

	plant.Id = helper.GenerateUuid()
	plant.CreatedAt = time.Now()
	plant.UpdatedAt = time.Now()

	tx, err := repository.DB.Beginx()
	if err != nil {
		return nil, errors.New(err.Error())
	}

	result, err := tx.ExecContext(ctx, query, plant.Id, plant.Name, plant.Location, plant.Description, plant.CreatedAt, plant.UpdatedAt)
	if err != nil {
		tx.Rollback()
		return nil, errors.New(err.Error())
	}

	tx.Commit()
	if row, _ := result.RowsAffected(); row < 0 {
		return nil, errors.New(err.Error())
	}

	return &plant.Id, nil
}

func (repository *plantRepository) Find(ctx context.Context, name string) (*[]domain.Plants, error) {
	query := "SELECT * FROM plants WHERE name like ?"
	var plants []domain.Plants

	tx, err := repository.DB.Beginx()
	if err != nil {
		return nil, errors.New(err.Error())
	}

	err = tx.SelectContext(ctx, &plants, query, "%"+name+"%")
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &plants, nil
}

func (repository *plantRepository) FindById(ctx context.Context, id string) (*domain.Plants, error) {
	query := "SELECT * FROM plants WHERE id = ?"
	plant := domain.Plants{}

	tx, err := repository.DB.Beginx()
	if err != nil {
		return nil, errors.New(err.Error())
	}

	err = tx.GetContext(ctx, &plant, query, id)
	if err != nil {
		return nil, errors.New("plant tidak ditemukan")
	}

	return &plant, nil
}

func (repository *plantRepository) Update(ctx context.Context, id string, plant *domain.Plants) error {
	query := "UPDATE plants SET name = ?, location = ?, description = ?, updated_at = ?  WHERE id = ?"

	plant.UpdatedAt = time.Now()

	tx, err := repository.DB.Beginx()
	if err != nil {
		return errors.New(err.Error())
	}

	_, err = repository.FindById(ctx, id)
	if err != nil {
		return errors.New(err.Error())
	}

	result, err := tx.ExecContext(ctx, query, plant.Name, plant.Location, plant.Description, plant.UpdatedAt, id)
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

func (repository *plantRepository) Delete(ctx context.Context, id string) error {
	query := "DELETE FROM plants WHERE id = ?"

	tx, err := repository.DB.Beginx()
	if err != nil {
		return errors.New(err.Error())
	}

	_, err = repository.FindById(ctx, id)
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
