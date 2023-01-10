package domain

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type JWTClaims struct {
	jwt.StandardClaims
	EmpId uuid.UUID `json:"emp_id,omitempty"`
}
