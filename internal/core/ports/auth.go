package ports

import "context"

type AuthRepository interface {
	Save(ctx context.Context, token string) error
	VerifyRefreshToken(ctx context.Context, token string) error
	DeleteRefreshToken(ctx context.Context, token string) error
}

type AuthService interface {
	AddRefreshToken(ctx context.Context, token string) error
	VerifyRefreshToken(ctx context.Context, token string) error
	DeleteRefreshToken(ctx context.Context, token string) error
}
