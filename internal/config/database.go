package config

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

func NewDB(configuration Config) (*sqlx.DB, error) {
	db, err := sqlx.Open(configuration.Get("DB_DRIVER"), configuration.Get("DB_SOURCE"))
	if err != nil {
		return nil, fmt.Errorf("DB error : %s", err.Error())
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("DB Connetion error : %s", err.Error())
	}

	return db, nil
}
