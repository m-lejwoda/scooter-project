// Database interface mostly interfaces
package user

import (
	"context"
)

type UserRepository interface {
	GetByID(ctx context.Context, id string) (*User, error)
	GetByUsername(ctx context.Context, username string) (*User, error)
	Save(ctx context.Context, user *UserRegister) (*User, error)
}

type TokenRepository interface {
	GetAccessToken(ctx context.Context, name string) string
	GetRefreshToken(ctx context.Context, name string) string
	DeleteAccessToken(ctx context.Context, name string)
	DeleteRefreshToken(ctx context.Context, name string)
	CreateTokens(ctx context.Context, name string, tokenResponse TokenResponse, tokensTTL TokensTTL) (TokenResponse, error)
}
