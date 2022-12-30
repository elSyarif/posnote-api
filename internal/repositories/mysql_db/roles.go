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
		return nil, errors.New(err.Error())
	}

	if row, _ := result.RowsAffected(); row > 0 {
		return role, nil
	}

	return nil, errors.New(err.Error())
}
