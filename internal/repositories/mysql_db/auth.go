package mysql_db

import (
	"context"
	"errors"
	"github.com/elSyarif/posnote-api.git/internal/core/domain"

	"github.com/elSyarif/posnote-api.git/internal/core/ports"
	"github.com/jmoiron/sqlx"
)

type authRepository struct {
	DB *sqlx.DB
}

func NewAuthRepository(connect *sqlx.DB) ports.AuthRepository {
	return &authRepository{
		DB: connect,
	}
}

func (respository *authRepository) Save(ctx context.Context, token string) error {
	query := "INSERT INTO authentication VALUES (?)"

	tx, err := respository.DB.Beginx()
	if err != nil {
		return errors.New(err.Error())
	}

	result, err := tx.ExecContext(ctx, query, token)
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

func (respository *authRepository) VerifyRefreshToken(ctx context.Context, token string) error {
	query := "SELECT token FROM authentication WHERE token = ?"

	auth := domain.Authentication{}
	tx, err := respository.DB.Beginx()
	if err != nil {
		return errors.New(err.Error())
	}

	err = tx.GetContext(ctx, &auth, query, token)
	if err != nil {
		return errors.New("token tidak ditemukan")
	}

	return nil
}

func (respository *authRepository) DeleteRefreshToken(ctx context.Context, token string) error {
	query := "DELETE FROM authentication WHERE token = ?"

	tx, err := respository.DB.Beginx()
	if err != nil {
		return errors.New(err.Error())
	}

	result, err := tx.ExecContext(ctx, query, token)
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
