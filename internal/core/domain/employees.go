package domain

import "time"

type Employees struct {
	Id        string    `json:"id" db:"id" `
	Fullname  string    `json:"fullname" db:"fullname" binding:"required"`
	Username  string    `json:"username" db:"username" binding:"required,lowercase"`
	Password  string    `json:"-" db:"password" binding:"required"`
	RoleId    string    `json:"role_id" db:"role_id" binding:"required"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
