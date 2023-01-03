package helper

import (
	"errors"
	"os"
	"time"

	"github.com/elSyarif/posnote-api.git/internal/core/domain"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func GenerateAccessToken(payload uuid.UUID) (interface{}, error) {
	claims := domain.JWTClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    "posnote",
			ExpiresAt: time.Now().Add(time.Minute * 1).Unix(),
		},
		EmpId: payload,
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	accessToken, err := token.SignedString([]byte(os.Getenv("ACCESS_TOKEN_KEY")))
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return accessToken, nil
}

func GenerateRefreshToken(payload uuid.UUID) (interface{}, error) {
	claims := domain.JWTClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    "posnote",
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
		EmpId: payload,
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	refreshToken, err := token.SignedString([]byte(os.Getenv("REFRESH_TOKEN_KEY")))
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return refreshToken, nil
}

func VerifyToken(token string, secret string) (*jwt.Token, error) {
	parsed, err := jwt.ParseWithClaims(token, &domain.JWTClaims{}, func(t *jwt.Token) (any, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	return parsed, nil
}

func GetJWTData(token string, secret string) (*domain.JWTClaims, error) {
	tokenVerified, err := VerifyToken(token, secret)
	if err != nil {
		return nil, err
	}

	claims, ok := tokenVerified.Claims.(*domain.JWTClaims)
	if !ok && !tokenVerified.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
