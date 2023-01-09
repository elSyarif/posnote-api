package domain

import "time"

type Plants struct {
	Id          string    `db:"id" json:"id"`
	Name        string    `db:"name" json:"name" binding:"required"`
	Description string    `db:"description" json:"description"`
	CreatedAt   time.Time `db:"created_at" json:"created_at" `
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}
