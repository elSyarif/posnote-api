package domain

import (
	"time"

	"github.com/google/uuid"
)

type Token struct {
	EmpId     uuid.UUID  `json:"emp_id,omitempty"`
	Token     string     `json:"token,omitempty"`
	ExpiredAt *time.Time `json:"expired_at,omitempty"`
}
