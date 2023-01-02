package domain

import "time"

type Employees struct {
	Id        string    `json:"id"`
	Fullname  string    `json:"fullname" binding:"required"`
	Username  string    `json:"username" binding:"required,lowercase"`
	Password  string    `json:"password" binding:"required"`
	RoleId    string    `json:"role_id" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
