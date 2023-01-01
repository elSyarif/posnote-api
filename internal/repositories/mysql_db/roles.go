package mysql_db

import (
	"context"
	"errors"
	"github.com/elSyarif/posnote-api.git/internal/core/domain"
	"github.com/elSyarif/posnote-api.git/internal/core/ports"
	"github.com/elSyarif/posnote-api.git/internal/helper"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type roleRepositroy struct {
	DB *sqlx.DB
}

func NewRolesRepository(connect *sqlx.DB) ports.RoleRepository {
	return &roleRepositroy{
		DB: connect,
	}
}

func (repository *roleRepositroy) Save(ctx context.Context, role *domain.Roles) (*domain.Roles, error) {
	query := "INSERT INTO roles VALUES(?, ?, ?)"
	role.Id = helper.GenerateUuid()

	tx, err := repository.DB.Beginx()
	if err != nil {
		return nil, errors.New(err.Error())
	}
	result, err := tx.ExecContext(ctx, query, role.Id, role.Name, role.Description)
	if err != nil {
		tx.Rollback()
		return nil, errors.New(err.Error())
	}

	tx.Commit()
	if row, _ := result.RowsAffected(); row < 0 {
		return nil, errors.New(err.Error())
	}

	return role, nil
}

func (repository *roleRepositroy) FindRole(ctx context.Context, name string) (*[]domain.Roles, error) {
	query := "SELECT * FROM roles WHERE name LIKE ?"
	var roles []domain.Roles

	tx, err := repository.DB.Beginx()
	if err != nil {
		return nil, errors.New(err.Error())
	}

	err = tx.SelectContext(ctx, &roles, query, "%"+name+"%")
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &roles, nil
}

func (repository *roleRepositroy) Update(ctx context.Context, roleId string, role *domain.Roles) error {
	query := "UPDATE roles set name = ?, description = ? WHERE id = ?"

	tx, err := repository.DB.Beginx()
	if err != nil {
		return errors.New(err.Error())
	}

	result, err := tx.ExecContext(ctx, query, role.Name, role.Description, roleId)
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

func (repository *roleRepositroy) Delete(ctx context.Context, roleId string) error {
	query := "DELETE FROM roles WHERE id = ?"

	tx, err := repository.DB.Beginx()
	if err != nil {
		return errors.New(err.Error())
	}

	_, err = tx.ExecContext(ctx, query, roleId)
	if err != nil {
		tx.Rollback()
		return errors.New(err.Error())
	}

	tx.Commit()
	return nil
}
